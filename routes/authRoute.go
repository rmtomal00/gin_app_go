package route

import (
	auth "app.team71.link/controller"
	"github.com/gin-gonic/gin"
)

func AuthRouteFunc(ctx *gin.Engine) {

	group := ctx.Group("/api/v1/auth")

	group.POST("/register", auth.Register)
	group.GET("/confirm-email/:token", auth.ConfirmEmail)
	group.POST("/login", auth.Login)
	group.POST("/reset-password", auth.ResetPassword)
	group.Any("/confirm-reset-password/:token", auth.ConfirmResetPassword)
	group.Any("/resend-verification-email", auth.ResendVerificationEmail)
	group.Any("/check-token/:token", auth.CkeckToken)
}