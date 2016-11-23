package net

import (
	"bytes"
	"fmt"
)

type Packet struct {
	Packet uint16
	Size   uint16
	Data   *bytes.Buffer
}

func (p *Packet) Hex() string {
	return fmt.Sprintf("%x", p.Packet)
}

func (p *Packet) String() string {
	return fmt.Sprintf("%x (%d bytes)", p.Packet, p.Size)
}
