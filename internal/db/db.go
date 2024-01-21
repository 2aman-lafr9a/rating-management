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
	port:=os.Getenv("DB_PORT")
	db_name:=os.Getenv("DB_NAME")

	if host == "" || user == "" || password == "" {
		host = "localhost"
		user = "postgres"
		password = "password"
	}
	return "host=" + host + " user=" + user + " password=" + password + " dbname=" + db_name + " port="+port
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
