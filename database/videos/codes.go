package videos

import (
	dbInstance "backend/database"
	codesDiagnostics "backend/diagnostics/database/videos"
	codesModel "backend/models/videos"
)

func GenerateCodes(codes []codesModel.Code) error {
	err := dbInstance.GetDBConnection().Omit("used_by_user_id").Create(&codes).Error
	codesDiagnostics.WriteCodesErr(err, "Create", codes)
	return err
}
