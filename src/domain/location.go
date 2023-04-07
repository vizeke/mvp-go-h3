package domain

import "github.com/uber/h3-go/v4"

type (
	Location struct {
		Id        int32
		Latitude  float64
		Longitude float64
	}

	HeatMap struct {
		Cell     h3.Cell
		Boundary h3.CellBoundary
		Count    int32
		Valid    bool
	}

	HeatMapCount struct {
		MaxCount int32
		HeatMap  map[int32][]*HeatMap
	}
)
