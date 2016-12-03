package net

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/packets"
)

const (
	MaxPacketSize = 1024 * 16
)

type DisconnectFunc func(err error)

type GameClient struct {
	conn        net.Conn
	stream      *bufio.ReadWriter
	handler     PacketHandler
	db          *packets.PacketDatabase
	log         *logrus.Entry
	forcedClose bool
}

func NewGameClient(conn net.Conn, handler PacketHandler, db *packets.PacketDatabase) *GameClient {
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)

	return &GameClient{
		conn:    conn,
		stream:  bufio.NewReadWriter(r, w),
		handler: handler,
		db:      db,
		log:     logrus.WithField("component", "gameclient"),
	}
}

func (c *GameClient) Start() {
	go c.run()
}

func (c *GameClient) Disconnect() error {
	return c.DisconnectWithError(nil)
}

func (c *GameClient) DisconnectWithError(err error) error {
	c.forcedClose = true

	if err := c.conn.Close(); err != nil {
		return err
	}

	c.handler.OnDisconnect(err)

	return nil
}

func (c *GameClient) SendRaw(data interface{}) error {
	err := binary.Write(c.stream, binary.LittleEndian, data)

	if err != nil {
		return err
	}

	return c.stream.Flush()
}

func (c *GameClient) Send(p packets.OutgoingPacket) error {
	def, raw, err := c.db.Write(p)

	if err != nil {
		return err
	}

	_, err = c.stream.Write(raw.Bytes())

	if err != nil {
		return err
	}

	c.log.WithFields(logrus.Fields{
		"packet": def.Name,
		"id":     def.ID,
		"length": raw.Len(),
		"parsed": p,
	}).Debug("packet sent")

	return c.stream.Flush()
}

func (c *GameClient) run() {
	for true {
		var packet uint16
		var size int

		if err := binary.Read(c.stream, binary.LittleEndian, &packet); err != nil {
			c.log.WithError(err).Error("error reading from socket")
			c.DisconnectWithError(err)
			return
		}

		size, ok := c.db.GetSize(packet)
		variable := size == -1

		if !ok {
			c.log.WithField("packet", fmt.Sprintf("%04x", packet)).Error("invalid packet")
			c.Disconnect()
			return
		}

		if variable {
			var s uint16

			if err := binary.Read(c.stream, binary.LittleEndian, &s); err != nil {
				c.log.WithError(err).Error("error reading from socket")
				c.DisconnectWithError(err)
				return
			}

			size = int(s)
		}

		data := make([]byte, size)

		if _, err := c.stream.Read(data); err != nil {
			c.log.WithError(err).Error("error reading from socket")
			c.DisconnectWithError(err)
			return
		}

		raw := packets.NewRawPacketFromBuffer(packet, data)

		def, parsed, err := c.db.Parse(raw)

		if err != nil {
			c.log.WithField("packet", fmt.Sprintf("%04x", packet)).Error("invalid packet")
			c.Disconnect()
			return
		}

		c.handler.HandlePacket(def, parsed)
	}

	if !c.forcedClose {
		c.handler.OnDisconnect(nil)
	}
}
