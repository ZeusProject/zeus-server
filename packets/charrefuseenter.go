package packets

type CharRefuseEnter struct {
	Reason byte
}

func (r *CharRefuseEnter) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(r.Reason)

	return nil
}
