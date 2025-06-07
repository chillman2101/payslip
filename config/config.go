package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB         string
	ServerPort string
	RedisUrl   string
	AuthKey    string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Config{
		DB:         os.Getenv("DB"),
		ServerPort: os.Getenv("SERVER_PORT"),
		RedisUrl:   os.Getenv("REDIS_URL"),
		AuthKey:    os.Getenv("AUTH_KEY"),
	}, nil
}
