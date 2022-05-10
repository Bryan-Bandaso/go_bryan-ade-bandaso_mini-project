package content

import (
	"fmt"
	"net/http"
	"project-art-museum/api/v1/creator/request"
	contentBusiness "project-art-museum/business/creator/content"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service contentBusiness.Service
}

func NewController(service contentBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// func (controller *Controller) GetAll(c echo.Context) error {
// 	contents, err := controller.service.GetContents()

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	response := response.NewGetContents(contents)

// 	return c.JSON(http.StatusOK, response)
// }

func (controller *Controller) Create(c echo.Context) error {
	createContentRequest := new(request.CreateContentRequest)
	if err := c.Bind(createContentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println("createContentRequest v2: ", createContentRequest)

	req := *createContentRequest.ToSpec()

	err := controller.service.CreateContent(req)
	if err != nil {
		fmt.Println("masuk err sini")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
