package packets

type NullPacket struct {
	Raw *RawPacket
}

func (r *NullPacket) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	r.Raw = p
	return nil
}

func (r *NullPacket) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	return nil
}
