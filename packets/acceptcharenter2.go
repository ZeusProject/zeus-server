package packets

type AcceptCharEnter2 struct {
	NormalSlots     byte
	PremiumSlots    byte
	BillingSlots    byte
	ProducibleSlots byte
	ValidSlots      byte
	Chars           []*CharacterInfo
}

func (r *AcceptCharEnter2) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Grow(25)

	p.Write(r.NormalSlots)
	p.Write(r.PremiumSlots)
	p.Write(r.BillingSlots)
	p.Write(r.ProducibleSlots)
	p.Write(r.ValidSlots)
	p.Skip(20)

	for _, ch := range r.Chars {
		ch.Write(db, d, p)
	}

	return nil
}
