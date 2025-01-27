package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	authDto "app.team71.link/dto"
	jwtTokenManageer "app.team71.link/jwt"
	mailSetup "app.team71.link/mailConfig"
	response "app.team71.link/responseStruct"
	authService "app.team71.link/services"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	

	var register authDto.Register

	if err := ctx.ShouldBindJSON(&register); err != nil{
		fmt.Print("error: ",err, "\n")
		ctx.JSON(http.StatusBadRequest, response.BadRequest(400, err.Error()))
		return
	}

	create, err := authService.CreateUser(&register)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BadRequest(400, err.Error()))
		return
	}

	token, err := jwtTokenManageer.GenerateJWT(jwtTokenManageer.JwtMapData{
		Email: strings.ToLower(register.Email), 
		ID: int(create["userId"].(int64)), 
		EXPIRE: time.Now().Add(24 * time.Hour).Unix(),
	});

	if err != nil {
		ctx.JSON(http.StatusOK, response.Success(200, create, "Success but you can't get the email for verificatio"));
		return;
	}

	
	go mailSetup.SendMail(register.Email, "Email confirm", token);
	ctx.JSON(http.StatusOK, response.Success(200, create, "Success"));
}