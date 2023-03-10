package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mvp-beebee-h3/src/config"

	_ "github.com/lib/pq"
)

func GetOrdersLocation() []Location {
	db, err := sql.Open("postgres", config.GetDbConnection())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(
		`select o.id, l.lat, l.lng
		from orders o
		inner join locations l ON l."orderId" = o.id
		where l."sequence" = 0
		and l.deleted_at is null
		and o."statusId" = 6`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var locations []Location = make([]Location, 50)
	for rows.Next() {
		var location Location
		err := rows.Scan(&location.Id, &location.Latitude, &location.Longitude)
		if err != nil {
			fmt.Println(err)
		}
		locations = append(locations, location)
	}
	return locations
}

func GetDriversCurrentPosition() []Location {
	db, err := sql.Open("postgres", config.GetDbConnection())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT id, name, "position", "positionLastUpdate", "latestPong"
		FROM users
		where "driverStatus" = 3
		and not "positionLastUpdate" is null`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var driverLocations []DriverLocation = make([]DriverLocation, 50)
	for rows.Next() {
		var driverLocation DriverLocation
		err := rows.Scan(
			&driverLocation.Id,
			&driverLocation.Name,
			&driverLocation.Position,
			&driverLocation.LastUpdate,
			&driverLocation.LatestPong)
		if err != nil {
			fmt.Println(err)
		}
		driverLocations = append(driverLocations, driverLocation)
	}

	var locations []Location = make([]Location, 50)
	for _, v := range driverLocations {
		if v.Position.Valid {
			var p DriverPosition
			json.Unmarshal([]byte(v.Position.String), &p)
			if p.Latitude == 0 {
				continue
			}
			locations = append(locations, Location{Id: v.Id, Latitude: p.Latitude, Longitude: p.Longitude})
		}
	}

	return locations
}

type Location struct {
	Id        int32
	Latitude  float64
	Longitude float64
}

type DriverLocation struct {
	Id         int32
	Name       string
	Position   sql.NullString
	LastUpdate sql.NullString
	LatestPong sql.NullString
}

type DriverPosition struct {
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
