package organizations

import (
	organizationAdminsDBInteractions "backend/database/organizations"
	organizationAdminsModel "backend/models/organizations"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
)

func GetOrganizationAdminsByClass(c echo.Context) error {
	class := c.QueryParam("selectedClass")
	admins := organizationAdminsDBInteractions.GetOrganizationAdminsByClass(class)
	return c.JSON(http.StatusOK, echo.Map{
		"admins": admins,
	})
}

func AddAdminToOrganization(c echo.Context) error {
	userID := utils.ConvertToUInt(c.FormValue("userID"))
	selectedClass := c.FormValue("selectedClass")
	class := organizationAdminsDBInteractions.GetClassByHash(selectedClass)
	organizationAdmin := organizationAdminsModel.OrganizationAdmin{
		UserID: userID,
		OrganizationHash: class.OrganizationHash,
	}
	err := organizationAdminsDBInteractions.CreateOrganizationAdmin(&organizationAdmin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (admin not added), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Admin added successfully to organization - make sure to mark him as admin in this class",
	})
}