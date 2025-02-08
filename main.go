package main

import (
	"fmt"
	"time"

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
		time.Sleep(100 * time.Millisecond)
		response.Success(ctx, 200, "", "Success")
		return
	})

	//auth controller
	authRoute.AuthRouteFunc(route);

	
	//error handle
	route.Use(func(ctx *gin.Context) {
		response.BadRequest(ctx, 404, "Not found")
		return
	})

	route.Run(":8080")
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                response.ServerError(c, fmt.Sprintf("Internal Server Error: %v", err))
                c.Abort()
            }
        }()
        c.Next()
    }
}