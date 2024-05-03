package initializers

import (
	"github.com/labstack/gommon/log"
	"github.com/lpernett/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
