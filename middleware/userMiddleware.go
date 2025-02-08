package middleware

import (
	"strings"

	db "app.team71.link/dbConfig"
	jwtTokenManageer "app.team71.link/jwt"
	model "app.team71.link/models"
	response "app.team71.link/responseStruct"
	"github.com/gin-gonic/gin"
)

func Usersmiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		token := ctx.GetHeader("Authorization");

		token1, found := strings.CutPrefix(token, "Bearer ");
		if !found {
			response.BadRequest(ctx, 401, "Token is missing")
			return;
		}

		data, err := jwtTokenManageer.ValidJwtToken(strings.TrimSpace(token1));
		if err != nil {
			response.BadRequest(ctx, 401, err.Error())
			return;
		}
		var user model.User
		userData := db.ConnectToDb().Where(model.User{ID: int64(data.ID)}).First(&user);
		if userData.Error != nil {
			response.BadRequest(ctx, 401, userData.Error.Error())
			return
		}

		if user.UserDisable {
			response.BadRequest(ctx, 401, "Your Account is disabled")
			return
		}

		if user.LoginToken != strings.TrimSpace(token1) {
			response.BadRequest(ctx, 401, "Token is old. You have login with another device")
			return
		}

		ctx.Set("userData", user);

		ctx.Next();
	}
}

