package packets

type Ping struct {
	AccountID uint32
}

func (r *Ping) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Read(&r.AccountID)

	return nil
}
