package videos

import "gorm.io/gorm"

type VideoPart struct {
	gorm.Model
	Video   Video  `json:"video" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VideoID uint   `json:"videoID"`
	Link    string `json:"link"`
	Number  int    `json:"number"`
	Name    string `json:"name"`
	IsVideo bool   `json:"isVideo"`
}
