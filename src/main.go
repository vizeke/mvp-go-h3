package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mvp-beebee-h3/src/db"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/cors"
	"github.com/uber/h3-go/v4"
	"hawx.me/code/route"
)

func main() {
	// Tests
	latLng := h3.NewLatLng(37.775938728915946, -122.41795063018799)
	resolution := 6 // between 0 (biggest cell) and 15 (smallest cell)

	cell := h3.LatLngToCell(latLng, resolution)
	center := h3.CellToLatLng(cell)
	bound := h3.CellToBoundary(cell)
	res := cell.Resolution()
	cellNumber := cell.BaseCellNumber()
	strCell := cell.String()
	uIntCell := h3.IndexFromString(strCell)

	icosahedrons := cell.IcosahedronFaces()

	neighbors2 := cell.GridDisk(2)
	distances2 := cell.GridDiskDistances(2)

	fmt.Println(cell)
	fmt.Println(center)
	fmt.Println(bound)
	fmt.Println(res)
	fmt.Println(cellNumber)
	fmt.Println(strCell)
	fmt.Println(uIntCell)
	fmt.Println(cell.IsValid())
	fmt.Println(icosahedrons)
	fmt.Println(neighbors2)
	fmt.Println(len(neighbors2))
	fmt.Println(distances2)
	// Output:
	// 8928308280fffff

	route.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	route.HandleFunc("/orders/:resolution", func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)

		// string to int
		resolution, err := strconv.Atoi(vars["resolution"])
		if err != nil {
			// ... handle error
			panic(err)
		}

		locations := db.GetOrdersLocation()

		maxCount, heatMap := createHeatMap(locations, resolution)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(HeatMapResponse{MaxCount: maxCount, HeatMap: heatMap})
	})

	route.HandleFunc("/drivers/:resolution", func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)

		// string to int
		resolution, err := strconv.Atoi(vars["resolution"])
		if err != nil {
			// ... handle error
			panic(err)
		}

		locations := db.GetDriversCurrentPosition()

		maxCount, heatMap := createHeatMap(locations, resolution)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(HeatMapResponse{MaxCount: maxCount, HeatMap: heatMap})
	})

	fs := http.FileServer(http.Dir("static"))
	route.Handle("/", fs)

	handler := cors.Default().Handler(route.Default)

	s := &http.Server{
		Addr:           ":8000",
		Handler:        handler,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.

	fmt.Println("Listening on 8000")
	log.Fatal(s.ListenAndServe())
}

func createHeatMap(locations []db.Location, resolution int) (int32, map[int32][]*HeatMap) {
	cells := map[string]*HeatMap{}
	for _, l := range locations {
		latLng := h3.NewLatLng(l.Latitude, l.Longitude)
		cell := h3.LatLngToCell(latLng, resolution)
		cString := cell.String()

		a := cells[cString]
		if a == nil || !a.Valid {
			heatMap := &HeatMap{
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

	heatMap := map[int32][]*HeatMap{}
	for _, v := range cells {
		if heatMap[v.Count] == nil {
			heatMap[v.Count] = make([]*HeatMap, countHeatMap[v.Count])
		}
		heatMap[v.Count][countHeatMap[v.Count]-1] = v
		countHeatMap[v.Count]--
	}

	return maxCount, heatMap
}

type HeatMapResponse struct {
	MaxCount int32
	HeatMap  map[int32][]*HeatMap
}

type HeatMap struct {
	Cell     h3.Cell
	Boundary h3.CellBoundary
	Count    int32
	Valid    bool
}
