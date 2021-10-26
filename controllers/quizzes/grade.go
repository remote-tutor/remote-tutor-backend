package quizzes

import (
	authController "backend/controllers/auth"
	classesDBInteractions "backend/database/organizations"
	quizzesDBInteractions "backend/database/quizzes"
	quizzesModel "backend/models/quizzes"
	usersModel "backend/models/users"
	gradesPDFHandler "backend/pdf/handlers/quizzes"
	"backend/utils"
	"github.com/jinzhu/now"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// GetGradesByQuiz Fetches logged-in user's grade for a quiz
func GetGradesByQuiz(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	quizID := utils.ConvertToUInt(c.QueryParam("quizID"))

	quizGrade := quizzesDBInteractions.GetGradesByQuizID(userID, quizID)
	return c.JSON(http.StatusOK, echo.Map{
		"quizGrade": []quizzesModel.QuizGrade{quizGrade},
	})
}

// CreateQuizGrade creates a quiz grade record for the logged-in user
func CreateQuizGrade(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	quizID := utils.ConvertToUInt(c.FormValue("quizID"))
	quiz := quizzesDBInteractions.GetQuizByID(quizID)
	validTill := getSmallestDate(quiz.EndTime, time.Now().Add(time.Duration(quiz.StudentTime)*time.Minute))
	quizGrade := quizzesModel.QuizGrade{
		Grade:     0,
		QuizID:    quizID,
		UserID:    userID,
		StartAt:   time.Now(),
		ValidTill: validTill,
	}
	err := quizzesDBInteractions.CreateGrade(&quizGrade)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred, please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"validTill": quizGrade.ValidTill,
	})
}

func GetStudentRemainingTime(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	quizID := utils.ConvertToUInt(c.QueryParam("quizID"))

	remainingTime, recordFound := quizzesDBInteractions.GetStudentRemainingTime(userID, quizID)
	return c.JSON(http.StatusOK, echo.Map{
		"studentTime": remainingTime,
		"recordFound": recordFound,
	})
}

// GetGradesByMonthAndUser fetches the grades of a specific user for a specific month
func GetGradesByMonthAndUser(c echo.Context) error {
	month := utils.ConvertToTime(c.QueryParam("month"))
	startOfMonth := now.With(month).BeginningOfMonth()
	endOfMonth := now.With(month).EndOfMonth()
	class := c.QueryParam("selectedClass")
	userID := authController.FetchLoggedInUserID(c)
	quizzes, grades := quizzesDBInteractions.GetGradesByMonthAndUser(class, userID, startOfMonth, endOfMonth)
	organizedGrades, quizzesTotalMarks, _ := organizeGrades(quizzes, grades)
	return c.JSON(http.StatusOK, echo.Map{
		"quizzes":           quizzes,
		"quizzesTotalMarks": quizzesTotalMarks,
		"grades":            organizedGrades,
	})
}

func GetGradesByMonthForAllUsers(c echo.Context) error {
	month := utils.ConvertToTime(c.QueryParam("month"))
	startOfMonth := now.With(month).BeginningOfMonth()
	endOfMonth := now.With(month).EndOfMonth()
	class := c.QueryParam("selectedClass")
	quizzes, grades := quizzesDBInteractions.GetGradesByMonthForAllUsers(class, startOfMonth, endOfMonth)
	organizedGrades, quizzesTotalMarks, _ := organizeGrades(quizzes, grades)
	return c.JSON(http.StatusOK, echo.Map{
		"quizzes":           quizzes,
		"quizzesTotalMarks": quizzesTotalMarks,
		"grades":            organizedGrades,
	})
}

func organizeGrades(quizzes []quizzesModel.Quiz, grades []quizzesModel.QuizGrade) ([]map[string]interface{}, int, [][]int) {
	quizzesIDsMap := make(map[uint]int)  // map to hold the index of each quizID
	studentsIDsMap := make(map[uint]int) // map to hold the index of each studentID
	quizzesTotalMarks := 0
	for index, quiz := range quizzes {
		quizzesIDsMap[quiz.ID] = index      // store each ID to the corresponding index in the 2D array
		quizzesTotalMarks += quiz.TotalMark // add to the quizzes total marks
	}
	studentIndex := 0 // index of UNIQUE users
	studentsArr := make([]usersModel.User, 0)
	for _, quizGrade := range grades { // loop over the grades
		_, isFound := studentsIDsMap[quizGrade.UserID] // check if the userID has been seen before
		if !isFound {                                  // if NOT
			studentsIDsMap[quizGrade.UserID] = studentIndex   // store its index
			studentsArr = append(studentsArr, quizGrade.User) // add the user to the users array
			studentIndex++                                    // increment the unique users index
		}
	}
	grades2DArray := make([][]int, studentIndex) // 2D array to store grades
	for arr := range grades2DArray {
		grades2DArray[arr] = make([]int, len(quizzes)) // initialize each inner array
	}
	for _, quizGrade := range grades {
		currentStudentIndex := studentsIDsMap[quizGrade.UserID]                // get the index of the student
		currentQuizIndex := quizzesIDsMap[quizGrade.QuizID]                    // get the index of the quiz
		grades2DArray[currentStudentIndex][currentQuizIndex] = quizGrade.Grade // store the quizGrade
	}
	returnedObject := make([]map[string]interface{}, len(grades2DArray)) // array of custom objects to be returned to the frontend
	for index, student := range studentsArr {
		returnedObject[index] = map[string]interface{}{} // initialize the current object
		returnedObject[index]["user"] = student          // add student object to the returned value
	}
	for index := range grades2DArray {
		totalMarks := 0                                              // calculate total marks for each user
		for gradeIndex, studentGrade := range grades2DArray[index] { // add each quizGrade under its correct quizID
			returnedObject[index][strconv.Itoa(int(quizzes[gradeIndex].ID))] = 0
			returnedObject[index][strconv.Itoa(int(quizzes[gradeIndex].ID))] = studentGrade
			totalMarks += studentGrade
		}
		returnedObject[index]["total"] = totalMarks
	}
	return returnedObject, quizzesTotalMarks, grades2DArray
}

func GenerateGradesPDF(c echo.Context) error {
	month := utils.ConvertToTime(c.QueryParam("month"))
	startOfMonth := now.With(month).BeginningOfMonth()
	endOfMonth := now.With(month).EndOfMonth()
	classHash := c.QueryParam("selectedClass")
	class := classesDBInteractions.GetClassByHash(classHash)
	quizzes, grades := quizzesDBInteractions.GetGradesByMonthForAllUsers(classHash, startOfMonth, endOfMonth)
	organizedGrades, quizzesTotalMarks, gradesOnly := organizeGrades(quizzes, grades)
	pdfGenerator, err := gradesPDFHandler.DeliverGradesPDF(organizedGrades, quizzesTotalMarks,
		class.Organization.TeacherName, class.Name, startOfMonth, endOfMonth, quizzes, gradesOnly)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{})
	}
	c.Response().Header().Set("Content-Type", "application/pdf")
	return c.Blob(http.StatusOK, "application/pdf", pdfGenerator.Bytes())
}

func getSmallestDate(first, second time.Time) time.Time {
	if first.Before(second) {
		return first
	}
	return second
}
