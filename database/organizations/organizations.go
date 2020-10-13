package organizations

import (
	dbInstance "backend/database"
	organizationsModel "backend/models/organizations"
)

func GetOrganizationAdminsByClass(class string) []organizationsModel.OrganizationAdmin {
	admins := make([]organizationsModel.OrganizationAdmin, 0)
	subQuery := dbInstance.GetDBConnection().Select("organization_hash").
		Where("hash = ?", class).Table("classes")
	dbInstance.GetDBConnection().Where("organization_hash IN (?)", subQuery).Preload("User").Find(&admins)
	return admins
}

