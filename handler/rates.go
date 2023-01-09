package handler

import (
	"net/http"

	"github.com/danisbagus/go-cache/service"
	"github.com/labstack/echo/v4"
)

type ResourceHandler struct {
	service service.IRatesService
}

func NewRatesHandler(service service.IRatesService) *ResourceHandler {
	return &ResourceHandler{
		service: service,
	}
}

func (h *ResourceHandler) GetUSDToIDRRates(c echo.Context) error {
	rates, err := h.service.GetUSDToIDRRates()
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully get rates",
		"data":    rates,
	})
}
