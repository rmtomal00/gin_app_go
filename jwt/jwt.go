package jwtTokenManageer

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type JwtMapData struct{
	Email string
	ID int
	EXPIRE int64
}

var jwtSecret = []byte("Rmtomal10@")

func GenerateJWT(jwtM JwtMapData) (string, error) {
	claims := jwt.MapClaims{
		"email": jwtM.Email,
		"id": jwtM.ID,
		"exp": jwtM.EXPIRE,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims);

	return token.SignedString(jwtSecret)
}

func ValidJwtToken(unSolveToken string) (*JwtMapData, error){
	nullSta := JwtMapData{}
	token, err := jwt.Parse(unSolveToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil{
		return &nullSta, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		exp, ok := claims["exp"].(float64)
		if !ok {
			return &nullSta, fmt.Errorf("Invalid expiration format")
		}

		if int64(exp) < time.Now().Unix() {
			return &nullSta, fmt.Errorf("Token expired")
		}

		nullSta := &JwtMapData{
			Email: claims["email"].(string),
			ID: int(claims["id"].(float64)),
			EXPIRE: int64(claims["exp"].(float64)),
		}

		return nullSta, nil
	}

	return nil, nil
}

