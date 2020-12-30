package assignments

import (
	assignmentsModel "backend/models/assignments"
	"time"
)

type SubmissionsPDF struct {
	Assignment  assignmentsModel.Assignment
	Submissions []assignmentsModel.AssignmentSubmission
	TeacherName string
	ClassName   string
}

func (submissionsPDF SubmissionsPDF) IsSubmissionInTime(deadline, submittedAt time.Time) bool {
	return submittedAt.Before(deadline)
}
