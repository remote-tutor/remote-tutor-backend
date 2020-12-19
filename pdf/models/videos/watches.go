package videos

import watchesModel "backend/models/videos"

type WatchesPDF struct {
	Part    watchesModel.VideoPart
	Watches []watchesModel.UserWatch
}
