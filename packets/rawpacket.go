package packets

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type RawPacket struct {
	*bytes.Buffer

	ID uint16
}

func NewRawPacket(id uint16) *RawPacket {
	return NewRawPacketFromBuffer(id, make([]byte, 0))
}

func NewRawPacketFromBuffer(id uint16, buffer []byte) *RawPacket {
	return &RawPacket{
		Buffer: bytes.NewBuffer(buffer),
		ID:     id,
	}
}

func (p *RawPacket) Skip(n int) {
	p.Buffer.Write(make([]byte, n))
}

func (p *RawPacket) Read(v interface{}) {
	binary.Read(p.Buffer, binary.LittleEndian, v)
}

func (p *RawPacket) ReadString(len int, s *string) {
	b := make([]byte, len)

	p.Buffer.Read(b)

	*s = string(b)
}

func (p *RawPacket) Write(v interface{}) error {
	return binary.Write(p.Buffer, binary.LittleEndian, v)
}

func (p *RawPacket) WriteString(size int, s string) {
	str := []byte(s)
	data := make([]byte, size)

	if len(str) < size {
		copy(data, str)
	} else {
		copy(data, str[:size])
	}

	p.Write(data)
}

func (p *RawPacket) Hex() string {
	return fmt.Sprintf("%04x", p.ID)
}

func (p *RawPacket) String() string {
	return fmt.Sprintf("%04x (%d bytes)", p.ID, p.Len())
}
