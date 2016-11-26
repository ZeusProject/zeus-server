package packets

type ZoneAid struct {
	AccountID uint32
}

func (r *ZoneAid) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(r.AccountID)

	return nil
}
