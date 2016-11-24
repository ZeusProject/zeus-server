package packets

type CharacterInfo struct {
}

func (r *CharacterInfo) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Grow(144)

	return nil
}
