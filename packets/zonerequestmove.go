package packets

type ZoneRequestMove struct {
	Position Position
}

func (r *ZoneRequestMove) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	r.Position.Parse(db, d, p)

	return nil
}
