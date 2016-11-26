package packets

import (
	"encoding/binary"
	"errors"
	"reflect"
)

type PacketDatabase struct {
	Version uint32

	incomingMap map[uint16]*Definition
	outgoingMap map[reflect.Type]*Definition
}

func New(version uint32) (*PacketDatabase, error) {
	db := &PacketDatabase{
		Version:     version,
		incomingMap: make(map[uint16]*Definition),
		outgoingMap: make(map[reflect.Type]*Definition),
	}

	pv, found := PacketVersions[int(version)]

	if !found {
		return nil, errors.New("invalid packet version")
	}

	for id, d := range pv.Packets {
		db.Register(d.Packet, id, d.Size)
	}

	return db, nil
}

func (db *PacketDatabase) Register(name string, id uint16, size int) {
	var parser IncomingPacket
	var writer OutgoingPacket

	parser, isParser := incomingTypeMap[name]

	if !isParser {
		parser = &NullPacket{}
	}

	writer, isWriter := outgoingTypeMap[name]

	if !isWriter {
		writer = &NullPacket{}
	}

	d := &Definition{
		Name:   name,
		ID:     id,
		Size:   size,
		Parser: parser,
		Writer: writer,
	}

	db.incomingMap[id] = d

	if isWriter {
		db.outgoingMap[reflect.TypeOf(writer).Elem()] = d
	}
}

func (db *PacketDatabase) GetSize(packet uint16) (int, bool) {
	def, ok := db.incomingMap[packet]

	if !ok {
		return 0, false
	}

	return def.Size, true
}

func (db *PacketDatabase) Parse(raw *RawPacket) (*Definition, IncomingPacket, error) {
	def, ok := db.incomingMap[raw.ID]

	if !ok {
		return nil, nil, errors.New("invalid packet")
	}

	typ := reflect.TypeOf(def.Parser).Elem()
	packet := reflect.New(typ).Interface().(IncomingPacket)

	err := packet.Parse(db, def, raw)

	if err != nil {
		return nil, nil, err
	}

	return def, packet, nil
}

func (db *PacketDatabase) Write(p OutgoingPacket) (*Definition, *RawPacket, error) {
	typ := reflect.TypeOf(p).Elem()
	def, ok := db.outgoingMap[typ]

	if !ok {
		return nil, nil, errors.New("invalid packet")
	}

	len := 0
	variable := def.Size == -1

	if variable {
		len = 4
	} else {
		len = def.Size
	}

	raw := NewRawPacket(def.ID, len)

	raw.Write(uint16(def.ID))

	if def.Size == -1 {
		raw.Write(uint16(0))
	}

	err := p.Write(db, def, raw)

	if err != nil {
		return nil, nil, err
	}

	if def.Size == -1 {
		data := raw.Bytes()

		binary.LittleEndian.PutUint16(data[2:4], uint16(raw.Size))
	}

	return def, raw, nil
}
