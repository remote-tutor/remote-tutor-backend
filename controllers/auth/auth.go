package controllers

import (
	usersModel "backend/models/users"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GenerateToken(user *usersModel.User) (string, error) {
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

// FetchLoggedInUserID retrieves the logged-in user's ID
func FetchLoggedInUserID(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := uint(claims["id"].(float64))
	return id
}
