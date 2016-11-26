package packets

type ZoneNotifyTime struct {
	Tick uint32
}

func (r *ZoneNotifyTime) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(r.Tick)

	return nil
}
