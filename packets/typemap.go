package packets

var incomingTypeMap = map[string]IncomingPacket{
	"SS_NULL":  &NullPacket{},
	"CA_LOGIN": &LoginRequest{},
	"CH_ENTER": &CharEnter{},
}

var outgoingTypeMap = map[string]OutgoingPacket{
	"SS_NULL":                &NullPacket{},
	"AC_ACCEPT_LOGIN":        &AcceptLoginResponse{},
	"HC_ACCEPT_ENTER":        &AcceptCharEnter{},
	"HC_REFUSE_ENTER":        &RefuseCharEnter{},
	"HC_ACCEPT2":             &AcceptCharEnter2{},
	"HC_BLOCK_CHARACTER":     &BlockCharacter{},
	"HC_SECOND_PASSWD_LOGIN": &SecondPasswordLogin{},
}
