package codes

import (
	codesModel "backend/models/videos"
	"backend/pdf"
	codes2 "backend/pdf/models/codes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func DeliverCodesPDF(videoTitle string, codes []codesModel.Code) (*wkhtmltopdf.PDFGenerator, error) {
	//html template name
	templateName := "video-codes.html"
	requestPDF := pdf.NewRequestPdf("")
	//html template data
	codesPDF := codes2.CodesPDF{
		VideoTitle: videoTitle,
		CodesArray: codes,
	}
	codesPDF.ConstructCodes()

	if err := requestPDF.ParseTemplate(templateName, codesPDF); err != nil {
		return nil, err
	}
	return requestPDF.GeneratePDF()
}