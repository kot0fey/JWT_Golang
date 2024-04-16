package initializers

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	_ "log"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
