package videos

import (
	"backend/diagnostics"
	partsModel "backend/models/videos"
)

func WriteVideoPartErr(err error, errorType string, part *partsModel.VideoPart) {
	filePath := "database/videos/parts.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filePath, errorType, err, part)
}

