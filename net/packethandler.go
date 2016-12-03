package net

import "github.com/zeusproject/zeus-server/packets"

type PacketHandler interface {
	HandlePacket(d *packets.Definition, p packets.IncomingPacket)
	OnDisconnect(err error)
}
