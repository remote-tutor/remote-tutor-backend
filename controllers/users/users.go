package controllers

import (
	authController "backend/controllers/auth"
	usersDBInteractions "backend/database/users"
	usersModel "backend/models/users"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// Login performs the login operation
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	user := usersDBInteractions.GetUserByUsername(username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 || err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid login credentials",
		})
	}
	if !user.Activated {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Sorry, you haven't been verified yet",
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
		"admin":   user.Admin,
		"name":    user.FullName,
	})
}

// Register performs the registration logic
func Register(c echo.Context) error {
	fullName := c.FormValue("fullName")
	username := c.FormValue("username")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirmPassword")
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
		Username: username,
		Password: string(hashedPassword),
		FullName: fullName,
	}
	usersDBInteractions.CreateUser(&user)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User created successfully",
	})
}

// GetPendingUsers retrieves the non activated users to view to the admin
func GetPendingUsers(c echo.Context) error {
	queryParams := c.Request().URL.Query()
	sortDesc := utils.ConvertToBoolArray(queryParams["sortDesc[]"])
	sortBy := queryParams["sortBy[]"]
	page := utils.ConvertToInt(queryParams["page"][0])
	itemsPerPage := utils.ConvertToInt(queryParams["itemsPerPage"][0])
	searchByValue := c.QueryParam("searchByValue")
	searchByField := c.QueryParam("searchByField")

	pendingUsers := usersDBInteractions.GetPendingUsers(sortBy, sortDesc, page, itemsPerPage, searchByValue, searchByField)
	totalPendingUsers := usersDBInteractions.GetTotalNumberOfPendingUsers(searchByValue, searchByField)

	return c.JSON(http.StatusOK, echo.Map{
		"pendingStudents":      pendingUsers,
		"totalPendingStudents": totalPendingUsers,
	})
}

// UpdateUser updates the user in the database
func UpdateUser(c echo.Context) error {
	userID := utils.ConvertToUInt(c.FormValue("userID"))
	fullName := c.FormValue("fullName")
	year := utils.ConvertToInt(c.FormValue("year"))
	if fullName == "" || year < 1 || year > 3 {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"userID":  userID,
			"message": "Error while saving the data, make sure you entered a correct name and/or year",
		})
	}
	status := utils.ConvertToInt(c.FormValue("status"))
	user := usersDBInteractions.GetUserByUserID(userID)
	user.FullName = fullName
	if status == 1 {
		user.Activated = true
		user.Admin = true
	} else if status == 0 {
		user.Activated = true
	}
	usersDBInteractions.UpdateUser(&user)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User updated successfully",
	})
}

// CheckUserIsAdmin checks whether the user has admin rights or not
func CheckUserIsAdmin(c echo.Context) error {
	userid := uint(1)
	user := usersDBInteractions.GetUserByUserID(userid)
	isAdmin := false
	if user.Admin {
		isAdmin = true
	}
	return c.JSON(http.StatusOK, echo.Map{
		"admin": isAdmin,
	})
}
