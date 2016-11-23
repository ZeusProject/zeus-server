package packets

import (
	"errors"
	"reflect"
)

type PacketDatabase struct {
	Version uint32
	Packets map[uint16]*Definition
}

func New(version uint32) (*PacketDatabase, error) {
	db := &PacketDatabase{
		Version: version,
		Packets: make(map[uint16]*Definition),
	}

	db.Register("CA_LOGIN", 0x64, 0x37)

	return db, nil
}

func (db *PacketDatabase) Register(name string, id uint16, size int) {
	typ, ok := typeMap[name]

	if !ok {
		typ = &NullPacket{}
	}

	db.Packets[id] = &Definition{
		Name: name,
		ID:   id,
		Size: size,
		Type: typ,
	}
}

func (db *PacketDatabase) GetSize(packet uint16) (int, bool) {
	def, ok := db.Packets[packet]

	if !ok {
		return 0, false
	}

	return def.Size, true
}

func (db *PacketDatabase) Parse(raw *RawPacket) (*Definition, Packet, error) {
	def, ok := db.Packets[raw.ID]

	if !ok {
		return nil, nil, errors.New("invalid packet")
	}

	typ := reflect.TypeOf(def.Type).Elem()
	packet := reflect.New(typ).Interface().(Packet)

	err := packet.Parse(db, def, raw)

	if err != nil {
		return nil, nil, err
	}

	return def, packet, nil
}
