package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/mhaqiw/product-service/domain"
	"github.com/mhaqiw/product-service/interfaces"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type ProductHandler struct {
	ProductService interfaces.ProductService
}

func NewProductHandler(e *echo.Echo, m interfaces.ProductService) {
	handler := &ProductHandler{
		ProductService: m,
	}
	e.GET("/products", handler.Get)
	e.POST("/products", handler.Post)
}

func (h *ProductHandler) Get(c echo.Context) error {
	sort := c.QueryParam("sorting")

	response, err := h.ProductService.GetAll(c.Request().Context(), sort)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) Post(c echo.Context) error {
	var request domain.ProductRequestPayload
	err := c.Bind(&request)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusUnprocessableEntity, domain.Response{Message: err.Error()})
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
	}

	resp, err := h.ProductService.AddProduct(c.Request().Context(), request)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, resp)
}
