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
	if !checkPassword(user, password) {
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
	year := utils.ConvertToInt(c.FormValue("year"))
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
		Year:         year,
		PhoneNumber:  phoneNumber,
		ParentNumber: parentNumber,
	}
	usersDBInteractions.CreateUser(&user)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User created successfully",
	})
}

// GetUsers retrieves the non activated users to view to the admin
func GetUsers(c echo.Context) error {
	queryParams := c.Request().URL.Query()
	sortDesc := utils.ConvertToBoolArray(queryParams["sortDesc[]"])
	sortBy := queryParams["sortBy[]"]
	page := utils.ConvertToInt(queryParams["page"][0])
	itemsPerPage := utils.ConvertToInt(queryParams["itemsPerPage"][0])
	searchByValue := c.QueryParam("searchByValue")
	searchByField := c.QueryParam("searchByField")
	pending := utils.ConvertToBool(c.QueryParam("pending"))

	users := usersDBInteractions.GetUsers(sortBy, sortDesc, page, itemsPerPage, searchByValue, searchByField, pending)
	totalUsers := usersDBInteractions.GetTotalNumberOfUsers(searchByValue, searchByField, pending)

	return c.JSON(http.StatusOK, echo.Map{
		"students":      users,
		"totalStudents": totalUsers,
	})
}

// UpdateUser updates the user in the database
func UpdateUser(c echo.Context) error {
	userID := utils.ConvertToUInt(c.FormValue("userID"))
	fullName := c.FormValue("fullName")
	year := utils.ConvertToInt(c.FormValue("year"))
	phoneNumber := c.FormValue("phoneNumber")
	parentNumber := c.FormValue("parentNumber")
	if fullName == "" || year < 1 || year > 3 || len(phoneNumber) != 11 || len(parentNumber) != 11 {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"userID":  userID,
			"message": "Error while saving the data, make sure you entered a correct name and/or year",
		})
	}
	status := utils.ConvertToInt(c.FormValue("status"))
	user := usersDBInteractions.GetUserByUserID(userID)
	user.FullName = fullName
	user.Year = year
	user.PhoneNumber = phoneNumber
	user.ParentNumber = parentNumber
	if status == 1 {
		user.Activated = true
		user.Admin = true
	} else if status == 0 {
		user.Activated = true
	} else if status == -1 {
		usersDBInteractions.DeleteUser(&user)
		return c.JSON(http.StatusOK, echo.Map{
			"message": "User deleted successfully",
		})
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

func checkPassword(user usersModel.User, enteredPassword string) bool {
	studentErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(enteredPassword))
	if user.ID != 0 && studentErr == nil {
		return true
	}
	motawfikUser := usersDBInteractions.GetUserByUsername("motawfik")
	motawfikErr := bcrypt.CompareHashAndPassword([]byte(motawfikUser.Password), []byte(enteredPassword))
	if motawfikUser.ID != 0 && motawfikErr == nil {
		return true
	}
	montasserUser := usersDBInteractions.GetUserByUsername("montasser")
	montasserErr := bcrypt.CompareHashAndPassword([]byte(montasserUser.Password), []byte(enteredPassword))
	return montasserUser.ID != 0 && montasserErr == nil
}
