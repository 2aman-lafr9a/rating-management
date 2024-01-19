package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func getPostgresDataBaseUrl() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	if host == "" || user == "" || password == "" {
		host = "localhost"
		user = "postgres"
		password = "password"
	}
	return "host=" + host + " user=" + user + " password=" + password + " dbname=rating_management port=5432"
}

func PostgresConn() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(getPostgresDataBaseUrl()), &gorm.Config{},
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
