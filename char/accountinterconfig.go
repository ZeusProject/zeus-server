package char

type AccountInterConfig struct {
	Endpoint   string `yaml:"endpoint"`
	ID         string `yaml:"id"`
	Key        string `yaml:"key"`
	PublicIP   string `yaml:"public_ip"`
	PublicPort int    `yaml:"public_port"`
}
