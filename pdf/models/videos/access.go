package videos

import videosModel "backend/models/videos"

type VideoAccess struct {
	Codes []videosModel.Code
	VideoTitle string
}
