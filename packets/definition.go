package packets

type Definition struct {
	Name   string
	ID     uint16
	Size   int
	Parser IncomingPacket
	Writer OutgoingPacket
}
