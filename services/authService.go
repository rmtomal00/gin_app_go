package authService

import (
	"fmt"
	"strings"

	passwordHash "app.team71.link/common"
	db "app.team71.link/dbConfig"
	authDto "app.team71.link/dto"
	userModel "app.team71.link/models"
)

func CreateUser(user *authDto.Register) (interface{}, error){

	mp := make(map[string]interface{})


	hashPass, err := passwordHash.HashPassword(user.Password)

	if err != nil{
		mp["message"] = "Password Invalid"
		return mp, fmt.Errorf("Password Invalid")
	}


	email := strings.TrimSpace(strings.ToLower(user.Email))
	
	

	userData := userModel.User{
		Username: user.Username,
		Email: email,
		Password: hashPass,
	}

	result := db.ConnectToDb().Save(&userData)
	
	if result.Error != nil {
		mp["message"] = "Unsuccess";
		return mp, result.Error
	}
	
	mp["message"] = "Success";
	mp["userId"] = userData.ID
	return mp, nil
}