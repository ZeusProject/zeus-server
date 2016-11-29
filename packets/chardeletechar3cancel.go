package packets

type CharDeleteChar3Cancel struct {
	CharID uint
}

func (r *CharDeleteChar3Cancel) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Read(&r.CharID)

	return nil
}
