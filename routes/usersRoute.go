package route

import (
	"app.team71.link/controller"
	"app.team71.link/middleware"
	"github.com/gin-gonic/gin"
)

func UsersRouter(r *gin.Engine) {
	
	route := r.Group("/api/v1/users")

	route.Use(middleware.Usersmiddleware());
	{
		route.GET("/user-profile", controller.UserProfile)
	}

}