package packets

type CharSelectChar struct {
	Slot byte
}

func (r *CharSelectChar) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Read(&r.Slot)

	return nil
}
