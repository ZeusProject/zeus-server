package packets

type Position struct {
	X         int16
	Y         int16
	Direction byte
}

func (r *Position) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write([]byte{
		byte(r.X >> 2),
		byte((r.X << 6) | ((r.Y >> 4) & 0x3F)),
		byte((r.Y << 4) | (int16(r.Direction) & 0xF)),
	})

	return nil
}
