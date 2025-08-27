package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	Url    string
	Secret string
}

func Init() (*Config, error) {
	var config Config

	if err := LoadEnvVariables(); err != nil {
		return nil, err
	}

	config.Port = os.Getenv("PORT")
	config.Url = os.Getenv("DATABASE_URL")
	config.Secret = os.Getenv("SECRET_JWT")

	return &config, nil
}

func LoadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	return nil
}
