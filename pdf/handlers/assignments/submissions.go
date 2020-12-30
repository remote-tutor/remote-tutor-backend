package assignments

import (
	assignmentsModel "backend/models/assignments"
	classesModel "backend/models/organizations"
	"backend/pdf"
	assignmentsPDFModel "backend/pdf/models/assignments"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func DeliverAssignmentsPDF(assignment *assignmentsModel.Assignment,
	submissions []assignmentsModel.AssignmentSubmission, class *classesModel.Class) (*wkhtmltopdf.PDFGenerator, error) {
	//html template name
	templateName := "assignment-submissions.html"
	requestPDF := pdf.NewRequestPdf("")
	//html template data
	submissionsPDF := assignmentsPDFModel.SubmissionsPDF{
		Assignment:  *assignment,
		Submissions: submissions,
		TeacherName: class.Organization.TeacherName,
		ClassName:   class.Name,
	}
	if err := requestPDF.ParseTemplate(templateName, submissionsPDF); err != nil {
		return nil, err
	}
	return requestPDF.GeneratePDF()
}
