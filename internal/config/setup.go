package config

import (
	"fmt"
	"log"

	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/app_context"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/db"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/routers"
)

var AppContext *app_context.Context

func Setup() {
	LoadEnv()

	AppContext = &app_context.Context{}

	setupDatabase()
	startServer()
}

// Loads the port
func loadPort() string {
	portString := GetEnv("PORT", "8000")
	if portString == "" {
		log.Fatal("PORT is undefined in the .env file")
	}

	return portString
}

// Setup the database
func setupDatabase() {
	AppContext.DB = &db.DBContext{
		Queries: db.Connect(
			GetEnv("DB_HOST", "localhost"),
			GetEnv("DB_USER", "postgres"),
			GetEnv("DB_PASSWORD", "Postgres@WebomindApps"),
			GetEnv("DB_NAME", "postgres"),
			GetEnv("DB_PORT", "5432"),
			GetEnv("SSL_MODE", "disable"),
		),
	}
	fmt.Printf("Database running on %s", GetEnv("DB_HOST", "localhost"))
}

// Defines and start the router
func startServer() {
	port := loadPort()
	r := routers.NewRouter(AppContext)
	r.Run(":" + port)
}
