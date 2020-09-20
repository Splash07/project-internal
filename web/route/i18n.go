package route

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/Splash07/project-internal/web/handler/i18n"
)

// Set up i18n routes
func setUpI18nRoutes(g *echo.Group) {
	g.Any("/i18n/code/get", i18n.Code.Get)
	g.Any("/i18n/code/get/all", i18n.Code.GetAll)
	g.Any("/i18n/code/create", i18n.Code.Create)
	g.Any("/i18n/code/update", i18n.Code.Update)
	g.Any("/i18n/code/delete", i18n.Code.Delete)
	g.Any("/i18n/code/search", i18n.Code.Search)

	g.Any("/i18n/language/get", i18n.Language.Get)
	g.Any("/i18n/language/get/all", i18n.Language.GetAll)
	g.Any("/i18n/language/create", i18n.Language.Create)
}
