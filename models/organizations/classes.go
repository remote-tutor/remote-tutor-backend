package organizations

import (
	hashUtils "backend/utils/hash"
	"gorm.io/gorm"
	"os"
)

type Class struct {
	gorm.Model
	Name             string       `json:"name"`
	Year             int          `json:"year"`
	OrganizationHash string       `json:"organizationHash" gorm:"size:25"`
	Organization     Organization `json:"organization" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OrganizationHash;references:Hash"`
	Hash             string       `json:"hash" gorm:"size:25;uniqueIndex"`
}

// this function generates the hash then update the Class created
func (class *Class) AfterCreate(tx *gorm.DB) (err error) {
	hash := hashUtils.GenerateHash([]uint{class.ID}, os.Getenv("ORGANIZATIONS_SALT"))
	tx.Model(class).UpdateColumn("hash", hash)
	return
}
