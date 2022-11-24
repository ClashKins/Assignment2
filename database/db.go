package database

import (
	"LATIHAN1/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "demage12"
	dbPort   = 5432
	dbname   = "postgres"
	db *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, dbPort)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}
	db.AutoMigrate(&models.GormModel{}, &models.Order{}, &models.Items{})
}

func GetDB() *gorm.DB{
	return db
}