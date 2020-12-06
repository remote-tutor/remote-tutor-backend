package videos

import (
	"backend/diagnostics"
	codesModel "backend/models/videos"
)

func WriteCodeErr(err error, errorType string, code *codesModel.Code) {
	filePath := "database/videos/codes.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filePath, errorType, err, code)
}

func WriteCodesErr(err error, errorType string, codes []codesModel.Code) {
	filepath := "database/payments/bulk-codes.log"
	for i := 0; i < len(codes); i++ {
		diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, &codes[i])
	}
}
