package api

import (
	content "project-art-museum/api/v1/creator"

	"github.com/labstack/echo"
)

// Controller to define controller that we use
type Controller struct {
	ContentController *content.Controller
}

//RegisterPath Registera V1 API path
func RegisterPath(e *echo.Echo, ctrl Controller) {
	//content
	contentV1 := e.Group("/v1/creator")
	contentV1.GET("/v1/creator/:id", ctrl.ContentController.GetContentByID)
	contentV1.POST("/v1/creator:id", ctrl.ContentController.CreateNewContent)
	contentV1.PUT("/v1/creator/:id", ctrl.ContentController.UpdateContent)
}
