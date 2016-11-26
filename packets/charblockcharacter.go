package packets

type CharBlockCharacter struct {
}

func (r *CharBlockCharacter) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	return nil
}
