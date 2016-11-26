package packets

type Position2 struct {
	X0 uint16
	Y0 uint16
	X1 uint16
	Y1 uint16
	SX uint16
	SY uint16
}

func (r *Position2) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write([]byte{
		byte(r.X0 >> 2),
		byte((r.X0 << 6) | ((r.Y0 >> 4) & 0x3f)),
		byte((r.Y0 << 4) | ((r.X1 >> 6) & 0x0f)),
		byte((r.X1 << 2) | ((r.Y1 >> 8) & 0x03)),
		byte(r.Y1),
		byte((r.SX << 4) | (r.SY & 0x0f)),
	})

	return nil
}
