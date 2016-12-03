package char

import (
	gonet "net"

	"github.com/Sirupsen/logrus"
	"github.com/zeusproject/zeus-server/net"
	"github.com/zeusproject/zeus-server/packets"
)

type Client struct {
	*net.GameClient

	server *Server
	log    *logrus.Entry

	chars []*packets.CharacterInfo

	accountId uint32
}

func NewClient(conn gonet.Conn, server *Server) *Client {
	c := &Client{
		server: server,
		log:    logrus.WithField("component", "client"),
	}

	c.GameClient = net.NewGameClient(conn, c, server.packetDatabase)

	return c
}

func (c *Client) loadCharacters() {
	c.chars = []*packets.CharacterInfo{
		&packets.CharacterInfo{
			ID:        150000,
			HP:        10000,
			MaxHP:     10000,
			SP:        3000,
			MaxSP:     3000,
			WalkSpeed: 137,
			Job:       4065,
			Head:      1,
			Body:      1,
			Level:     99,
			JobLevel:  50,
			HairColor: 6,
			Name:      "PROTETAS",
			Slot:      1,
			Str:       99,
			Agi:       99,
			Vit:       99,
			Int:       99,
			Dex:       99,
			Luk:       99,
			MapName:   "prontera.gat",
			Sex:       true,
		},
	}
}

func (c *Client) Enter(p *packets.CharEnter) {
	c.accountId = p.AccountID

	// Send the AID to the client
	c.SendRaw(p.AccountID)

	c.loadCharacters()

	c.Send(&packets.CharSlotsInfo{
		NormalSlots:     9,
		PremiumSlots:    0,
		BillingSlots:    0,
		ProducibleSlots: 9,
		ValidSlots:      9,
	})

	// Send all characters
	c.Send(&packets.CharAcceptEnter{
		TotalSlotCount:   9,
		PremiumSlotStart: 9,
		PremiumSlotEnd:   9,
		Chars:            c.chars,
	})

	// Banned characters
	c.Send(&packets.CharBlockCharacter{})

	// Skip PIN check
	c.Send(&packets.CharSecondPasswordLogin{
		AccountID: p.AccountID,
		Seed:      0xDEADBEEF,
		Result:    0,
	})
}

func (c *Client) SelectChar(slot byte) {
	c.Send(&packets.CharNotifyZoneServer{
		CharID:  150000,
		MapName: "prontera.gat",
		Address: gonet.ParseIP("127.0.0.1"),
		Port:    5121,
	})
}

func (c *Client) CheckCharName(name string) int {
	return 0
}

func (c *Client) MakeChar(name string, slot byte, haircolor uint16, hairstyle uint16, startingjobid uint16, sex byte) {
	var error byte = 0

	errorcode := c.CheckCharName(name)

	if (startingjobid != 0) && (startingjobid != 4218) { //JOB_NOVICE && JOB_SUMMONER
		errorcode = -2
	}

	charsex := false //Female
	if sex == 1 {
		charsex = true //Male
	}

	switch errorcode {
	case -1:
		error = 0x00 //Charname already exists								CHAR_NAME_EXISTS (custom enum names -ZzZz-)
	case -2:
		error = 0xFF //Char creation denied								CHAR_CREATION_DENIED
	case -3:
		error = 0x01 //You are underaged									CHAR_UNDERAGED
	case -4:
		error = 0x02 //Symbols in Character Names are forbidden			CHAR_FORBIDDEN_SYMBOLS
	case -5:
		error = 0x03 //You are not elegible to open the Character Slot		CHAR_NO_SLOT
	}

	if errorcode < 0 {
		c.Send(&packets.CharRefuseMakeChar{
			ErrorCode: error,
		})
	} else {
		c.Send(&packets.CharAcceptMakeChar{
			CharInfo: &packets.CharacterInfo{
				ID:           150002,
				BaseExp:      0,
				Zeny:         0,
				JobExp:       0,
				JobLevel:     1,
				BodyState:    0,
				HealthState:  0,
				EffectState:  0,
				Virtue:       0,
				Honor:        0,
				JobPoints:    48,
				HP:           40,
				MaxHP:        40,
				SP:           11,
				MaxSP:        11,
				WalkSpeed:    150,
				Job:          startingjobid,
				Head:         hairstyle,
				Body:         0,
				Weapon:       0,
				Level:        1,
				SkillPoints:  0,
				HeadBottom:   0,
				Shield:       0,
				HeadTop:      0,
				HeadMid:      0,
				HairColor:    haircolor,
				ClothesColor: 0,
				Name:         name,
				Str:          1,
				Agi:          1,
				Vit:          1,
				Int:          1,
				Dex:          1,
				Luk:          1,
				Slot:         slot,
				Renamed:      true,
				MapName:      "prontera.gat",
				DeleteDate:   nil,
				Robe:         0,
				SlotChange:   0,
				Rename:       0,
				Sex:          charsex,
			},
		})
	}
}

func (c *Client) DeleteChar(charid uint) {
	// First we need a DB xd
	var result int = 1
	var deletedate uint32 = 1480443780 - 1480440180
	/*
		// Check if character exists
		if CharExists(charid) == nil {
			result = 3 // Database error
		}

		if db_test := DB.Char.getDummy(); db_test == nil {
			result = 3 // Database error
		}

		// Check if character is already on deletion list
		if char_dd := DB.Char.getDeleteDate(); char_dd {
			result = 0 // Unknown error
		}

		// Check if character is in a guild
		if hasGuild := DB.Char.getGuild(); hasGuild && c.server.config.guilddelete {
			result = 4 // To delete a character you must withdraw from the guild.
		}

		// Check if character is in a party
		if hasParty := DB.Char.getParty(); hasParty && c.server.config.partydelete {
			result = 5 // To delete a character you must withdraw from the party.
		}

		// Set deletion date
		deletiondate := Timer.Now()
		if result := DB.Char.setDeleteDate(); result == nil {
			result = 3 // Database error
		}
	*/
	c.Send(&packets.CharDeleteChar3ReservedAck{
		CharID:             charid,
		Result:             result,
		DeleteReservedDate: deletedate,
	})
}

func (c *Client) CancelDeleteChar(charid uint) {
	// 1 (0x718): none/success, (if char id not in deletion process): An unknown error has occurred.
	// 2 (0x719): A database error occurred.
	var result int = 1

	c.Send(&packets.CharDeleteChar3CancelAck{
		CharID: charid,
		Result: result,
	})
}

func (c *Client) HandlePacket(d *packets.Definition, p packets.IncomingPacket) {
	c.log.WithFields(logrus.Fields{
		"packet": d.Name,
		"id":     d.ID,
		"parsed": p,
	}).Debug("packet arrived")

	switch p := p.(type) {
	case *packets.CharEnter:
		c.Enter(p)
	case *packets.CharSelectChar:
		c.SelectChar(p.Slot)
	case *packets.CharMakeChar:
		c.MakeChar(p.Name, p.Slot, p.HairColor, p.HairStyle, p.StartingJobID, p.Sex)
	case *packets.CharDeleteChar3Reserved:
		c.DeleteChar(p.CharID)
	case *packets.CharDeleteChar3Cancel:
		c.CancelDeleteChar(p.CharID)
	case *packets.Ping:
	case *packets.NullPacket:
		c.log.WithFields(logrus.Fields{
			"packet": d.Name,
			"id":     d.ID,
		}).Warning("unhandled packet")
	}
}

func (c *Client) OnDisconnect(err error) {
}
