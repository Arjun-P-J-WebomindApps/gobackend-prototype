package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Defines a function that gives the output in required format
func RespondWithJSON(ctx *gin.Context, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to Marshall JSON response %s:", payload)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Data(code, "application/json", data)
}

func RespondWithError(ctx *gin.Context, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX Error")
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(ctx, code, errResponse{
		Error: msg,
	})
}
