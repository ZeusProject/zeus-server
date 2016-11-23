package packets

type OutgoingPacket interface {
	Write(db *PacketDatabase, d *Definition, r *RawPacket) error
}
