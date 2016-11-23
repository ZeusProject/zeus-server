package net

import "net"

type HandlerFn struct {
	Fn func(conn net.Conn)
}

func (h HandlerFn) Accept(conn net.Conn) {
	h.Fn(conn)
}
