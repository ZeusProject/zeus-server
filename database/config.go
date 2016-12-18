package database

import (
	"fmt"
	"net/url"
)

type Config struct {
	Redis    RedisConfig    `yaml:"redis"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type PostgresConfig struct {
	Host     string            `yaml:"host"`
	Port     int               `yaml:"port"`
	User     string            `yaml:"user"`
	Password string            `yaml:"password"`
	Database string            `yaml:"database"`
	Options  map[string]string `yaml:"options"`
}

func (c *PostgresConfig) ConnectionUrl() string {
	query := make(url.Values)

	for k, v := range c.Options {
		query[k] = []string{v}
	}

	url := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.User, c.Password),
		Host:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:     c.Database,
		RawQuery: query.Encode(),
	}

	return url.String()
}
