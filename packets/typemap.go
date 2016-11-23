package packets

var incomingTypeMap = map[string]IncomingPacket{
	"SS_NULL":  &NullPacket{},
	"CA_LOGIN": &LoginRequest{},
}

var outgoingTypeMap = map[string]OutgoingPacket{
	"SS_NULL":         &NullPacket{},
	"AC_ACCEPT_LOGIN": &AcceptLoginResponse{},
}
