package main

import (
	"fmt"
	"net/http"

	response "app.team71.link/responseStruct"
	authRoute "app.team71.link/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	route := gin.Default()


	route.Use(ErrorHandlerMiddleware());

	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file\n")
	}

	route.GET("/", func(ctx *gin.Context) {
		res := response.Success(200, "", "Success")

		ctx.JSON(http.StatusOK, res)
	})

	//auth controller
	authRoute.AuthRouteFunc(route);

	
	//error handle
	route.Use(func(ctx *gin.Context) {
		res := response.BadRequest(404, "Not found")

		ctx.JSON(http.StatusForbidden, res)
	})

	route.Run("localhost:8080")
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                c.JSON(500, response.ServerError(fmt.Sprintf("Internal Server Error: %v", err)))
                c.Abort()
            }
        }()
        c.Next()
    }
}