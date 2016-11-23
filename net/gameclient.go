package net

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/packets"
)

const (
	MaxPacketSize = 1024 * 16
)

type GameClient struct {
	conn    net.Conn
	handler PacketHandler
	db      *packets.PacketDatabase
	log     *logrus.Entry
}

func NewGameClient(conn net.Conn, handler PacketHandler, db *packets.PacketDatabase) *GameClient {
	return &GameClient{
		conn:    conn,
		handler: handler,
		db:      db,
		log:     logrus.WithField("component", "client"),
	}
}

func (c *GameClient) Start() {
	go c.run()
}

func (c *GameClient) Disconnect() error {
	return c.conn.Close()
}

func (c *GameClient) Send(p packets.OutgoingPacket) error {
	def, raw, err := c.db.Write(p)

	if err != nil {
		return err
	}

	totalLen := 2 + raw.Len()

	if def.Size == -1 {
		totalLen += 2
	}

	data := bytes.NewBuffer(make([]byte, totalLen))

	binary.Write(data, binary.LittleEndian, raw.ID)

	if def.Size == -1 {
		binary.Write(data, binary.LittleEndian, totalLen)
	}

	binary.Write(data, binary.LittleEndian, raw.Bytes())

	_, err = c.conn.Write(data.Bytes())

	return err
}

func (c *GameClient) run() {
	var packet uint16
	var size uint16

	raw := make([]byte, MaxPacketSize)
	buffer := bytes.NewBuffer(raw)
	offset := 0
	state := 0

	for true {
		read, err := c.conn.Read(raw[offset:])

		if err != nil {
			c.log.WithError(err).Error("error reading from socket")
			c.Disconnect()
			return
		}

		offset += read

		if state == 0 {
			if offset >= 2 {
				err := binary.Read(buffer, binary.LittleEndian, &packet)

				if err != nil {
					c.log.WithError(err).Error("error reading from socket")
					c.Disconnect()
					return
				}

				state++
			}
		}

		if state == 1 {
			s, ok := c.db.GetSize(packet)

			if !ok {
				c.log.WithField("packet", fmt.Sprintf("%04x", packet)).Error("invalid packet")
				c.Disconnect()
				return
			}

			if s == -1 {
				if offset >= 4 {
					err := binary.Read(buffer, binary.LittleEndian, &size)

					if err != nil {
						c.log.WithError(err).Error("error reading from socket")
						c.Disconnect()
						return
					}

					state++
				}
			} else {
				size = uint16(s)
				state++
			}
		}

		if state == 2 {
			if offset >= int(size) {
				raw := &packets.RawPacket{
					Buffer: bytes.NewBuffer(raw[:size]),
					ID:     packet,
					Size:   size,
				}

				def, packet, err := c.db.Parse(raw)

				if err != nil {
					c.log.WithField("packet", fmt.Sprintf("%04x", packet)).Error("invalid packet")
					c.Disconnect()
					return
				}

				c.handler(def, packet)

				state = 0
				offset = 0
				buffer.Reset()
			}
		}
	}
}
