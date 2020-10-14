package organizations

import (
	dbInstance "backend/database"
	"backend/database/diagnostics"
	organizationAdminsModel "backend/models/organizations"
)

func GetOrganizationAdminsByClass(class string) []organizationAdminsModel.OrganizationAdmin {
	admins := make([]organizationAdminsModel.OrganizationAdmin, 0)
	subQuery := dbInstance.GetDBConnection().Select("organization_hash").
		Where("hash = ?", class).Table("classes")
	dbInstance.GetDBConnection().Where("organization_hash IN (?)", subQuery).Preload("User").Find(&admins)
	return admins
}


func CreateOrganizationAdmin(organizationAdmin *organizationAdminsModel.OrganizationAdmin) error {
	err := dbInstance.GetDBConnection().Create(organizationAdmin).Error
	diagnostics.WriteError(err, "CreateOrganizationAdmin")
	return err
}