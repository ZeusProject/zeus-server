package packets

type BlockCharacter struct {
}

func (r *BlockCharacter) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	return nil
}
