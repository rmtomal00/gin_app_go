package authRoute

import (
	auth "app.team71.link/controller"
	"github.com/gin-gonic/gin"
)

func AuthRouteFunc(ctx *gin.Engine) {

	group := ctx.Group("/api/v1/auth")

	group.POST("/register", auth.Register)
	
}