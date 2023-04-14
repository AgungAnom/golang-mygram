package database

import (
	"fmt"
	"golang-mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     	= "localhost"
	user     	= "postgres1"
	password 	= "postgres1"
	port     	= 5432
	dbname		= "mygram"

)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host,port,user,password,dbname)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{},models.Socialmedia{})
	fmt.Println("Connected to Database")
}

func GetDB() *gorm.DB {
	return db
}