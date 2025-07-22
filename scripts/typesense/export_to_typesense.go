package typesense

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/app_context"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/config"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/db"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/search"
	"github.com/google/uuid"
	"github.com/typesense/typesense-go/typesense/api"
)

var AppContext *app_context.Context

type ProductTableMetaData struct {
	ID        uuid.UUID
	Company   string
	Model     string
	ModelType sql.NullString
	Brand     string
	Category  string
	PartNo    string
}

func Setup() {
	config.LoadEnv()

	AppContext = &app_context.Context{}

	setupDatabase()
	startSearchEngine()

}

// Setup the database
func setupDatabase() {
	AppContext.DB = &db.DBContext{
		Queries: db.Connect(
			config.GetEnv("DB_HOST", "localhost"),
			config.GetEnv("DB_USER", "postgres"),
			config.GetEnv("DB_PASSWORD", "Postgres@WebomindApps"),
			config.GetEnv("DB_NAME", "postgres"),
			config.GetEnv("DB_PORT", "5432"),
			config.GetEnv("SSL_MODE", "disable"),
		),
	}
	fmt.Printf("Database running on %s", config.GetEnv("DB_HOST", "localhost"))
}

func startSearchEngine() {
	AppContext.Search = &search.TypesenseContext{
		SearchContext: search.ConnectTypesense(
			config.GetEnv("TYPESENSE_URL", "http://localhost:8108/"),
			config.GetEnv("TYPESENSE_API", "xyz"),
		),
	}
}

func SetupTable(ctx context.Context) {

	_, err := AppContext.Search.SearchContext.Collection("product_parts").Delete(ctx)

	if err != nil && !strings.Contains(err.Error(), "Not Found") {
		fmt.Println("No table to delete")
	} else if err == nil {
		log.Println("Deleted table")
	}

	schema := api.CollectionSchema{
		Name: "product_parts",
		Fields: []api.Field{
			{Name: "id", Type: "string"},
			{Name: "company", Type: "string"},
			{Name: "model", Type: "string"},
			{Name: "model_type", Type: "string"},
			{Name: "brand", Type: "string"},
			{Name: "category", Type: "string"},
			{Name: "part_no", Type: "string"},
		},
	}

	_, errSchema := AppContext.Search.SearchContext.Collections().Create(ctx, &schema)

	if errSchema != nil {
		log.Fatal("couldnt create schema %s", err.Error())
	}

	log.Println("Table Created")
}

func ExportData() {
	ctx := context.Background()
	SetupTable(ctx)
	rows, err := AppContext.DB.Queries.GetProductTableMetaData(ctx)

	if err != nil {
		fmt.Printf("couldt get data from database %s", err.Error())
	}

	log.Printf("length of table %d", len(rows))
	i := 0
	for _, row := range rows {

		doc := map[string]interface{}{
			"id":      row.ID.String(),
			"company": row.Company,
			"model":   row.Model,
			"model_type": func() string {
				if row.ModelType.Valid {
					return row.ModelType.String
				}

				return ""
			}(),
			"brand":    row.Brand,
			"category": row.Category,
			"part_no":  row.PartNo,
		}

		_, err := AppContext.Search.SearchContext.Collection("product_parts").Documents().Upsert(ctx, doc)
		if err != nil {
			log.Printf("Failed to push %s due to %s", row.PartNo, err.Error())
		} else {
			log.Printf("Pushed : %d", i)
		}
		i++
	}
}
