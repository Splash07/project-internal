package handler

import (
	"github.com/labstack/echo/v4"
	apiInternalSSO "gitlab.ghn.vn/online/common/apicall/v5/sso"
	tokenAPI "gitlab.ghn.vn/online/common/apicall/v5/token"
	cStruct "gitlab.ghn.vn/online/common/gstruct/token"
)

// User handler
var User userHandler

func init() {
	User = userHandler{}
}

type userHandler struct {
}

// Info func
func (o userHandler) Info(c echo.Context) (err error) {
	apiSSO := apiInternalSSO.New(getRequestID(c), cfg.API)
	token := c.Get("token").(cStruct.Token)
	employeeInfo, err := apiSSO.GetStaffDetail(token.UserID)
	if err != nil {
		return
	}

	return c.JSON(success(employeeInfo))
}

func (o userHandler) InfoStaff(c echo.Context) (err error) {
	type myRequest struct {
		UserID int `json:"user_id" query:"user_id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}

	if err = c.Validate(request); err != nil {
		return
	}
	apiSSO := apiInternalSSO.New(getRequestID(c), cfg.API)
	staffInfo, err := apiSSO.GetEmployeeDetail(request.UserID)
	if err != nil {
		return
	}

	return c.JSON(success(staffInfo))
}

// Logout
func (o userHandler) Logout(c echo.Context) (err error) {
	tokenAPI := tokenAPI.New(getRequestID(c), cfg.API)

	token := c.Get("token").(cStruct.Token)

	codeMessage, err := tokenAPI.ClearByDevice(token.UserID, "employee", token.DeviceID)
	if err != nil {
		return
	}

	return c.JSON(success(codeMessage))
}
