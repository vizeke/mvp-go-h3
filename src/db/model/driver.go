package model

import "database/sql"

type (
	DriverLocation struct {
		Id         int32
		Name       string
		Position   sql.NullString
		LastUpdate sql.NullString
		LatestPong sql.NullString
	}

	DriverPosition struct {
		Altitude           float64
		Provider           string
		Bearing            float64
		LocationProvider   int32
		Latitude           float64
		Accuracy           float64
		Time               uint64
		Radius             float64
		Speed              float64
		Longitude          float64
		IsFromMockProvider bool
	}
)
