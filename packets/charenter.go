package packets

type CharEnter struct {
	AccountID          uint32
	AuthenticationCode uint32
	UserLevel          uint32
	ClientType         uint16
	Sex                byte
}

func (r *CharEnter) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Read(&r.AccountID)
	p.Read(&r.AuthenticationCode)
	p.Read(&r.UserLevel)
	p.Read(&r.ClientType)
	p.Read(&r.Sex)

	return nil
}
