package videos

import (
	dbInstance "backend/database"
	videoParts "backend/models/videos"
)

func CreatePart(part *videoParts.VideoPart) {
	dbInstance.GetDBConnection().Create(part)
}