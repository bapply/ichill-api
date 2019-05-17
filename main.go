package main

import (
	"log"
	"os"

	"github.com/bapply/ichill-api/config"
	"github.com/bapply/ichill-api/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &config.App

	app.Router.Any("/songs", models.SongHandler)

	app.Router.Run(os.Getenv("PORT"))
}
