package i18n

import (
	"github.com/labstack/echo/v4"
	apicallUtilsLanguage "gitlab.ghn.vn/online/common/apicall/v5/utils/language"
	"gitlab.ghn.vn/online/common/gstuff/handler"
)

// Language handler
var Language languageHandler

func init() {
	Language = languageHandler{}
}

type languageHandler struct{}

// Get
func (languageHandler) Get(c echo.Context) (err error) {
	type myRequest struct {
		Key string `json:"key" query:"key"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	utilsAPI := apicallUtilsLanguage.New(getRequestID(c), cfg.API)
	i18nLang, err := utilsAPI.GetOneLanguage(request.Key)
	if err != nil {
		return
	}
	return c.JSON(handler.Success(i18nLang))
}

// GetAll
func (languageHandler) GetAll(c echo.Context) (err error) {

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
	utilsAPI := apicallUtilsLanguage.New(getRequestID(c), cfg.API)
	i18nLang, err := utilsAPI.GetAllLanguage(request.Skip, request.Limit)
	if err != nil {
		return
	}
	return c.JSON(handler.Success(i18nLang))
}

//Create
func (languageHandler) Create(c echo.Context) (err error) {
	type myRequest struct {
		Key      string `json:"key" query:"key" validate:"required"`
		Language string `json:"language" query:"language" validate:"required"`
	}

	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	utilsAPI := apicallUtilsLanguage.New(getRequestID(c), cfg.API)
	i18nLang, err := utilsAPI.CreateLanguage(request.Key, request.Language)
	if err != nil {
		return
	}
	return c.JSON(handler.Success(i18nLang))
}
