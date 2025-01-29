package controller

import (
	"fmt"
	"os"
	"strings"
	"time"

	"app.team71.link/common"
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
		response.BadRequest(ctx, 400, err.Error())
		return
	}

	create, err := authService.CreateUser(&register)

	if err != nil {
		response.BadRequest(ctx, 400, err.Error())
		return
	}

	userData := create.(userModel.User);
	token, err := jwtTokenManageer.GenerateJWT(jwtTokenManageer.JwtMapData{
		Email: strings.ToLower(register.Email+"Register"), 
		ID: int(userData.ID), 
		EXPIRE: time.Now().Add(24 * time.Hour).Unix(),
	});

	link := fmt.Sprintf("%s/api/v1/auth/confirm-email/%s",os.Getenv("Base_url"), token);

	if err != nil {
		response.Success(ctx, 200, create, "Success but you can't get the email for verification")
		return;
	}

	userData.TempToken = token;

	go authService.SaveUserData(&userData)

	
	go mailSetup.SendMail2(register.Email, "Email confirm", link);
	response.Success(ctx, 200, map[string]interface{}{"userId": userData.ID}, "Successfull, Please comfirm you email. Check your inbox")
}


func ConfirmEmail(context *gin.Context){

	token := context.Param("token")
	strings.TrimSpace(token)
	if token == ""{
		response.BadRequest(context, 400, "Token not found");
		return;
	}

	tokenData, err := jwtTokenManageer.ValidJwtToken(strings.TrimSpace(token)); 
	
	if err != nil{
		response.BadRequest(context, 400, err.Error())
		return;
	}

	dataFind := map[string]interface{}{"id": tokenData.ID}

	userDataGetFromDb, err := authService.FindBy(dataFind);

	if err != nil{
		response.BadRequest(context, 400, err.Error())
		return;
	}

	if userDataGetFromDb.TempToken != token {
		response.BadRequest(context, 400, "Invalid link")
		return;
	}

	userDataGetFromDb.TempToken = "";

	userDataGetFromDb.EmailVerify = true;

	msg , err := authService.SaveUserData(userDataGetFromDb);

	if err != nil {
		fmt.Println(err)
		response.BadRequest(context, 400, "Email not verify, due to internal error")
		return;
	}
	_ = msg;

	response.Success(context, 200,"", "Email verification successfully done")
}

func ResendVerificationEmail(ctx *gin.Context){
	email := ctx.Query("email")
	if email == "" {
		response.BadRequest(ctx, 400, "Email can't be empty")
		return
	}
	email = strings.ToLower(strings.TrimSpace(email));

	user, err := authService.FindBy(map[string]interface{}{"email": email});

	if checkErr(err) {
		response.BadRequest(ctx, 400, err.Error())
		return
	}

	if user.EmailVerify || !user.UserDisable{
		response.BadRequest(ctx, 400, "User already verified or Account is disable")
		return
	}

	token, err := jwtTokenManageer.GenerateJWT(jwtTokenManageer.JwtMapData{
		Email: strings.ToLower(email+"ResetPassword"), 
		ID: int(user.ID), 
		EXPIRE: time.Now().Add(24 * time.Hour).Unix(),
	});

	link := fmt.Sprintf("%s/api/v1/auth/confirm-email/%s",os.Getenv("Base_url"), token);

	if err != nil {
		response.Success(ctx, 200, "", "Success but you can't get the email for verification")
		return;
	}

	user.TempToken = token;

	go authService.SaveUserData(user)

	
	go mailSetup.SendMail2(user.Email, "Email confirm", link);
	response.Success(ctx, 200, "", "Please check your email")

}

func ResetPassword(ctx* gin.Context){
	var resetPass authDto.ResetPassword;

	err := ctx.ShouldBind(&resetPass)

	if err != nil {
		response.BadRequest(ctx, 400, err.Error())
		return;
	}

	pass, err := common.HashPassword(resetPass.Password);

	if err != nil {
		fmt.Println(err)
		response.BadRequest(ctx, 400, "This password you can't use")
		return
	}

	user, err := authService.FindBy(map[string]interface{}{"email": strings.ToLower(strings.TrimSpace(resetPass.Email))});

	if err != nil {
		fmt.Println(err)
		response.BadRequest(ctx, 400, "Email not found")
		return
	}

	if user.UserDisable || !user.EmailVerify {
		response.BadRequest(ctx, 400, "User account disbale or email not verified")
		return
	}

	token, err := jwtTokenManageer.GenerateJWT(jwtTokenManageer.JwtMapData{Email: pass, ID: int(user.ID), EXPIRE: time.Now().Add(24 * time.Hour).Unix()});

	if err != nil {
		fmt.Println(err)
		response.BadRequest(ctx, 400, err.Error())
		return
	}

	user.TempToken = token;

	link := fmt.Sprintf("%s/api/v1/auth/confirm-reset-password/%s", os.Getenv("Base_Url"), token)
	go mailSetup.SendMail(strings.ToLower(strings.TrimSpace(resetPass.Email)), "Reset Password", link)

	response.Success(ctx, 200, "", "Please check your email, We send a link");
}

func ConfirmResetPassword(ctx* gin.Context){
	token := ctx.Param("token");
	if token == "" {
		response.BadRequest(ctx, 400, "Token can't be empty");
		return;
	}
	token = strings.TrimSpace(token);

	tokenData, err := jwtTokenManageer.ValidJwtToken(token);

	if err != nil {
		response.BadRequest(ctx, 400, err.Error())
		return
	}

	user, err := authService.FindBy(map[string]interface{}{"id": tokenData.ID});
	if checkErr(err) {
		response.BadRequest(ctx, 400, err.Error());
		return;
	}

	user.Password = tokenData.Email;

	upadate, err := authService.SaveUserData(user);

	if checkErr(err) {
		response.BadRequest(ctx, 400, err.Error())
		return;
	}
	_ = upadate;

	response.Success(ctx, 200, "", "Password update success");

}

func Login(ctx* gin.Context){
	var loginData authDto.Login;
	err := ctx.ShouldBind(&loginData)

	if checkErr(err) {
		response.BadRequest(ctx, 400, err.Error())
		return
	}

	user, err := authService.FindBy(map[string]interface{}{"email": strings.ToLower(strings.TrimSpace(loginData.Email))})
	if checkErr(err) {
		response.BadRequest(ctx, 400, "Email not found")
		return
	}

	if !common.IsValidPass(user.Password, loginData.Password) {
		response.BadRequest(ctx, 400, "Password not match")
		return
	}

	if !user.EmailVerify {
		response.BadRequest(ctx, 400, "Email not varified")
		return
	}

	if user.UserDisable {
		response.BadRequest(ctx, 400, "Account is disable")
		return
	}

	token, err := jwtTokenManageer.GenerateJWT(jwtTokenManageer.JwtMapData{Email: user.Email, ID: int(user.ID), EXPIRE: time.Now().Add(24 * time.Hour).Unix()})

	user.LoginToken = token;

	u, err := authService.SaveUserData(user)

	_ = u;

	if checkErr(err) {
		response.BadRequest(ctx, 400, "Login fail")
		return
	}

	response.Success(ctx, 200, map[string]interface{}{"id": user.ID, "token": token}, "Login successfull")

}


func checkErr(err error) bool{
	if err != nil {
		return true
	}
	return false
}

func CkeckToken(ctx *gin.Context){
	token := ctx.Param("token")

	if token == "" {
		response.BadRequest(ctx, 400, "Token can't be empty")
		return
	}

	d, err := jwtTokenManageer.ValidJwtToken(token); 

	if err != nil{
		response.BadRequest(ctx, 400, err.Error())
		return
	}

	user, err := authService.FindBy(map[string]interface{}{"id": d.ID})
	if checkErr(err) || user.LoginToken != strings.TrimSpace(token) {
		response.BadRequest(ctx, 400, "Token not match")
		return
	}

	response.Success(ctx, 200, map[string]bool{"isValid": true}, "Token is valid")
}