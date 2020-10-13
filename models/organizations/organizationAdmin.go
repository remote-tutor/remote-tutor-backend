package organizations

import (
	usersModel "backend/models/users"
	"gorm.io/gorm"
)

// OrganizationAdmin holds the admins users for each organization
type OrganizationAdmin struct {
	gorm.Model
	UserID           uint            `json:"userID" gorm:"uniqueIndex:idx_admin_organization,sort:asc"`
	User             usersModel.User `json:"user"`
	OrganizationHash string          `json:"organizationHash" gorm:"size:255;uniqueIndex:idx_admin_organization,sort:asc"`
	Organization     Organization    `json:"organization" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OrganizationHash;references:Hash"`
}
