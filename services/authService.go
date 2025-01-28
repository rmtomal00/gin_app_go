package service

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

	result := db.ConnectToDb().Create(&userData)
	
	if result.Error != nil {
		mp["message"] = "Unsuccess";
		return mp, result.Error
	}
	
	
	return userData, nil
}

func SaveUserData(user *userModel.User) (string, error){
	result := db.ConnectToDb().Save(&user);

	if result.Error != nil {
		return result.Error.Error(), result.Error
	}

	return "Update Data Successfully", nil
}

func FindBy(data map[string]interface{}) (*userModel.User, error){
	var user userModel.User
	result := db.ConnectToDb().Where(data).First(&user);
	if result.Error != nil {
		return &user, result.Error
	}
	return &user, nil
}