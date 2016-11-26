package packets

var incomingTypeMap = map[string]IncomingPacket{
	"SS_NULL": &NullPacket{},
	"SS_PING": &Ping{},

	"CA_LOGIN": &AccountLogin{},

	"CH_ENTER":       &CharEnter{},
	"CH_SELECT_CHAR": &CharSelectChar{},

	"CZ_ENTER":            &ZoneEnter{},
	"CZ_REQUEST_TIME":     &ZoneRequestTime{},
	"CZ_NOTIFY_ACTORINIT": &ZoneNotifyActorInit{},
	"CZ_REQNAME":          &ZoneNameRequest{},
}

var outgoingTypeMap = map[string]OutgoingPacket{
	"SS_NULL": &NullPacket{},

	"AC_ACCEPT_LOGIN": &AccountAcceptLogin{},

	"HC_ACCEPT_ENTER":        &CharAcceptEnter{},
	"HC_REFUSE_ENTER":        &CharRefuseEnter{},
	"HC_SLOT_INFO":           &CharSlotsInfo{},
	"HC_BLOCK_CHARACTER":     &CharBlockCharacter{},
	"HC_SECOND_PASSWD_LOGIN": &CharSecondPasswordLogin{},
	"HC_NOTIFY_ZONESVR":      &CharNotifyZoneServer{},

	"ZC_AID":          &ZoneAid{},
	"ZC_ACCEPT_ENTER": &ZoneAcceptEnter{},
	"ZC_NOTIFY_TIME":  &ZoneNotifyTime{},
	"ZC_ACK_REQNAME":  &ZoneAckNameRequest{},
}
