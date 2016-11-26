package packets

type ZoneNotifyPlayerMove struct {
	Tick     uint32
	Position Position2
}

func (r *ZoneNotifyPlayerMove) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(r.Tick)
	r.Position.Write(db, d, p)

	return nil
}
