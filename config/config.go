package config

import (
	"log"

	"github.com/Netflix/go-env"
)

var Config struct {
	DatabaseURL  string `env:"DATABASE_URL"`
	RedisAddress string `env:"REDIS_ADDR"`
}

func init() {
	if _, err := env.UnmarshalFromEnviron(&Config); err != nil {
		log.Fatal(err)
	}
	println(Config.DatabaseURL)
}
