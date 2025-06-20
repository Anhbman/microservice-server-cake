package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
}

func SetupEnv() (cfg Config, err error) {

	godotenv.Load()

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return Config{}, errors.New("HTTP_PORT environment variable is not set")
	}

	return Config{
		ServerPort: httpPort,
	}, nil
}
