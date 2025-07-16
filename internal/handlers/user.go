package handlers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func GetUsers(ctx *gin.Context){
// 	ctx.JSON(http.StatusOK,gin.H{"message":"Get All Users"})
// }

// func GetUser(ctx *gin.Context){
// 	id:=ctx.Param("id")
// 	ctx.JSON(http.StatusOK,gin.H{"message":"Get User by Id","id":id})
// }

// func CreateUser(ctx *gin.Context){
// 	ctx.JSON(http.StatusCreated,gin.H{"message":"Create a New user"})
// }

// func UpdateUser(ctx *gin.Context){
// 	id:=ctx.Param("id")
// 	ctx.JSON(http.StatusOK,gin.H{"message":"Updated the user","id":id})
// }

// func DeleteUser(ctx *gin.Context){
// 	id:=ctx.Param("id")
// 	ctx.JSON(http.StatusOK,gin.H{"message":"Delete a user","id":id})
// }