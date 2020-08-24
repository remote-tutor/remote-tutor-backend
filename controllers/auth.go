package controllers

import (
	md "backend/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateToken(user *md.User) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["admin"] = user.Admin
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}
