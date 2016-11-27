package packets

type CharMakeChar struct {
	Name		  string
	Slot		  byte
	HairColor	  uint16
	HairStyle	  uint16
	StartingJobID uint16
	//Unknown     uint16 (or [2]byte)
	Sex			  byte
}

func (r *CharMakeChar) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.ReadString(24, &r.Name)
	p.Read(&r.Slot)
	p.Read(&r.HairColor)
	p.Read(&r.HairStyle)
	p.Read(&r.StartingJobID)
	p.Skip(2) //FIXME: "Sex" always 0
	p.Read(&r.Sex)

	return nil
}