package packets

import (
	"net"
)

type AcceptLoginResponse struct {
	AuthenticationCode uint32
	AccountID          uint32
	AccountLevel       uint32
	Sex                byte
	Servers            []CharServer
}

type CharServer struct {
	IP       net.IP
	Port     uint16
	Name     string
	Users    uint16
	State    uint16
	Property uint16
}

func (r *AcceptLoginResponse) Write(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Grow(43 + len(r.Servers)*32)

	p.Write(r.AuthenticationCode)
	p.Write(r.AccountID)
	p.Write(r.AccountLevel)
	p.Skip(30)
	p.Write(r.Sex)

	for _, s := range r.Servers {
		p.Write([]byte(s.IP.To4()))
		p.Write(s.Port)
		p.WriteString(20, s.Name)
		p.Write(s.Users)
		p.Write(s.State)
		p.Write(s.Property)
	}

	return nil
}
