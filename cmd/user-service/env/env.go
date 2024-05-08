package env

import (
	"github.com/joho/godotenv"
	"log"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv() map[string]string {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return env
}
