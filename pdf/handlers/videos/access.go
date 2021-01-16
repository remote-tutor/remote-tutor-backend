package videos

import (
	videosModel "backend/models/videos"
	"backend/pdf"
	videosPDF "backend/pdf/models/videos"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func DeliverVideoAccessPDF(videoTitle string, codes []videosModel.Code) (*wkhtmltopdf.PDFGenerator, error) {
	//html template name
	templateName := "video-access.html"
	requestPDF := pdf.NewRequestPdf("")
	//html template data
	accessPDF := videosPDF.VideoAccess{
		Codes:      codes,
		VideoTitle: videoTitle,
	}

	if err := requestPDF.ParseTemplate(templateName, accessPDF); err != nil {
		return nil, err
	}
	return requestPDF.GeneratePDF()
}
