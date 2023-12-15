package user_model

import (
	"os"
	"time"

	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	"github.com/golang-jwt/jwt"
)

const (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_errors.RestError) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_errors.NewForbiddenError("Error to generate token")
	}

	return tokenString, nil
}
