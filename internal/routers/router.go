package routers

import (
	"github.com/gin-gonic/gin"
)



func NewRouter() *gin.Engine{
	router := gin.Default()
	
	// userRoutes:=router.Group("/user")
	// {
	// 	// userRoutes.GET("/",handlers.GetUsers)
	// 	// userRoutes.GET("/:id",handlers.GetUser)
	// 	// userRoutes.POST("/",handlers.CreateUser)
	// 	// userRoutes.PUT("/:id",handlers.UpdateUser)
	// 	// userRoutes.DELETE("/:id",handlers.DeleteUser)
	// }
	

	return router
}