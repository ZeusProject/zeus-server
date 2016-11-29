package packets

type CharDeleteChar3CancelAck struct {
	CharID uint
	Result int
}

func (r *CharDeleteChar3CancelAck) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(uint32(r.CharID))
	p.Write(r.Result)

	return nil
}
