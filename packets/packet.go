package packets

type Packet interface {
	Parse(db *PacketDatabase, d *Definition, r *RawPacket) error
}
