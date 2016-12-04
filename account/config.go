package account

type Config struct {
	Endpoint      string `yaml:"endpoint"`
	InterEndpoint string `yaml:"inter_endpoint"`
	PacketVersion uint32 `yaml:"packet_version"`

	CharServers []*CharServerDefinition `yaml:"char_servers"`
}
