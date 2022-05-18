package content

import (
	"net/http"
	"project-art-museum/api/common"
	"project-art-museum/api/v1/creator/request"
	"project-art-museum/api/v1/creator/response"
	"project-art-museum/business"
	contentBusiness "project-art-museum/business/content"

	v10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

//Controller Get content API controller
type Controller struct {
	service   contentBusiness.Service
	validator *v10.Validate
}

//NewController Construct content API controller
func NewController(service contentBusiness.Service) *Controller {
	return &Controller{
		service,
		v10.New(),
	}
}

//GetContentByID Get content by ID echo handler
func (controller *Controller) GetContentByID(c echo.Context) error {
	ID := c.Param("id")
	content, err := controller.service.FindContentByID(ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	} else if content == nil {
		return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
	}

	response := response.NewGetContentByIDResponse(*content)
	return c.JSON(http.StatusOK, response)
}

//CreateNewContent Create new content echo handler
func (controller *Controller) CreateNewContent(c echo.Context) error {
	createContentRequest := new(request.CreateContentRequest)
	if err := c.Bind(createContentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	ID, err := controller.service.CreateContent(*createContentRequest.ToSpec())

	if err != nil {
		if err == business.ErrInvalidSpec {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewCreateNewContentResponse(ID)
	return c.JSON(http.StatusCreated, response)
}

//UpdateContent update content echo handler
func (controller *Controller) UpdateContent(c echo.Context) error {
	updateContentRequest := new(request.CreateContentRequest)

	if err := c.Bind(updateContentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	err := controller.validator.Struct(updateContentRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	err = controller.service.UpdateContent(
		c.Param("id"),
		*updateContentRequest.ToSpec(),
		updateContentRequest.Version,
		"updater")

	return c.NoContent(http.StatusNoContent)
}
