package search

import "github.com/typesense/typesense-go/typesense"

type TypesenseContext struct {
	SearchContext *typesense.Client
}

func ConnectTypesense(server string, apiKey string) *typesense.Client {
	client := typesense.NewClient(typesense.WithServer(server), typesense.WithAPIKey(apiKey))

	return client
}
