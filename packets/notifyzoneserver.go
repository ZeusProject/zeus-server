package packets

import "net"

type NotifyZoneServer struct {
	CharID  uint
	MapName string
	Address net.IP
	Port    uint16
}

func (r *NotifyZoneServer) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Write(uint32(r.CharID))
	p.WriteString(16, r.MapName)
	p.Write(r.Address.To4())
	p.Write(uint16(r.Port))

	return nil
}
