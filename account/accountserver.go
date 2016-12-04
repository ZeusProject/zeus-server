package account

import "github.com/zeusproject/zeus-server/packets"

type AccountServer interface {
	CharServerStore() CharServerStore
	PacketDB() *packets.PacketDatabase
}
