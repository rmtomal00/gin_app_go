package controller

import (
	model "app.team71.link/models"
	response "app.team71.link/responseStruct"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func UserProfile(ctx *gin.Context) {
	userData, ok := ctx.Get("userData");

	if !ok {
		response.BadRequest(ctx, 400, "Data not found in context")
		return
	}

	var user model.User
	if err := mapstructure.Decode(userData, &user); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to decode user data"})
		return
	}

	response.Success(ctx, 200, user, "Success")
}