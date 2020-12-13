package quizzes

import (
	quizzesModel "backend/models/quizzes"
	"backend/pdf"
	gradesPDFModel "backend/pdf/models/quizzes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"time"
)

func DeliverGradesPDF(grades []map[string]interface{}, quizzesTotalMark int, teacherName, className string,
	startDate, endDate time.Time, quizzes []quizzesModel.Quiz, gradesOnly [][]int) (*wkhtmltopdf.PDFGenerator, error) {
	//html template name
	templateName := "quizzes-grades.html"
	requestPDF := pdf.NewRequestPdf("")
	//html template data
	gradesPDF := gradesPDFModel.Grades{
		Grades:           grades,
		GradesOnly:       gradesOnly,
		TeacherName:      teacherName,
		ClassName:        className,
		StartDate:        startDate,
		EndDate:          endDate,
		Quizzes:          quizzes,
		QuizzesTotalMark: quizzesTotalMark,
	}

	if err := requestPDF.ParseTemplate(templateName, gradesPDF); err != nil {
		return nil, err
	}
	return requestPDF.GeneratePDF()
}
