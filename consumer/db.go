package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	dns = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
)

func ConnectToDb() {
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed To Connect With Db")
	}
}
