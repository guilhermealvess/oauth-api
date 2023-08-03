package config

import (
	"log"

	"github.com/Netflix/go-env"
)

type config struct {
	DatabaseURL string `env:"DATABASE_URL"`
}

var Config *config

func init() {
	if _, err := env.UnmarshalFromEnviron(Config); err != nil {
		log.Fatal(err)
	}
}
