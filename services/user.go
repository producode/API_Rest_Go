package services

import (
	"github.com/apirestgo/models"
	"github.com/apirestgo/utils"
	"github.com/dgrijalva/jwt-go"

	"strings"
)

func LogIn() (string, error) {
	user := models.User{
		Id:       1,
		Username: "admin",
		Password: "admin",
	}
	token := utils.Generate_JWT(user)
	return token, nil
}

func Authenticate(token string) (int, error) {
	token = strings.Replace(token, "Bearer ", "", 1)
	claims := &models.UserClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, err
	}
	if !tkn.Valid {
		return 0, err
	}
	return int(claims.User_id), nil
}
