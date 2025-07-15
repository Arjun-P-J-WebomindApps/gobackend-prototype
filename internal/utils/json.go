package response

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func respondWithJSON(ctx *gin.Context,code int,payload interface{}){
	data,err := json.Marshal(payload)

	if err!=nil {
		log.Print("Failed to Marshall JSON response %s:",payload)
		ctx.Status(500)
		return
	}

	ctx.Header("Content-Type","application/json")
	ctx.Writer.WriteHeader(code)
	ctx.Writer.Write(data)
}