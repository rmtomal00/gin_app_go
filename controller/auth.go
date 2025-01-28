package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	authDto "app.team71.link/dto"
	jwtTokenManageer "app.team71.link/jwt"
	mailSetup "app.team71.link/mailConfig"
	userModel "app.team71.link/models"
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

	userData := create.(userModel.User);
	token, err := jwtTokenManageer.GenerateJWT(jwtTokenManageer.JwtMapData{
		Email: strings.ToLower(register.Email), 
		ID: int(userData.ID), 
		EXPIRE: time.Now().Add(24 * time.Hour).Unix(),
	});

	link := fmt.Sprintf("http://102.168.2.106:8080/api/v1/auth/confirm-email/%s", token);

	if err != nil {
		ctx.JSON(http.StatusOK, response.Success(200, create, "Success but you can't get the email for verificatio"));
		return;
	}

	userData.TempToken = token;

	go authService.SaveUserData(&userData)

	
	go mailSetup.SendMail2(register.Email, "Email confirm", link);
	ctx.JSON(http.StatusOK, response.Success(200, map[string]interface{}{"userId": userData.ID}, "Success"));
}


func 