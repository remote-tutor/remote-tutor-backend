package pdf

import (
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

//pdf requestpdf struct
type RequestPdf struct {
	body string
}

//new request to pdf function
func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

//parsing template function
func (r *RequestPdf) ParseTemplate(templateName string, data interface{}) error {
	var templatePath string
	if os.Getenv("APP_ENV") == "development" {
		templatePath = "pdf/templates/" + templateName
	} else {
		templatePath = "/home/ubuntu/project/pdf/templates/" + templateName
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

//generate pdf function
func (r *RequestPdf) GeneratePDF() (*wkhtmltopdf.PDFGenerator, error) {
	t := time.Now().Unix()
	// write whole the body
	var filename string
	if os.Getenv("APP_ENV") == "development" {
		filename = "pdf/generated-html/" + strconv.FormatInt(t, 10) + ".html"
	} else {
		filename = "/home/ubuntu/project/pdf/generated-html/" + strconv.FormatInt(t, 10) + ".html"
	}
	err := ioutil.WriteFile(filename, []byte(r.body), 0644)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filename)
	if file != nil {
		defer file.Close()
		defer os.Remove(filename)
	}
	if err != nil {
		return nil, err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pageReader := wkhtmltopdf.NewPageReader(file)
	pageReader.FooterRight.Set("[page]/[toPage]")
	pageReader.HeaderLeft.Set(time.Now().Format("02/01/2006"))


	pdfg.AddPage(pageReader)

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg, nil
}
