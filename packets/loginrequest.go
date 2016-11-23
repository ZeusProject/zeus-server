package packets

import "encoding/binary"

type LoginRequest struct {
	Version    uint32
	Username   string
	Password   string
	ClientType byte
}

func (r *LoginRequest) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	binary.Read(p, binary.LittleEndian, &r.Version)

	p.ReadString(24, &r.Username)
	p.ReadString(24, &r.Password)

	binary.Read(p, binary.LittleEndian, &r.ClientType)

	return nil
}
