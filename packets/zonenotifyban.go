package packets

type ZoneNotifyBan struct {
	Reason byte
}

func (r *ZoneNotifyBan) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(r.Reason)

	return nil
}
