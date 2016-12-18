package database

import (
	"github.com/jinzhu/gorm"
	"github.com/o1egl/gormrus"
	"gopkg.in/redis.v5"

	// We just support postgres, so we load its driver here
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database interface {
	Sql() *gorm.DB
	Redis() *redis.Client

	AuthCookie() AuthCookieStore
}

type CommonDatabase struct {
	sql   *gorm.DB
	redis *redis.Client
}

func NewDatabase(config *Config) (*CommonDatabase, error) {
	sql, err := setupGorm(config)

	if err != nil {
		return nil, err
	}

	redisclient, err := setupRedis(config)

	if err != nil {
		return nil, err
	}

	return &CommonDatabase{
		sql:   sql,
		redis: redisclient,
	}, nil
}

func (d *CommonDatabase) Sql() *gorm.DB {
	return d.sql
}

func (d *CommonDatabase) Redis() *redis.Client {
	return d.redis
}

func setupGorm(config *Config) (*gorm.DB, error) {
	gorm, err := gorm.Open("postgres", config.Postgres.ConnectionUrl())

	if err != nil {
		return nil, err
	}

	gorm.LogMode(true)
	gorm.SetLogger(gormrus.New())

	return gorm, nil
}

func setupRedis(config *Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})

	return client, nil
}
