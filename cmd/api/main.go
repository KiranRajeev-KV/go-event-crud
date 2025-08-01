package main

import (
	"database/sql"
	"go-event-crud/internal/database"
	"go-event-crud/internal/env"
	"log"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/joho/godotenv/autoload"
)

type application struct {
	port int
	jwtSecret string
	models database.Models
}

func main() {
	db,err := sql.Open("sqlite3","./data.db")
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.Close()

	// init modals
	models := database.NewModels(db)

	app  := &application{
		port: env.GetEnvInt("PORT",6969),
		jwtSecret: env.GetEnvString("JWT_SECRET","random-secret"),
		models: models,
	}

	if err:= app.serve(); err != nil {
		log.Fatal(err)
	}

	

}
