package packets

type ZoneRefuseEnter struct {
	Reason byte
}

func (r *ZoneRefuseEnter) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(r.Reason)

	return nil
}
