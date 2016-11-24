package char

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

	switch p := p.(type) {
	case *packets.CharEnter:
		c.SendRaw(p.AccountID)

		// c.Send(&packets.RefuseCharEnter{
		// 	Reason: 0,
		// })

		c.Send(&packets.AcceptCharEnter2{
			NormalSlots:     9,
			PremiumSlots:    0,
			BillingSlots:    0,
			ProducibleSlots: 9,
			ValidSlots:      9,
		})

		c.Send(&packets.AcceptCharEnter{
			TotalSlotCount:   9,
			PremiumSlotStart: 9,
			PremiumSlotEnd:   9,
		})

		c.Send(&packets.BlockCharacter{})

		c.Send(&packets.SecondPasswordLogin{
			AccountID: p.AccountID,
			Seed:      0xDEADBEEF,
			Result:    4,
		})
	case *packets.NullPacket:
		c.log.WithFields(logrus.Fields{
			"packet": d.Name,
			"id":     d.ID,
		}).Warning("unhandled packet")
	}
}
