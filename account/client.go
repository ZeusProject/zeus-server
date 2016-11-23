package account

import (
	gonet "net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
	"github.com/zeusproject/zeus-server/packets"
)

type Client struct {
	*net.GameClient

	server *Server
	log    *logrus.Entry
}

func NewClient(conn gonet.Conn, server *Server) *Client {
	c := &Client{
		server: server,
		log:    logrus.WithField("component", "client"),
	}

	c.GameClient = net.NewGameClient(conn, c.handlePacket, server.packetDatabase)

	return c
}

func (c *Client) handlePacket(d *packets.Definition, p packets.IncomingPacket) {
	c.log.WithFields(logrus.Fields{
		"packet": d.Name,
		"id":     d.ID,
		"parsed": p,
	}).Debug("packet arrived")

	switch p.(type) {
	case *packets.LoginRequest:
		res := &packets.AcceptLoginResponse{
			AuthenticationCode: 0xDEADBEEF,
			AccountID:          0x1337AAAA,
			AccountLevel:       0xBAADCAFE,
			Sex:                0,
			Servers: []packets.CharServer{
				packets.CharServer{
					IP:       gonet.ParseIP("127.0.0.1"),
					Port:     6121,
					Name:     "Zeus Project",
					Users:    1000,
					State:    0,
					Property: 0,
				},
			},
		}

		c.Send(res)
	}
}
