package masterdata

import (
	"github.com/labstack/echo/v4"
	apiMasterData "gitlab.ghn.vn/online/common/apicall/v5/masterdata"
	cStruct "gitlab.ghn.vn/online/common/gstruct/token"
)

type wardHandler struct{}

// WardHandler ..
var WardHandler wardHandler

// Wards Danh sách Wards ..
func (wardHandler) Wards(c echo.Context) (err error) {

	apiMasterData := apiMasterData.New(getRequestID(c), cfg.API)

	type myRequest struct {
		DistrictID int `json:"district_id" query:"district_id" validate:"required"`
	}

	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	wards, err := apiMasterData.Wards(request.DistrictID)

	if err != nil {
		return
	}

	return c.JSON(success(wards))
}

// Chi tiết Wards ..
func (wardHandler) WardDetail(c echo.Context) (err error) {

	apiMasterData := apiMasterData.New(getRequestID(c), cfg.API)

	type myRequest struct {
		WardCode string `json:"ward_code" query:"ward_code" validate:"required"`
	}

	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	ward, err := apiMasterData.Ward(request.WardCode)

	if err != nil {
		return
	}

	return (c.JSON(success(ward)))
}

//==========================V2 ==========================//

// GetByIDV2 ..
func (wardHandler) GetByCodeV2(c echo.Context) (err error) {
	type myRequest struct {
		WardCode string `json:"code" query:"code" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	ward, err := masterDataAPI.WardDetailV2ByCode(request.WardCode)
	if err != nil {
		return
	}
	return c.JSON(success(ward))
}

// GetAllV2 ..
func (wardHandler) WardsByDistrictIDV2(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID int `json:"district_id" query:"district_id"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}

	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	wards, err := masterDataAPI.WardsV2ByDistrictID(request.DistrictID)
	if err != nil {
		return
	}
	return c.JSON(success(wards))
}

// RemoveV2 ..
func (wardHandler) RemoveV2(c echo.Context) (err error) {
	type myRequest struct {
		WardCode string `json:"code" query:"code" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	token := c.Get("token").(cStruct.Token)
	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	err = masterDataAPI.RemoveWardV2(apiMasterData.RemoveWardRequest{
		WardCode: request.WardCode,
		//
		UpdatedIP:       c.RealIP(),
		UpdatedEmployee: token.UserID,
		UpdatedSource:   "internal",
	})
	if err != nil {
		return
	}
	return c.JSON(success(nil))
}

// AddV2 ..
func (wardHandler) AddV2(c echo.Context) (err error) {
	type myRequest struct {
		WardName      string   `json:"ward_name" validate:"required"`
		DistrictID    int      `json:"district_id"`
		NameExtension []string `json:"name_extension"`
		WardEncode    string   `json:"ward_encode"`
		Priority      int      `json:"priority"`
		CanUpdateCOD  *bool    `json:"can_update_cod" validate:"required"`
		SupportType   int      `json:"support_type" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	token := c.Get("token").(cStruct.Token)
	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	err = masterDataAPI.AddWardV2(apiMasterData.AddWardRequest{
		WardName:        request.WardName,
		DistrictID:      request.DistrictID,
		NameExtension:   request.NameExtension,
		WardEncode:      request.WardEncode,
		Priority:        request.Priority,
		CanUpdateCOD:    *request.CanUpdateCOD,
		SupportType:     request.SupportType,
		CreatedIP:       c.RealIP(),
		CreatedEmployee: token.UserID,
		CreatedSource:   "internal",
	})
	if err != nil {
		return
	}
	return c.JSON(success(nil))
}

// UpdateV2 ..
func (wardHandler) UpdateV2(c echo.Context) (err error) {
	type myRequest struct {
		WardCode              string   `json:"code" validate:"required"`
		WardName              string   `json:"ward_name" validate:"required"`
		Priority              int      `json:"priority"`
		WardEncode            string   `json:"ward_encode"`
		CanUpdateCOD          *bool    `json:"can_update_cod" validate:"required"`
		SupportType           int      `json:"support_type" validate:"required"`
		NameExtensionToAdd    []string `json:"name_extension_add"`
		NameExtensionToRemove []string `json:"name_extension_remove"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	token := c.Get("token").(cStruct.Token)
	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	err = masterDataAPI.UpdateWardV2(apiMasterData.UpdateWardRequest{
		WardCode:              request.WardCode,
		WardName:              request.WardName,
		Priority:              request.Priority,
		WardEncode:            request.WardEncode,
		CanUpdateCOD:          *request.CanUpdateCOD,
		SupportType:           request.SupportType,
		NameExtensionToAdd:    request.NameExtensionToAdd,
		NameExtensionToRemove: request.NameExtensionToRemove,

		UpdatedIP:       c.RealIP(),
		UpdatedEmployee: token.UserID,
		UpdatedSource:   "internal",
	})
	if err != nil {
		return
	}

	return c.JSON(success(nil))
}

// SwitchStatus ..
func (wardHandler) SwitchStatus(c echo.Context) (err error) {
	type myRequest struct {
		WardCode string `json:"code" query:"code" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	token := c.Get("token").(cStruct.Token)
	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	err = masterDataAPI.SwitchStatusWardV2(apiMasterData.SwitchStatusWardRequest{
		WardCode:        request.WardCode,
		UpdatedIP:       c.RealIP(),
		UpdatedEmployee: token.UserID,
		UpdatedSource:   "internal",
	})
	if err != nil {
		return
	}
	return c.JSON(success(nil))
}
