// Package main Event CRUD API
//
//	@title			Event CRUD API
//	@version		1.0
//	@description	A simple event management API with authentication
//
//	@host		localhost:6969
//	@BasePath	/api/v1
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.
package main

import (
	"database/sql"
	"go-event-crud/internal/database"
	"go-event-crud/internal/env"
	"log"

	_ "go-event-crud/docs" // Import generated docs

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// init modals
	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 6969),
		jwtSecret: env.GetEnvString("JWT_SECRET", "random-secret"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}

}
