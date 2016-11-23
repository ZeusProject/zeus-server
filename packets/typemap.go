package packets

var typeMap = map[string]Packet{
	"SS_NULL":  &NullPacket{},
	"CA_LOGIN": &LoginRequest{},
}
