package service

import (
	"drivers-location-h3/src/domain"

	"github.com/uber/h3-go/v4"
)

type (
	LocationService interface {
		CreateHeatMap(locations []domain.Location, resolution int) (int32, map[int32][]*domain.HeatMap)
	}

	locatoinService struct{}
)

// NewDeckService will create new an deckService object representation of domain.DeckService interface.
func NewLocationService() LocationService {
	return &locatoinService{}
}

func (l *locatoinService) CreateHeatMap(locations []domain.Location, resolution int) (int32, map[int32][]*domain.HeatMap) {
	cells := map[string]*domain.HeatMap{}
	for _, l := range locations {
		latLng := h3.NewLatLng(l.Latitude, l.Longitude)
		cell := h3.LatLngToCell(latLng, resolution)
		cString := cell.String()

		a := cells[cString]
		if a == nil || !a.Valid {
			heatMap := &domain.HeatMap{
				Valid:    true,
				Cell:     cell,
				Boundary: cell.Boundary(),
				Count:    1}
			cells[cString] = heatMap
		} else {
			a.Count += 1
		}
	}

	countHeatMap := map[int32]int32{}
	maxCount := int32(0)
	for _, v := range cells {
		countHeatMap[v.Count]++
		if v.Count > maxCount {
			maxCount = v.Count
		}
	}

	heatMap := map[int32][]*domain.HeatMap{}
	for _, v := range cells {
		if heatMap[v.Count] == nil {
			heatMap[v.Count] = make([]*domain.HeatMap, countHeatMap[v.Count])
		}
		heatMap[v.Count][countHeatMap[v.Count]-1] = v
		countHeatMap[v.Count]--
	}

	return maxCount, heatMap
}
