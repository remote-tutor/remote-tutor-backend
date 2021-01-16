package videos

import watchesModel "backend/models/videos"

type WatchesPDF struct {
	Part    watchesModel.VideoPart
	Watches []watchesModel.UserWatch
	PartNumber int
}

func (watchesPDF *WatchesPDF) AdjustPartsNumbers(parts []watchesModel.VideoPart, length int) {
	videoPartNumber := 1
	for i := 0; i < length; i++ {
		if watchesPDF.Part.ID == parts[i].ID {
			watchesPDF.PartNumber = videoPartNumber
		}
		if parts[i].IsVideo {
			videoPartNumber++
		}
	}
}