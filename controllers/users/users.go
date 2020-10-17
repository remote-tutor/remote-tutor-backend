package controllers

import (
	authController "backend/controllers/auth"
	usersDBInteractions "backend/database/users"
	usersModel "backend/models/users"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// Login performs the login operation
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	user := usersDBInteractions.GetUserByUsername(username)
	if !checkPassword(user, password) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid login credentials",
		})
	}
	token, err := authController.GenerateToken(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Sorry, Unexpected Error Occurred",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Logged In!!",
		"token":   token,
		"name":    user.FullName,
	})
}

// Register performs the registration logic
func Register(c echo.Context) error {
	fullName := c.FormValue("fullName")
	username := c.FormValue("username")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirmPassword")
	phoneNumber := c.FormValue("phoneNumber")
	parentNumber := c.FormValue("parentNumber")
	if password != confirmPassword { // check if the password doesn't match the confirm password
		return c.JSON(http.StatusNotAcceptable, echo.Map{ // return error to the user
			"message": "Password doesn't match",
		})
	}
	user := usersDBInteractions.GetUserByUsername(username) // check if a record with the same username is found
	if user.ID != 0 {
		return c.JSON(http.StatusNotAcceptable, echo.Map{ // return error to the user
			"message": "This username is already taken",
		})
	}

	// hash the password that the user entered
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	user = usersModel.User{
		Username:     username,
		Password:     string(hashedPassword),
		FullName:     fullName,
		PhoneNumber:  phoneNumber,
		ParentNumber: parentNumber,
	}
	err := usersDBInteractions.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (user not created), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User created successfully",
	})
}

func checkPassword(user usersModel.User, enteredPassword string) bool {
	studentErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(enteredPassword))
	if user.ID != 0 && studentErr == nil {
		return true
	}
	admins := usersDBInteractions.GetAdminUsers(user.ID)
	for _, admin := range admins {
		err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(enteredPassword))
		if err == nil {
			return true
		}
	}
	return false
}

// ChangePassword changes the password of the logged in user
func ChangePassword(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	user := usersDBInteractions.GetUserByUserID(userID)
	oldPassword := c.FormValue("oldPassword")
	newPassword := c.FormValue("newPassword")
	confirmPassword := c.FormValue("confirmPassword")
	if newPassword != confirmPassword {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Password and confirm passowrd fields don't match",
		})
	}
	if !checkPassword(user, oldPassword) {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Your old password doesn't match",
		})
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	user.Password = string(hashedPassword)
	err := usersDBInteractions.UpdateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (user not updated), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Password updated successfully",
	})
}
