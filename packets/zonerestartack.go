package packets

type ZoneRestartAck struct {
	Reason byte
}

func (r *ZoneRestartAck) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(r.Reason)

	return nil
}
