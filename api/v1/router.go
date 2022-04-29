package api

import (
	contentV1 "clean-hexa/api/v1/content"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	ContentV1Controller *contentV1.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	contentV1 := e.Group("/v1/content")
	contentV1.POST("", controller.ContentV1Controller.Create)
	contentV1.GET("", controller.ContentV1Controller.GetByName)
	contentV1.PATCH("", controller.ContentV1Controller.Update)
	contentV1.DELETE("", controller.ContentV1Controller.DELETE)
}
