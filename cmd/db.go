package main

import (
	"fmt"
	"github.com/EugeneTsydenov/go-user-service/cmd/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	host := env.GetEnv()["DB_HOST"]
	user := env.GetEnv()["DB_USERNAME"]
	password := env.GetEnv()["DB_PASSWORD"]
	dbName := env.GetEnv()["DB_NAME"]
	port := env.GetEnv()["DB_PORT"]
	sslMode := env.GetEnv()["DB_SSL_MODE"]

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslMode)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("Connected to database successfully!")
}
