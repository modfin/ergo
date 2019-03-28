package config

import (
	"github.com/caarlos0/env"
	"log"
	"sync"
)

type Config struct {
	ValA int `env:"VAL_A" envDefault:"1"`
	ValB int `env:"VAL_B" envDefault:"2"`
}

var (
	once sync.Once
	cfg  Config
)


func Get() *Config{
	once.Do(func() {
		cfg = Config{}
		if err := env.Parse(&cfg); err != nil {
			log.Panic("Couldn't parse Config from env: ", err)
		}
	})
	return &cfg
}
