package packets

type ZoneNotifyNewEntry struct {
	Type               byte
	Aid                uint32
	Gid                uint32
	Speed              uint16
	BodyState          uint16
	HealthState        uint16
	EffectState        uint32
	Job                uint16
	Head               uint16
	Weapon             uint32
	HeadBottom         uint16
	HeadTop            uint16
	HeadMid            uint16
	HairColor          uint16
	ClothesColor       uint16
	HeadDirection      uint16
	Robe               uint16
	GuildID            uint32
	GuildEmblemVersion uint16
	Honor              uint16
	Virtue             uint32
	IsPK               bool
	Gender             bool
	Position           Position
	XSize              byte
	YSize              byte
	CLevel             uint16
	Font               uint16
	MaxHP              uint32
	HP                 uint32
	IsBoss             bool
	Body               uint16
	Name               string
}

func (r *ZoneNotifyNewEntry) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	return nil
}
