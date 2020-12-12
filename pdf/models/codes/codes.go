package codes

import (
	codesModel "backend/models/videos"
	"math"
)

type CodesPDF struct {
	VideoTitle string
	CodesArray []codesModel.Code
	Codes      [][]codesModel.Code
	CodesMap   map[string][][]codesModel.Code
}

func (codesPDF *CodesPDF) ConstructCodes() {
	numberOfCodes := len(codesPDF.CodesArray)
	chunk := 4
	numberOfRows := int(math.Ceil(float64(numberOfCodes) / float64(chunk)))
	codesPDF.Codes = make([][]codesModel.Code, numberOfRows)
	for i := 0; i < numberOfRows; i++ {
		codesPDF.Codes[i] = make([]codesModel.Code, chunk)
		for j := 0; j < chunk && i * chunk + j < numberOfCodes; j++ {
			codesPDF.Codes[i][j] = codesPDF.CodesArray[i * chunk + j]
		}
	}
}