package masterdata

import (
	"github.com/labstack/echo/v4"
	apiMasterData "gitlab.ghn.vn/online/common/apicall/v5/masterdata"
	cStruct "gitlab.ghn.vn/online/common/gstruct/token"
)

type districtHandler struct{}

// DistrictHandler ..
var DistrictHandler districtHandler

// Districts Danh sách Province ..
func (districtHandler) Districts(c echo.Context) (err error) {

	apiMasterData := apiMasterData.New(getRequestID(c), cfg.API)

	type myRequest struct {
		ProvinceID int `json:"province_id" query:"province_id" validate:"required"`
	}

	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	districts, err := apiMasterData.Districts(request.ProvinceID)

	if err != nil {
		return
	}

	return c.JSON(success(districts))
}

// Chi tiết District ..
func (districtHandler) DistrictDetail(c echo.Context) (err error) {

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

	province, err := apiMasterData.District(request.DistrictID)

	if err != nil {
		return
	}

	return (c.JSON(success(province)))
}

//////////////////////////V2//////////////////////////

// GetByIDV2 ..
func (districtHandler) GetByIDV2(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID int `json:"id" query:"id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	district, err := masterDataAPI.DistrictDetailV2ByID(request.DistrictID)
	if err != nil {
		return
	}
	return c.JSON(success(district))
}

// GetAllV2 ..
func (districtHandler) DistrictsByProvinceIDV2(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID int `json:"province_id" query:"province_id"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}

	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	districts, err := masterDataAPI.DistrictsV2ByProvinceID(request.ProvinceID)
	if err != nil {
		return
	}
	return c.JSON(success(districts))
}

// RemoveV2 ..
func (districtHandler) RemoveV2(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID int `json:"id" query:"id" validate:"required"`
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

	err = masterDataAPI.RemoveDistrictV2(apiMasterData.RemoveDistrictRequest{
		DistrictID: request.DistrictID,
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
func (districtHandler) AddV2(c echo.Context) (err error) {
	type myRequest struct {
		DistrictName   string   `json:"district_name" validate:"required"`
		ProvinceID     int      `json:"province_id"`
		NameExtension  []string `json:"name_extension"`
		DistrictEncode string   `json:"district_encode"`
		Priority       int      `json:"priority"`
		Type           int      `json:"type"`
		SupportType    int      `json:"support_type" validate:"required"`
		CanUpdateCOD   *bool    `json:"can_update_cod" validate:"required"`
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

	err = masterDataAPI.AddDistrictV2(apiMasterData.AddDistrictRequest{
		DistrictName:    request.DistrictName,
		ProvinceID:      request.ProvinceID,
		NameExtension:   request.NameExtension,
		DistrictEncode:  request.DistrictEncode,
		Priority:        request.Priority,
		Type:            request.Type,
		SupportType:     request.SupportType,
		CanUpdateCOD:    *request.CanUpdateCOD,
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
func (districtHandler) UpdateV2(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID            int      `json:"id" validate:"required"`
		DistrictName          string   `json:"district_name" validate:"required"`
		Priority              int      `json:"priority"`
		DistrictEncode        string   `json:"district_encode"`
		Type                  int      `json:"type"`
		SupportType           int      `json:"support_type"`
		CanUpdateCOD          *bool    `json:"can_update_cod" validate:"required"`
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

	err = masterDataAPI.UpdateDistrictV2(apiMasterData.UpdateDistrictRequest{
		DistrictID:            request.DistrictID,
		DistrictName:          request.DistrictName,
		Priority:              request.Priority,
		DistrictEncode:        request.DistrictEncode,
		Type:                  request.Type,
		SupportType:           request.SupportType,
		CanUpdateCOD:          *request.CanUpdateCOD,
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
func (districtHandler) SwitchStatus(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID int `json:"id" query:"id" validate:"required"`
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

	err = masterDataAPI.SwitchStatusDistrictV2(apiMasterData.SwitchStatusDistrictRequest{
		DistrictID:      request.DistrictID,
		UpdatedIP:       c.RealIP(),
		UpdatedEmployee: token.UserID,
		UpdatedSource:   "internal",
	})
	if err != nil {
		return
	}
	return c.JSON(success(nil))
}
