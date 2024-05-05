package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	host := getEnv()["DB_HOST"]
	user := getEnv()["DB_USERNAME"]
	password := getEnv()["DB_PASSWORD"]
	dbName := getEnv()["DB_NAME"]
	port := getEnv()["DB_PORT"]
	sslMode := getEnv()["DB_SSL_MODE"]

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslMode)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("Connected to database successfully!")
}
