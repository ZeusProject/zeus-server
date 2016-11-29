package packets

type CharDeleteChar3Reserved struct {
	CharID uint
}

func (r *CharDeleteChar3Reserved) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Read(&r.CharID)

	return nil
}
