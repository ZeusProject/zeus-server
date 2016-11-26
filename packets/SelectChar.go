package packets

type SelectChar struct {
	Slot byte
}

func (r *SelectChar) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Read(&r.Slot)

	return nil
}
