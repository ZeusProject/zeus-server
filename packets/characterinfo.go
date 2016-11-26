package packets

import "time"

type CharacterInfo struct {
	ID uint

	BaseExp  uint
	Zeny     uint
	JobExp   uint
	JobLevel uint

	BodyState   uint
	HealthState uint
	EffectState uint

	Virtue uint
	Honor  uint

	JobPoints uint16

	HP    uint
	MaxHP uint

	SP    uint16
	MaxSP uint16

	WalkSpeed uint16

	Job    uint16
	Head   uint16
	Body   uint16
	Weapon uint16
	Level  uint16

	SkillPoints uint16

	HeadBottom   uint16
	Shield       uint16
	HeadTop      uint16
	HeadMid      uint16
	HairColor    uint16
	ClothesColor uint16

	Name string

	Str byte
	Agi byte
	Vit byte
	Int byte
	Dex byte
	Luk byte

	Slot       byte
	HairColor2 byte

	Renamed bool

	MapName string

	DeleteDate *time.Time

	Robe       uint
	SlotChange uint
	Rename     uint
	Sex        bool
}

func (r *CharacterInfo) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Grow(147)

	p.Write(uint32(r.ID))
	p.Write(uint32(r.BaseExp))
	p.Write(uint32(r.Zeny))
	p.Write(uint32(r.JobExp))
	p.Write(uint32(r.JobLevel))
	p.Write(uint32(r.BodyState))
	p.Write(uint32(r.HealthState))
	p.Write(uint32(r.EffectState))
	p.Write(uint32(r.Virtue))
	p.Write(uint32(r.Honor))
	p.Write(uint16(r.JobPoints))
	p.Write(uint32(r.HP))
	p.Write(uint32(r.MaxHP))
	p.Write(uint16(r.SP))
	p.Write(uint16(r.MaxSP))
	p.Write(uint16(r.WalkSpeed))
	p.Write(uint16(r.Job))
	p.Write(uint16(r.Head))
	p.Write(uint16(r.Body))
	p.Write(uint16(r.Weapon))
	p.Write(uint16(r.Level))
	p.Write(uint16(r.SkillPoints))
	p.Write(uint16(r.HeadBottom))
	p.Write(uint16(r.Shield))
	p.Write(uint16(r.HeadTop))
	p.Write(uint16(r.HeadMid))
	p.Write(uint16(r.HairColor))
	p.Write(uint16(r.ClothesColor))

	p.WriteString(24, r.Name)

	p.Write(byte(r.Str))
	p.Write(byte(r.Agi))
	p.Write(byte(r.Vit))
	p.Write(byte(r.Int))
	p.Write(byte(r.Dex))
	p.Write(byte(r.Luk))
	p.Write(byte(r.Slot))
	p.Write(byte(r.HairColor2))

	if r.Renamed {
		p.Write(uint16(1))
	} else {
		p.Write(uint16(0))
	}

	p.WriteString(16, r.MapName)

	if r.DeleteDate != nil {
		p.Write(uint32(r.DeleteDate.Unix()))
	} else {
		p.Write(uint32(0))
	}

	p.Write(uint32(r.Robe))
	p.Write(uint32(r.SlotChange))
	p.Write(uint32(r.Rename))

	if r.Sex {
		p.Write(byte(1))
	} else {
		p.Write(byte(0))
	}

	return nil
}
