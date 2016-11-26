package packets

type CharSecondPasswordLogin struct {
	Seed      uint32
	AccountID uint32
	Result    uint16
}

func (r *CharSecondPasswordLogin) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(r.Seed)
	p.Write(r.AccountID)
	p.Write(r.Result)

	return nil
}
