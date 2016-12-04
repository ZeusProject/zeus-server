package char

type Config struct {
	Endpoint      string `yaml:"endpoint"`
	PacketVersion uint32 `yaml:"packet_version"`

	AccountInter AccountInterConfig `yaml:"account_inter"`
}
