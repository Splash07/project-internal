package masterdata

import (
	"github.com/labstack/echo/v4"
	apiMasterData "gitlab.ghn.vn/online/common/apicall/v5/masterdata"
	cStruct "gitlab.ghn.vn/online/common/gstruct/token"
)

type provinceHandler struct{}

// ProvinceHandler ..
var ProvinceHandler provinceHandler

// Provinces Danh sách Province ..
func (provinceHandler) Provinces(c echo.Context) (err error) {

	apiMasterData := apiMasterData.New(getRequestID(c), cfg.API)

	provinces, err := apiMasterData.Provinces()

	if err != nil {
		return
	}

	return c.JSON(success(provinces))
}

// Chi tiết Province ..
func (provinceHandler) ProvinceDetail(c echo.Context) (err error) {

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

	province, err := apiMasterData.Province(request.ProvinceID)

	if err != nil {
		return
	}

	return (c.JSON(success(province)))
}

//////////////////////////V2//////////////////////////

// GetByIDV2 ..
func (provinceHandler) GetByIDV2(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID int `json:"id" query:"id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	province, err := masterDataAPI.ProvinceDetailV2ByID(request.ProvinceID)
	if err != nil {
		return
	}
	return c.JSON(success(province))
}

// GetAllV2 ..
func (provinceHandler) GetAllV2(c echo.Context) (err error) {
	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	allProvinces, err := masterDataAPI.AllProvincesV2()
	if err != nil {
		return
	}
	return c.JSON(success(allProvinces))
}

// RemoveV2 ..
func (provinceHandler) RemoveV2(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID int `json:"id" query:"id" validate:"required"`
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

	err = masterDataAPI.RemoveProvinceV2(apiMasterData.RemoveProvinceRequest{
		ProvinceID: request.ProvinceID,
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
func (provinceHandler) AddV2(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceName   string   `json:"province_name" validate:"required"`
		NameExtension  []string `json:"name_extension"`
		ProvinceEncode string   `json:"province_encode"`
		Priority       int      `json:"priority"`
		CanUpdateCOD   *bool    `json:"can_update_cod" validate:"required"`
		RegionID       int      `json:"region_id" validate:"required"`
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

	err = masterDataAPI.AddProvinceV2(apiMasterData.AddProvinceRequest{
		ProvinceName:    request.ProvinceName,
		NameExtension:   request.NameExtension,
		ProvinceEncode:  request.ProvinceEncode,
		Priority:        request.Priority,
		CanUpdateCOD:    *request.CanUpdateCOD,
		RegionID:        request.RegionID,
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
func (provinceHandler) UpdateV2(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID            int      `json:"id" validate:"required"`
		ProvinceName          string   `json:"province_name" validate:"required"`
		Priority              int      `json:"priority"`
		Code                  string   `json:"code"`
		ProvinceEncode        string   `json:"province_encode"`
		CanUpdateCOD          *bool    `json:"can_update_cod" validate:"required"`
		RegionID              int      `json:"region_id" validate:"required"`
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

	err = masterDataAPI.UpdateProvinceV2(apiMasterData.UpdateProvinceRequest{
		ProvinceID:            request.ProvinceID,
		ProvinceName:          request.ProvinceName,
		Priority:              request.Priority,
		Code:                  request.Code,
		ProvinceEncode:        request.ProvinceEncode,
		CanUpdateCOD:          *request.CanUpdateCOD,
		RegionID:              request.RegionID,
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
func (provinceHandler) SwitchStatus(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID int `json:"id" query:"id" validate:"required"`
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

	err = masterDataAPI.SwitchStatusProvinceV2(apiMasterData.SwitchStatusProvinceRequest{
		ProvinceID:      request.ProvinceID,
		UpdatedIP:       c.RealIP(),
		UpdatedEmployee: token.UserID,
		UpdatedSource:   "internal",
	})
	if err != nil {
		return
	}

	return c.JSON(success(nil))
}
