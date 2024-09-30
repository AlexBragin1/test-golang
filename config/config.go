package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"gitlab.com/usdtkg/payout/logger"
)

var Global Config

type Config struct {
	AppName string `env:"APP_NAME,required"`
	AppUrl  string `env:"APP_URL,required"`

	ServerHost string `env:"SERVER_HOST"`
	ServerPort int    `env:"SERVER_PORT"`

	DB

	JWTSecret string `env:"JWT_SECRET"`
}

type DB struct {
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASS"`
	Host string `env:"DB_HOST"`
	Port int    `env:"DB_PORT"`
	Name string `env:"DB_NAME"`
}

func init() {
	if err := env.Parse(&Global); err != nil {
		fmt.Printf("%+v\n", err)
	}

	if Global.AppName == "" {
		panic("APP_NAME not set")
	}

	logger.Info("Init config", "APP_NAME", Global.AppName)
}
