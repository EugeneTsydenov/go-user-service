package pkg

import (
	"github.com/joho/godotenv"
	"log"
)

func GetEnv() map[string]string {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return env
}
