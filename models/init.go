package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitialDatabase() error {

	var err error
	dsn := "host=localhost port=5432 dbname=postgres user=postgres password=64021531 connect_timeout=10 sslmode=prefer"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	autoMigrateDatabase()

	fmt.Println("Database connection successfully...")
	return nil
}

func autoMigrateDatabase() {
	Db.AutoMigrate(&User{})
}
