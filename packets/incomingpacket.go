package packets

type IncomingPacket interface {
	Parse(db *PacketDatabase, d *Definition, r *RawPacket) error
}
