package net

import (
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

func (c *GameClient) SendRaw(data interface{}) error {
	return binary.Write(c.conn, binary.LittleEndian, data)
}

func (c *GameClient) Send(p packets.OutgoingPacket) error {
	def, raw, err := c.db.Write(p)

	if err != nil {
		return err
	}

	_, err = c.conn.Write(raw.Bytes())

	if err != nil {
		return err
	}

	c.log.WithFields(logrus.Fields{
		"packet": def.Name,
		"id":     def.ID,
		"length": raw.Len(),
		"parsed": p,
	}).Debug("packet sent")

	return err
}

func (c *GameClient) run() {
	buffer := make([]byte, MaxPacketSize)
	offset := 0

	for true {
		read, err := c.conn.Read(buffer[offset:])

		if err != nil {
			c.log.WithError(err).Error("error reading from socket")
			c.Disconnect()
			return
		}

		offset += read

		header := 2
		if offset < header {
			continue
		}

		packet := binary.LittleEndian.Uint16(buffer[0:2])
		size, ok := c.db.GetSize(packet)

		if !ok {
			c.log.WithField("packet", fmt.Sprintf("%04x", packet)).Error("invalid packet")
			c.Disconnect()
			return
		}

		if size == -1 {
			header += 2

			if offset < 4 {
				continue
			}

			size = int(binary.LittleEndian.Uint16(buffer[2:4]))
		}

		if offset < size {
			continue
		}

		raw := packets.NewRawPacketFromBuffer(packet, buffer[:size])
		raw.Next(header)

		def, parsed, err := c.db.Parse(raw)

		if err != nil {
			c.log.WithField("packet", fmt.Sprintf("%04x", packet)).Error("invalid packet")
			c.Disconnect()
			return
		}

		c.handler(def, parsed)

		copy(buffer, buffer[size:offset])
		offset = 0
	}
}
