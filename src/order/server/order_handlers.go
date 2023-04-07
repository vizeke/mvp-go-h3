package handlers

import (
	"mvp-beebee-h3/src/db"
	"mvp-beebee-h3/src/domain"
	"mvp-beebee-h3/src/location/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	OrderHandler struct {
		LocationService service.LocationService
	}
)

func NewOrderHandler(e *echo.Echo) {
	handler := &OrderHandler{
		LocationService: service.NewLocationService(),
	}

	e.GET("/orders/:resolution", handler.OrderIndexes)
}

func (h *OrderHandler) OrderIndexes(c echo.Context) error {
	resolution, err := strconv.Atoi(c.Param("resolution"))
	if err != nil {
		return err
	}

	locations := db.GetOrdersLocation()

	maxCount, heatMap := h.LocationService.CreateHeatMap(locations, resolution)

	c.JSON(http.StatusOK, domain.HeatMapCount{MaxCount: maxCount, HeatMap: heatMap})

	return nil
}
