package packets

type Position struct {
	X         uint16
	Y         uint16
	Direction byte
}

func (r *Position) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write([]byte{
		byte(r.X >> 2),
		byte((r.X << 6) | ((r.Y >> 4) & 0x3F)),
		byte((r.Y << 4) | (uint16(r.Direction) & 0xF)),
	})

	return nil
}

func (r *Position) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	raw := make([]byte, 3)

	p.Read(raw)

	r.X = uint16(((raw[0] & 0xFF) << 2) | (raw[1] >> 6))
	r.Y = uint16(((raw[1] & 0x3F) << 4) | (raw[2] >> 4))
	r.Direction = raw[2] & 0xF

	return nil
}
