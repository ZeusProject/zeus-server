package inter

type Config struct {
	Endpoint      string `yaml:"endpoint"`
	Secret        string `yaml:"secret"`
	PacketVersion uint32 `yaml:"packet_version"`
}
