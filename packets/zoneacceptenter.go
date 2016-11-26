package packets

type ZoneAcceptEnter struct {
	Tick     uint32
	Position Position
	XSize    byte
	YSize    byte
	Font     uint16
	Sex      bool
}

func (r *ZoneAcceptEnter) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(uint32(r.Tick))

	r.Position.Write(db, d, p)

	p.Write(byte(r.XSize))
	p.Write(byte(r.YSize))
	p.Write(uint16(r.Font))

	if r.Sex {
		p.Write(byte(1))
	} else {
		p.Write(byte(0))
	}

	return nil
}
