package i18n

import (
	"github.com/labstack/echo/v4"
	"gitlab.ghn.vn/online/common/config"
	ghandler "gitlab.ghn.vn/online/common/gstuff/handler"
)

var cfg = config.GetConfig()

func getRequestID(c echo.Context) (requestID string) {
	return ghandler.GetRequestID(c)
}

// success func
func success(data interface{}) (int, interface{}) {
	return ghandler.Success(data)
}

func Error(code int, message string) (int, interface{}) {
	return code, ghandler.ResponseContent{
		Code:    code,
		Message: message,
	}
}

func ErrorPassword(status, code int, message string) (int, interface{}) {
	return status, ghandler.ResponseContent{
		Code:    code,
		Message: message,
	}
}
