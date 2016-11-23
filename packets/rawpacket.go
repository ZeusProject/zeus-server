package packets

import (
	"bytes"
	"fmt"
)

type RawPacket struct {
	*bytes.Buffer

	ID   uint16
	Size uint16
}

func (p *RawPacket) Hex() string {
	return fmt.Sprintf("%04x", p.ID)
}

func (p *RawPacket) String() string {
	return fmt.Sprintf("%04x (%d bytes)", p.ID, p.Size)
}
