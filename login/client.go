package login

import (
	gonet "net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
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

func (c *Client) handlePacket(p *net.Packet) {
	c.log.WithFields(logrus.Fields{
		"packet": p.Packet,
		"size":   p.Size,
	}).Debug("packet arrived")
}
