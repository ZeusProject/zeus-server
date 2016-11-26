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

func (c *Client) Enter(p *packets.CharEnter) {
	chars := []*packets.CharacterInfo{
		&packets.CharacterInfo{
			ID:        150000,
			HP:        10000,
			MaxHP:     10000,
			SP:        3000,
			MaxSP:     3000,
			WalkSpeed: 137,
			Job:       4065,
			Head:      1,
			Body:      1,
			Level:     99,
			JobLevel:  50,
			HairColor: 6,
			Name:      "PROTETAS",
			Slot:      1,
			Str:       99,
			Agi:       99,
			Vit:       99,
			Int:       99,
			Dex:       99,
			Luk:       99,
			MapName:   "prontera.gat",
			Sex:       true,
		},
	}

	c.accountId = p.AccountID

	// Send the AID to the client
	c.SendRaw(p.AccountID)

	c.Send(&packets.AcceptCharEnter2{
		NormalSlots:     9,
		PremiumSlots:    0,
		BillingSlots:    0,
		ProducibleSlots: 9,
		ValidSlots:      9,
	})

	// Send all characters
	c.Send(&packets.AcceptCharEnter{
		TotalSlotCount:   9,
		PremiumSlotStart: 9,
		PremiumSlotEnd:   9,
		Chars:            chars,
	})

	// Banned characters
	c.Send(&packets.BlockCharacter{})

	// Skip PIN check
	c.Send(&packets.SecondPasswordLogin{
		AccountID: p.AccountID,
		Seed:      0xDEADBEEF,
		Result:    0,
	})
}

func (c *Client) handlePacket(d *packets.Definition, p packets.IncomingPacket) {
	c.log.WithFields(logrus.Fields{
		"packet": d.Name,
		"id":     d.ID,
		"parsed": p,
	}).Debug("packet arrived")

	switch p := p.(type) {
	case *packets.CharEnter:
		c.Enter(p)
	case *packets.NullPacket:
		c.log.WithFields(logrus.Fields{
			"packet": d.Name,
			"id":     d.ID,
		}).Warning("unhandled packet")
	}
}
