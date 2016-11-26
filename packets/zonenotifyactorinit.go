package packets

type ZoneNotifyActorInit struct {
}

func (r *ZoneNotifyActorInit) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	return nil
}
