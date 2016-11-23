package login

import (
	gonet "net"

	"github.com/zeusproject/zeus-server/net"
)

type Client struct {
	*net.GameClient

	login *LoginServer
}

func NewClient(conn gonet.Conn, login *LoginServer) *Client {
	c := &Client{
		login: login,
	}

	c.GameClient = net.NewGameClient(conn, c.handlePacket, login.packetDatabase)

	return c
}

func (c *Client) handlePacket(p *net.Packet) {
}
