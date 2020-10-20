package videos

import (
	"backend/diagnostics"
	watchModel "backend/models/videos"
)

func WriteWatchErr(err error, errorType string, watch *watchModel.UserWatch) {
	filePath := "database/videos/watch.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filePath, errorType, err, watch)
}

