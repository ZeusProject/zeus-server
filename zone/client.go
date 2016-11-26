package zone

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

	accountId uint32
}

func NewClient(conn gonet.Conn, server *Server) *Client {
	c := &Client{
		server: server,
		log:    logrus.WithField("component", "client"),
	}

	c.GameClient = net.NewGameClient(conn, c.handlePacket, server.packetDatabase)

	return c
}

func (c *Client) Enter(p *packets.ZoneEnter) {
	c.accountId = p.AccountID

	c.Send(&packets.ZoneAid{
		AccountID: c.accountId,
	})

	c.Send(&packets.ZoneAcceptEnter{
		Tick: p.Tick,
		Position: packets.Position{
			X:         150,
			Y:         150,
			Direction: 1,
		},
		XSize: 5,
		YSize: 5,
		Font:  0,
		Sex:   p.Sex,
	})
}

func (c *Client) handlePacket(d *packets.Definition, p packets.IncomingPacket) {
	c.log.WithFields(logrus.Fields{
		"packet": d.Name,
		"id":     d.ID,
		"parsed": p,
	}).Debug("packet arrived")

	switch p := p.(type) {
	case *packets.ZoneEnter:
		c.Enter(p)
	case *packets.Ping:
	case *packets.NullPacket:
		c.log.WithFields(logrus.Fields{
			"packet": d.Name,
			"id":     d.ID,
		}).Warning("unhandled packet")
	}
}
