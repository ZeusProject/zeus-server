package packets

var outgoingTypeMap = map[string]OutgoingPacket{
	"SS_NULL": &NullPacket{},

	"AC_ACCEPT_LOGIN": &AccountAcceptLogin{},

	"HC_ACCEPT_ENTER":          &CharAcceptEnter{},
	"HC_REFUSE_ENTER":          &CharRefuseEnter{},
	"HC_SLOT_INFO":             &CharSlotsInfo{},
	"HC_BLOCK_CHARACTER":       &CharBlockCharacter{},
	"HC_SECOND_PASSWD_LOGIN":   &CharSecondPasswordLogin{},
	"HC_NOTIFY_ZONESVR":        &CharNotifyZoneServer{},
	"HC_REFUSE_MAKE_CHAR":      &CharRefuseMakeChar{},
	"HC_ACCEPT_MAKE_CHAR":      &CharAcceptMakeChar{},
	"HC_DELETE_CHAR3_RESERVED": &CharDeleteChar3ReservedAck{},
	"HC_DELETE_CHAR3":          &NullPacket{}, //&CharDeleteChar3Ack{},
	"HC_DELETE_CHAR3_CANCEL":   &CharDeleteChar3CancelAck{},

	"ZC_AID":               &ZoneAid{},
	"ZC_ACCEPT_ENTER":      &ZoneAcceptEnter{},
	"ZC_NOTIFY_TIME":       &ZoneNotifyTime{},
	"ZC_ACK_REQNAME":       &ZoneAckNameRequest{},
	"ZC_NOTIFY_PLAYERMOVE": &ZoneNotifyPlayerMove{},
	"ZC_REFUSE_ENTER":      &ZoneRefuseEnter{},
	"ZC_RESTART_ACK":       &ZoneRestartAck{},
	"ZC_NOTIFY_BAN":        &ZoneNotifyBan{},
	"ZC_NOTIFY_NEWENTRY":   &ZoneNotifyNewEntry{},
}
