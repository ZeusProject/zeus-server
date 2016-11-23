package net

import "net"

type Handler interface {
	Accept(conn net.Conn)
}
