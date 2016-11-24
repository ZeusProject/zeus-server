package packets

type AcceptCharEnter struct {
	TotalSlotCount   byte
	PremiumSlotStart byte
	PremiumSlotEnd   byte
	Chars            []*CharacterInfo
}

func (r *AcceptCharEnter) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Grow(23)

	p.Write(r.TotalSlotCount)
	p.Write(r.PremiumSlotStart)
	p.Write(r.PremiumSlotEnd)
	p.Skip(20)

	for _, ch := range r.Chars {
		ch.Write(db, d, p)
	}

	return nil
}
