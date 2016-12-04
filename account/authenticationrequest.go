package account

import "net"

type AuthenticationRequest struct {
	ID  string
	Key string

	PublicIP   net.IP
	PublicPort int
}
