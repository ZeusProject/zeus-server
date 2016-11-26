package packets

type ZoneNameRequest struct {
	ID uint32
}

func (r *ZoneNameRequest) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Read(&r.ID)

	return nil
}
