package route

import (
	"github.com/labstack/echo/v4"
	groute "gitlab.ghn.vn/online/common/gstuff/route"
)

// Internal route
func Internal(e *echo.Echo) {
	publicAPI := groute.PublicAPIRoute(e)

	setUpI18nRoutes(publicAPI)
	setUpMasterDataRoutes(publicAPI)
}
