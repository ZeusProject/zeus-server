package login

import (
	gonet "net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
	"github.com/zeusproject/zeus-server/packets"
)

type Client struct {
	*net.GameClient

	login *LoginServer
	log   *logrus.Entry
}

func NewClient(conn gonet.Conn, login *LoginServer) *Client {
	c := &Client{
		login: login,
		log:   logrus.WithField("component", "client"),
	}

	c.GameClient = net.NewGameClient(conn, c.handlePacket, login.packetDatabase)

	return c
}

func (c *Client) handlePacket(d *packets.Definition, p packets.Packet) {
	c.log.WithFields(logrus.Fields{
		"packet": d.Name,
		"id":     d.ID,
		"parsed": p,
	}).Debug("packet arrived")
}
