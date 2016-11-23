package packets

import "encoding/binary"

type LoginRequest struct {
	Version    uint32
	Username   string
	Password   string
	ClientType int
}

func (r *LoginRequest) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	binary.Read(p, binary.LittleEndian, &r.Version)

	return nil
}
