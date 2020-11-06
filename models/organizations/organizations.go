package organizations

import (
	hashUtils "backend/utils/hash"
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	TeacherName      string `json:"teacherName"`
	Subject          string `json:"subject"`
	S3BucketName     string `json:"s3BucketName"`
	CloudfrontDomain string `json:"cloudfrontDomain"`
	Hash             string `json:"hash" gorm:"size:25;uniqueIndex"`
}

// this function generates the hash then update the Organization created
func (organization *Organization) AfterCreate(tx *gorm.DB) (err error) {
	hash := hashUtils.GenerateHash([]uint{organization.ID})
	tx.Model(organization).UpdateColumn("hash", hash)
	return
}
