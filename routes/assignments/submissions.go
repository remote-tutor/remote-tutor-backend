package assignments

import (
	submissionsController "backend/controllers/assignments"
	"github.com/labstack/echo"
)

func InitializeSubmissionsRoutes(assignments *echo.Group, adminAssignments *echo.Group) {
	submissions := assignments.Group("/submissions")
	submissions.GET("/submission", submissionsController.GetSubmissionByUserAndAssignment)
	submissions.POST("", submissionsController.CreateOrUpdateSubmission)
	submissions.PUT("", submissionsController.CreateOrUpdateSubmission)

	adminSubmissions := adminAssignments.Group("/submissions")
	adminSubmissions.GET("", submissionsController.GetSubmissionsByAssignmentForAllUsers)
	adminSubmissions.PUT("", submissionsController.UpdateSubmissionByAdmin)
	adminSubmissions.GET("/pdf", submissionsController.GetSubmissionsPDFByAssignment)
}
