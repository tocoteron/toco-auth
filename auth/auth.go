package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/tocoteron/toco-auth/model"
)

func generateUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func GenerateToken(serverSetting *model.ServerSetting, user *model.User) (string, error) {
	jwtID, err := generateUUID()
	if err != nil {
		return "", err
	}

	fmt.Println(serverSetting)

	curerntTime := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":       serverSetting.Identifier,
		"sub":       serverSetting.Identifier,
		"aud":       serverSetting.Identifier,
		"exp":       curerntTime.Add(time.Hour * 24).Unix(),
		"nbf":       curerntTime.Unix(),
		"iat":       curerntTime.Unix(),
		"jti":       jwtID,
		"user_id":   user.ID,
		"user_name": user.Name,
	})

	tokenString, err := token.SignedString([]byte(serverSetting.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
