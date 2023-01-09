package routes

import (
	"github.com/danisbagus/go-cache/handler"
	"github.com/danisbagus/go-cache/repository"
	"github.com/danisbagus/go-cache/service"
	"github.com/danisbagus/go-cache/utils/cache"
	"github.com/labstack/echo/v4"
)

func ApiRoutes(route *echo.Echo) {
	cache := cache.NewCache()
	ratesRepository := repository.NewRatesRepository()

	ratesService := service.NewRatesService(ratesRepository, cache)

	rateseHandler := handler.NewRatesHandler(ratesService)

	ratesRoutes := route.Group(("/api/rates"))
	ratesRoutes.GET("", rateseHandler.GetUSDToIDRRates)
}
