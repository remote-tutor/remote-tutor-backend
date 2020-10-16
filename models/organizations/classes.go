package organizations

import (
	hashUtils "backend/utils/hash"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name             string       `json:"name"`
	Year             int          `json:"year"`
	OrganizationHash string       `json:"organizationHash" gorm:"size:255"`
	Organization     Organization `json:"organization" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OrganizationHash;references:Hash"`
	Hash             string       `json:"hash" gorm:"size:255;uniqueIndex"`
}

// this function generates the hash then update the Class created
func (class *Class) AfterCreate(tx *gorm.DB) (err error) {
	hash := hashUtils.GenerateHash([]uint{class.ID})
	tx.Model(class).UpdateColumn("hash", hash)
	return
}
