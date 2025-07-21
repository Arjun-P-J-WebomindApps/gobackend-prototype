package routers

import (
	"slices"
	"strconv"
	"strings"

	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/app_context"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/handlers"
	"github.com/gin-gonic/gin"
)

type CORS struct {
	AllowedOrigins   []string
	AllowedHeaders   []string
	AllowedMethods   []string
	AllowCredentials bool
}

func NewRouter(appCtx *app_context.Context) *gin.Engine {
	router := gin.Default()
	cors := CORS{
		AllowedOrigins: []string{"http://localhost:8000"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowedMethods: []string{"POST", "GET"},
	}
	router.Use(corsMiddleware(cors))

	graphql := &handlers.GraphQL{
		AppCtx: appCtx,
	}
	graphql.CreateGraphQLHandler()
	graphql.CreatePlaygroundHandler()

	router.POST("/query", graphql.GraphQlHandler)
	router.GET("/playground", graphql.PlaygroundHandler)

	return router
}

func corsMiddleware(cors CORS) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		isOriginAllowed := func(origin string, allowedOrigins []string) bool {
			return slices.Contains(allowedOrigins, origin)
		}

		origin := ctx.Request.Header.Get("Origin")

		if isOriginAllowed(origin, cors.AllowedOrigins) {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			ctx.Writer.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(cors.AllowCredentials))
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(cors.AllowedHeaders, ","))
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(cors.AllowedMethods, ","))
		}

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
