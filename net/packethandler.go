package net

import "github.com/zeusproject/zeus-server/packets"

type PacketHandler func(d *packets.Definition, p packets.IncomingPacket)
