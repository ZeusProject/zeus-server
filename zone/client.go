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

	accountId  uint32
	clientTick uint32

	x, y int
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
	c.x = 150
	c.y = 150
	c.accountId = p.AccountID

	c.Send(&packets.ZoneAid{
		AccountID: c.accountId,
	})

	c.Send(&packets.ZoneAcceptEnter{
		Tick: c.server.Time.GetTick(),
		Position: packets.Position{
			X:         uint16(c.x),
			Y:         uint16(c.y),
			Direction: 6,
		},
		XSize: 5,
		YSize: 5,
		Font:  0,
		Sex:   p.Sex,
	})
}

func (c *Client) SyncTime(tick uint32) {
	c.clientTick = tick

	c.Send(&packets.ZoneNotifyTime{
		Tick: c.server.Time.GetTick(),
	})
}

func (c *Client) SetCharLoaded() {
}

func (c *Client) NotifyName(id uint32) {
	c.Send(&packets.ZoneAckNameRequest{
		ID:   id,
		Name: "espadahabil",
	})
}

func (c *Client) Move(x, y, direction int) {
	c.Send(&packets.ZoneNotifyPlayerMove{
		Tick: c.clientTick,
		Position: packets.Position2{
			X0: uint16(c.x),
			Y0: uint16(c.y),
			X1: uint16(x),
			Y1: uint16(y),
			SX: 8,
			SY: 8,
		},
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
	case *packets.ZoneNotifyActorInit:
		c.SetCharLoaded()
	case *packets.ZoneRequestTime:
		c.SyncTime(p.Tick)
	case *packets.ZoneNameRequest:
		c.NotifyName(p.ID)
	case *packets.ZoneRequestMove:
		c.Move(int(p.Position.X), int(p.Position.Y), int(p.Position.Direction))
	case *packets.Ping:
	case *packets.NullPacket:
		c.log.WithFields(logrus.Fields{
			"packet": d.Name,
			"id":     d.ID,
		}).Warning("unhandled packet")
	}
}
