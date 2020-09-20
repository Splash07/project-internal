package masterdata

import (
	"github.com/labstack/echo/v4"
	apiMasterData "gitlab.ghn.vn/online/common/apicall/v5/masterdata"
	cStruct "gitlab.ghn.vn/online/common/gstruct/token"
)

// BankHandler ..
var BankHandler = bankHandler{}

type bankHandler struct{}

// GetOne ..
func (bankHandler) GetOne(c echo.Context) (err error) {
	type myRequest struct {
		BankID int `json:"id" query:"id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	bank, err := masterDataAPI.GetBankV2(request.BankID)
	if err != nil {
		return
	}
	return c.JSON(success(bank))
}

// GetAll ..
func (bankHandler) GetAll(c echo.Context) (err error) {
	masterDataAPI := apiMasterData.New(getRequestID(c), cfg.API)

	banks, err := masterDataAPI.GetAllBanksV2()
	if err != nil {
		return
	}
	return c.JSON(success(banks))
}

// AddNewBank ..
func (bankHandler) AddNewBank(c echo.Context) (err error) {
	type myRequest struct {
		Name      string `json:"name" validate:"required"`
		ShortName string `json:"short_name" validate:"required"`
		HasBranch bool   `json:"has_branch"`
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

	err = masterDataAPI.AddBankV2(apiMasterData.AddBankRequest{
		Name:      request.Name,
		ShortName: request.ShortName,
		HasBranch: request.HasBranch,

		//
		CreatedIP:       c.RealIP(),
		CreatedEmployee: token.UserID,
		CreatedSource:   "internal",
	})
	if err != nil {
		return
	}
	return c.JSON(success(nil))
}

// RemoveBank ..
func (bankHandler) RemoveBank(c echo.Context) (err error) {
	type myRequest struct {
		BankID int `json:"id" query:"id" validate:"required"`
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

	err = masterDataAPI.RemoveBankV2(apiMasterData.RemoveBankRequest{
		BankID:          request.BankID,
		UpdatedIP:       c.RealIP(),
		UpdatedEmployee: token.UserID,
		UpdatedSource:   "internal",
	})
	if err != nil {
		return
	}
	return c.JSON(success(nil))
}

// UpdateBank ..
func (bankHandler) UpdateBank(c echo.Context) (err error) {
	type myRequest struct {
		BankID    int    `json:"id" validate:"required"`
		Name      string `json:"name" validate:"required"`
		ShortName string `json:"short_name" validate:"required"`
		HasBranch *bool  `json:"has_branch" validate:"required"`
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

	err = masterDataAPI.UpdateBankV2(apiMasterData.UpdateBankRequest{
		BankID:    request.BankID,
		Name:      request.Name,
		ShortName: request.ShortName,
		HasBranch: *(request.HasBranch),

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

// SwitchStatus ..
func (bankHandler) SwitchStatus(c echo.Context) (err error) {
	type myRequest struct {
		BankID int `json:"id" query:"id" validate:"required"`
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

	err = masterDataAPI.SwitchBankStatusV2(apiMasterData.SwitchStatusBankRequest{
		BankID:          request.BankID,
		UpdatedIP:       c.RealIP(),
		UpdatedEmployee: token.UserID,
		UpdatedSource:   "internal",
	})
	if err != nil {
		return
	}
	return c.JSON(success(nil))
}
