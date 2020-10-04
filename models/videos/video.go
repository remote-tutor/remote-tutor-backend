package videos

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	AvailableFrom time.Time `json:"availableFrom"`
	Parts []VideoPart       `json:"parts"`
}
