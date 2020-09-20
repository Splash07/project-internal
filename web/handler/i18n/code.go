package i18n

import (
	"net/http"

	"github.com/labstack/echo/v4"
	apicallUtilsCode "gitlab.ghn.vn/online/common/apicall/v5/utils/code"
	"gitlab.ghn.vn/online/common/gstuff/handler"
)

// Code handler
var Code codeHandler

func init() {
	Code = codeHandler{}
}

type codeHandler struct{}

func (codeHandler) Get(c echo.Context) (err error) {

	type myRequest struct {
		Code string `json:"code" query:"code" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	utilsAPI := apicallUtilsCode.New(getRequestID(c), cfg.API)
	i18nCode, err := utilsAPI.GetOneCode(request.Code)
	// i18nCode, err := code.New().GetOneByCode(request.Code)
	if err != nil {
		return
	}
	// if i18nCode.ID.IsZero() {
	// 	return echo.NewHTTPError(http.StatusNotFound, "Mã từ khóa không tồn tại")
	// }
	return c.JSON(handler.Success(i18nCode))
}

func (codeHandler) GetAll(c echo.Context) (err error) {

	type myRequest struct {
		Skip  int `json:"skip" query:"skip"`
		Limit int `json:"limit" query:"limit"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	if request.Limit == 0 {
		request.Limit = 20
	}
	// i18nCode, total, err := code.New().GetAll(request.Skip, request.Limit)
	// type i18nCodeTotal struct {
	// 	Data  []model.I18NCode `json:"data"`
	// 	Total int              `json:"total"`
	// }
	// if err != nil {
	// 	return
	// }
	// result := i18nCodeTotal{i18nCode, int(total)}

	// if err != nil {
	// 	return fmt.Errorf("Không tìm thấy message: %s", err)
	// }
	utilsAPI := apicallUtilsCode.New(getRequestID(c), cfg.API)

	result, err := utilsAPI.GetAllCode(request.Skip, request.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Mã từ khóa không tồn tại")
	}
	return c.JSON(handler.Success(result))
}

func (codeHandler) Create(c echo.Context) (err error) {
	type myRequest struct {
		Code      string `json:"code" query:"code" validate:"required"`
		Languages []struct {
			LanguageID string `json:"language_id" query:"languageId" validate:"required"`
			Text       string `json:"text" query:"text" validate:"required"`
		} `json:"languages" query:"languages" validate:"required"`
	}

	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	utilsAPI := apicallUtilsCode.New(getRequestID(c), cfg.API)
	i18nCode, err := utilsAPI.GetOneCode(request.Code)
	i18nCode.Code = request.Code
	i18nCode.Languages = []struct {
		LanguageID string `json:"language_id" bson:"languageId"`
		Text       string `json:"text" bson:"text"`
	}(request.Languages)
	i18nCode, err = utilsAPI.CreateCode(i18nCode.Code, i18nCode.Languages)
	if err != nil {
		return
	}
	return c.JSON(handler.Success(i18nCode))
}

func (codeHandler) Update(c echo.Context) (err error) {

	type myRequest struct {
		Code      string `json:"code" query:"code" validate:"required"`
		Languages []struct {
			LanguageID string `json:"language_id" query:"languageId" validate:"required"`
			Text       string `json:"text" query:"text" validate:"required"`
		} `json:"languages" query:"languages" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	utilsAPI := apicallUtilsCode.New(getRequestID(c), cfg.API)
	i18nCode, err := utilsAPI.GetOneCode(request.Code)
	i18nCode.Code = request.Code
	i18nCode.Languages = []struct {
		LanguageID string `json:"language_id" bson:"languageId"`
		Text       string `json:"text" bson:"text"`
	}(request.Languages)

	i18nCode, err = utilsAPI.UpdateCode(i18nCode.Code, i18nCode.Languages)
	if err != nil {
		return
	}
	return c.JSON(handler.Success(i18nCode))
}

// Delete
func (codeHandler) Delete(c echo.Context) (err error) {

	type myRequest struct {
		Code string `json:"code" query:"code" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	utilsAPI := apicallUtilsCode.New(getRequestID(c), cfg.API)
	i18nCode, err := utilsAPI.GetOneCode(request.Code)
	if err != nil {
		return
	}
	i18nCode.Visible = 0
	i18nCode, err = utilsAPI.DeleteCode(request.Code)
	if err != nil {
		return
	}
	return c.JSON(handler.Success(i18nCode))
}

// Search...
func (codeHandler) Search(c echo.Context) (err error) {
	type myRequest struct {
		Code  string `json:"code" query:"code"`
		Skip  int    `json:"skip" query:"skip"`
		Limit int    `json:"limit" query:"limit"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	if request.Limit == 0 {
		request.Limit = 20
	}
	utilsAPI := apicallUtilsCode.New(getRequestID(c), cfg.API)
	i18nCode, err := utilsAPI.SearchCode(request.Code, request.Skip, request.Limit)
	if err != nil {
		return
	}
	return c.JSON(handler.Success(i18nCode))
}
