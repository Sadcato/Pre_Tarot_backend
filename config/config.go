package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ArkAPIKey   string
	EndpointID  string
	BaseURL     string
	RedisHost   string
	RedisPort   string
	RedisPasswd string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		ArkAPIKey:   os.Getenv("ARK_API_KEY"),
		EndpointID:  os.Getenv("ENDPOINT_ID"),
		BaseURL:     os.Getenv("BASE_URL"),
		RedisHost:   os.Getenv("REDIS_HOST"),
		RedisPort:   os.Getenv("REDIS_PORT"),
		RedisPasswd: os.Getenv("REDIS_PASSWORD"),
	}, nil
}
