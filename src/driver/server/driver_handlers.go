package handlers

import (
	"drivers-location-h3/src/db"
	"drivers-location-h3/src/domain"
	"drivers-location-h3/src/location/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Represents the httphandler for deck.
type DriverHandler struct {
	LocationService service.LocationService
}

// NewDeckHandler will initialize the decks/ resource endpoints.
func NewDriverHandler(e *echo.Echo) {
	handler := &DriverHandler{
		LocationService: service.NewLocationService(),
	}

	e.GET("/drivers/:resolution", handler.DriverIndexes)
}

func (h *DriverHandler) DriverIndexes(c echo.Context) error {
	// string to int
	resolution, err := strconv.Atoi(c.Param("resolution"))
	if err != nil {
		return err
	}

	locations := db.GetDriversCurrentPosition()

	maxCount, heatMap := h.LocationService.CreateHeatMap(locations, resolution)

	c.JSON(http.StatusOK, domain.HeatMapCount{MaxCount: maxCount, HeatMap: heatMap})

	return nil
}
