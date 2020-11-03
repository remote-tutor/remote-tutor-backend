package videos

import (
	"backend/diagnostics"
	videosModel "backend/models/videos"
)

func WriteVideoErr(err error, errorType string, video *videosModel.Video) {
	filePath := "database/videos/videos.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filePath, errorType, err, video)
}


