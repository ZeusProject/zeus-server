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

	c.GameClient = net.NewGameClient(conn, c, server.packetdb)

	return c
}

func (c *Client) Login(p *packets.AccountLogin) {
	servers, err := c.server.charserverStore.Servers()

	if err != nil {
		c.DisconnectWithError(err)
		return
	}

	packetServers := make([]packets.CharServer, len(servers))

	for i, s := range servers {
		instance := s.RandomInstance()

		if instance == nil {
			instance = &CharServerInstance{
				PublicIP:   gonet.ParseIP("0.0.0.0"),
				PublicPort: 0,
			}
		}

		packetServers[i] = packets.CharServer{
			IP:       instance.PublicIP,
			Port:     uint16(instance.PublicPort),
			Name:     s.Name,
			Users:    uint16(s.OnlinePlayers),
			State:    0,
			Property: 0,
		}
	}

	c.Send(&packets.AccountAcceptLogin{
		AuthenticationCode: 0xDEADBEEF,
		AccountID:          2000000,
		AccountLevel:       0xBAADCAFE,
		Sex:                0,
		Servers:            packetServers,
	})
}

func (c *Client) HandlePacket(d *packets.Definition, p packets.IncomingPacket) {
	c.log.WithFields(logrus.Fields{
		"packet": d.Name,
		"id":     d.ID,
		"parsed": p,
	}).Debug("packet arrived")

	switch p := p.(type) {
	case *packets.AccountLogin:
		c.Login(p)
	case *packets.NullPacket:
		c.log.WithFields(logrus.Fields{
			"packet": d.Name,
			"id":     d.ID,
		}).Warning("unhandled packet")
	}
}

func (c *Client) OnDisconnect(err error) {
}
