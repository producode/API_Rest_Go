package services

import (
	"github.com/apirestgo/models"
	"github.com/apirestgo/repository"

	"strings"
)

func Authenticate(token string) (models.User, error) {
	token_string := strings.Replace(token, "Bearer ", "", 1)
	user, err := repository.GetUserByToken(token_string)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
