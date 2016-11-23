package net

import "bytes"

type Packet struct {
	Packet uint16
	Size   uint16
	Data   *bytes.Buffer
}
