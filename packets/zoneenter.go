package packets

type ZoneEnter struct {
	AccountID          uint32
	CharID             uint32
	AuthenticationCode uint32
	Tick               uint32
	Sex                bool
}

func (r *ZoneEnter) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Read(&r.AccountID)
	p.Read(&r.CharID)
	p.Read(&r.AuthenticationCode)
	p.Read(&r.Tick)
	p.Read(&r.Sex)

	return nil
}
