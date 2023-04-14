package main

import (
	"golang-mygram/database"
	"golang-mygram/router"
)

func main(){
	database.StartDB()
	r := router.StartApp()
	r.Run(":3000")
}