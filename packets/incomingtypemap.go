package packets

var incomingTypeMap = map[string]IncomingPacket{
	"SS_NULL": &NullPacket{},
	"SS_PING": &Ping{},

	"CA_LOGIN": &AccountLogin{},

	"CH_ENTER":                 &CharEnter{},
	"CH_SELECT_CHAR":           &CharSelectChar{},
	"CH_MAKE_CHAR":             &CharMakeChar{},
	"CH_DELETE_CHAR3_RESERVED": &CharDeleteChar3Reserved{},
	"CH_DELETE_CHAR3":          &NullPacket{}, //&CharDeleteChar3{},
	"CH_DELETE_CHAR3_CANCEL":   &CharDeleteChar3Cancel{},

	"CZ_ENTER":            &ZoneEnter{},
	"CZ_REQUEST_TIME":     &ZoneRequestTime{},
	"CZ_NOTIFY_ACTORINIT": &ZoneNotifyActorInit{},
	"CZ_REQNAME":          &ZoneNameRequest{},
	"CZ_REQUEST_MOVE":     &ZoneRequestMove{},
}
