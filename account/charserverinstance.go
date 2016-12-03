package account

import "net"

type CharServerInstance struct {
	PublicIP   net.IP
	PublicPort int
}

func (i *CharServerInstance) Equal(other *CharServerInstance) bool {
	return i.PublicIP.Equal(other.PublicIP) && i.PublicPort == other.PublicPort
}
