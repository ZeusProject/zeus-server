package packets

type CharRefuseMakeChar struct {
	ErrorCode byte
}

func (r *CharRefuseMakeChar) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(r.ErrorCode)

	return nil
}
