package conf

import (
	"github.com/daodao97/xgo/xapp"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Addr     string `yaml:"addr" env:"REDIS_ADDR"`
	Password string `yaml:"password" env:"REDIS_PASSWORD"`
	DB       int    `yaml:"db" env:"REDIS_DB"`
}

type config struct {
	Redis *Redis `yaml:"redis"`
}

var Conf config

func Init() error {
	return xapp.InitConf(&Conf)
}

func GetRedis() *redis.Options {
	return &redis.Options{
		Addr:     Conf.Redis.Addr,
		Password: Conf.Redis.Password,
		DB:       Conf.Redis.DB,
	}
}

func Get() *config {
	return &Conf
}
