package videos

import (
	watchesModel "backend/models/videos"
	"backend/pdf"
	watchesPDFModel "backend/pdf/models/videos"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func DeliverWatchesPDF(part *watchesModel.VideoPart, watches []watchesModel.UserWatch) (*wkhtmltopdf.PDFGenerator, error) {
	//html template name
	templateName := "video-part-watches.html"
	requestPDF := pdf.NewRequestPdf("")
	//html template data
	watchesPDF := watchesPDFModel.WatchesPDF{
		Part: *part,
		Watches: watches,
	}
	if err := requestPDF.ParseTemplate(templateName, watchesPDF); err != nil {
		return nil, err
	}
	return requestPDF.GeneratePDF()
}
