package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	database = "coursesdb"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host = %s port =%d user =%s password=%s database=%s sslmode= disable", host, port, user, password, database)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		log.Println("Error Connecting to Database: ", err)
	}
	return db
}
