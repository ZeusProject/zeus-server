package packets

type RefuseCharEnter struct {
	Reason byte
}

func (r *RefuseCharEnter) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(r.Reason)

	return nil
}
