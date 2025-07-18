package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/db/models"
	_ "github.com/lib/pq"
)

type DBContext struct {
	Queries *models.Queries
}

func Connect(host string, user string, password string, dbName string, dbPort string, sslMode string) *models.Queries {
	dbUrl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		password,
		dbName,
		dbPort,
		sslMode,
	)

	database, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Error connecting to database" + err.Error())
	}

	return models.New(database)
}
