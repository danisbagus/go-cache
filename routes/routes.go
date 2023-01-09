package routes

import (
	"github.com/danisbagus/go-cache/handler"
	"github.com/danisbagus/go-cache/repository"
	"github.com/danisbagus/go-cache/service"
	"github.com/labstack/echo/v4"
)

func ApiRoutes(route *echo.Echo) {
	ratesRepository := repository.NewRatesRepository()

	ratesService := service.NewRatesService(ratesRepository)

	rateseHandler := handler.NewRatesHandler(ratesService)

	ratesRoutes := route.Group(("/api/rates"))
	ratesRoutes.GET("", rateseHandler.GetUSDToIDRRates)
}
