package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/context"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/graph"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/graph/resolvers"
)

type GraphQL struct {
	AppCtx     *context.Context
	handler    *handler.Server
	playground *http.HandlerFunc
}

func (graphql *GraphQL) CreateGraphQLHandler() {
	graphql.handler = handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{DB: graphql.AppCtx.DB}}))

	graphql.handler.AddTransport(transport.Options{})
	graphql.handler.AddTransport(transport.GET{})
	graphql.handler.AddTransport(transport.POST{})

	graphql.handler.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	graphql.handler.Use(extension.Introspection{})
	graphql.handler.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
}

// Defines the Handler for setting up graphql endpoints
func (graphql *GraphQL) GraphQlHandler(ctx *gin.Context) {

	graphql.handler.ServeHTTP(ctx.Writer, ctx.Request)
}

func (graphql *GraphQL) CreatePlaygroundHandler() {
	handler := playground.Handler("GraphQL playground", "/query")
	graphql.playground = &handler
}

// Defines the Hnalder fro the Playground
func (graphql *GraphQL) PlaygroundHandler(ctx *gin.Context) {

	graphql.playground.ServeHTTP(ctx.Writer, ctx.Request)
}
