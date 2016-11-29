package packets

type CharDeleteChar3ReservedAck struct {
	CharID             uint
	Result             int
	DeleteReservedDate uint32
}

func (r *CharDeleteChar3ReservedAck) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(uint32(r.CharID))
	p.Write(r.Result)
	p.Write(r.DeleteReservedDate)

	return nil
}
