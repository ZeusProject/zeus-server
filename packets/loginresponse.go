package packets

type LoginResponse struct {
}

func (r *LoginResponse) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	return nil
}
