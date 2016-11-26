package packets

type ZoneAckNameRequest struct {
	ID   uint32
	Name string
}

func (r *ZoneAckNameRequest) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(r.ID)
	p.WriteString(24, r.Name)

	return nil
}
