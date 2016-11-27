package utils

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"path"
)

func LoadConfig(name string, config interface{}) error {
	file, err := ioutil.ReadFile(path.Join("./config", name+".yml"))

	if err != nil {
		return err
	}

	return yaml.Unmarshal(file, config)
}
