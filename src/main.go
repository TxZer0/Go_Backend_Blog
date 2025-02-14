package main

import (
	"log"

	"github.com/TxZer0/Go_Backend_Blog/src/database"
	"github.com/TxZer0/Go_Backend_Blog/src/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Can't load file .env!", err)
	}

	database.InitDatabase()
	routes.InitRoutes()
}
