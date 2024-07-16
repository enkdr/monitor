package config

import (
	"errors"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	PORT          int
	DB_PORT       string
	DB_HOST       string
	DB_USER       string
	DB_NAME       string
	DB_PASSWORD   string
	TEMPLATE_PATH string
}

func GetConfig() (Config, error) {

	config := Config{}

	envVars := map[string]*string{
		"DB_PORT":       &config.DB_PORT,
		"DB_HOST":       &config.DB_HOST,
		"DB_USER":       &config.DB_USER,
		"DB_NAME":       &config.DB_NAME,
		"DB_PASSWORD":   &config.DB_PASSWORD,
		"TEMPLATE_PATH": &config.TEMPLATE_PATH,
	}

		// loop through env(s)
	for key, value := range envVars {
		envValue := os.Getenv(key)
		if envValue == "" {
			return Config{}, errors.New(key + " is required")
		}
		*value = envValue
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		return Config{}, errors.New("PORT is required")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return Config{}, errors.New("invalid PORT value")
	}
	config.PORT = port

	return config, nil
}
