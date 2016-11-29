package packets

type CharAcceptMakeChar struct {
	CharInfo *CharacterInfo
}

func (r *CharAcceptMakeChar) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	r.CharInfo.Write(db, d, p)

	return nil
}
