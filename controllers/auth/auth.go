package controllers

import (
	classUsersDBInteractions "backend/database/organizations"
	usersModel "backend/models/users"
	"os"
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
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))
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

// FetchLoggedInUserAdminStatus retrieves the logged-in user admin status
func FetchLoggedInUserAdminStatus(c echo.Context) bool {
	userID := FetchLoggedInUserID(c)
	class := c.QueryParam("selectedClass")
	classUser := classUsersDBInteractions.GetClassUserByUserIDAndClass(userID, class)
	return classUser.Admin
}