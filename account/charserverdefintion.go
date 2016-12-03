package account

type CharServerDefinition struct {
	ID      string `yaml:"id"`
	Key     string `yaml:"key"`
	Name    string `yaml:"name"`
	Enabled bool   `yaml:"enabled"`
}
