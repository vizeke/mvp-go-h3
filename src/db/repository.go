package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mvp-beebee-h3/src/config"
	"mvp-beebee-h3/src/db/model"
	"mvp-beebee-h3/src/domain"

	_ "github.com/lib/pq"
)

func GetOrdersLocation() []domain.Location {
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

	var locations []domain.Location = make([]domain.Location, 50)
	for rows.Next() {
		var location domain.Location
		err := rows.Scan(&location.Id, &location.Latitude, &location.Longitude)
		if err != nil {
			fmt.Println(err)
		}
		locations = append(locations, location)
	}
	return locations
}

func GetDriversCurrentPosition() []domain.Location {
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

	var driverLocations []model.DriverLocation = make([]model.DriverLocation, 50)
	for rows.Next() {
		var driverLocation model.DriverLocation
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

	var locations []domain.Location = make([]domain.Location, 50)
	for _, v := range driverLocations {
		if v.Position.Valid {
			var p model.DriverPosition
			json.Unmarshal([]byte(v.Position.String), &p)
			if p.Latitude == 0 {
				continue
			}
			locations = append(locations, domain.Location{Id: v.Id, Latitude: p.Latitude, Longitude: p.Longitude})
		}
	}

	return locations
}
