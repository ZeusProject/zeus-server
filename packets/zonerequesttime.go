package packets

type ZoneRequestTime struct {
	Tick uint32
}

func (r *ZoneRequestTime) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Read(&r.Tick)

	return nil
}
