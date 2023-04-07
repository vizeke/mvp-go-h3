package main

import (
	"fmt"
	"testing"

	"github.com/uber/h3-go/v4"
)

func TestH3ApI(t *testing.T) {
	// Tests
	latLng := h3.NewLatLng(37.775938728915946, -122.41795063018799)
	resolution := 6 // between 0 (biggest cell) and 15 (smallest cell)

	cell := h3.LatLngToCell(latLng, resolution)
	if cell.String() != "86283082fffffff" {
		t.Error("Wrong cell for location")
	}

	center := h3.CellToLatLng(cell)
	if center.Lat != 37.77351509723 && center.Lng != -122.41827103692466 {
		fmt.Println(center.Lng)
		t.Error("Wrong geolocation cell center")
	}

	bound := h3.CellToBoundary(cell)
	if len(bound) != 6 {
		t.Error("Wrong cell boundaries length")
	}

	res := cell.Resolution()
	if res != resolution {
		t.Error("Wrong cell resolution")
	}

	cellNumber := cell.BaseCellNumber()
	if cellNumber != 20 {
		t.Error("Wrong cell Number")
	}

	uIntCell := h3.IndexFromString(cell.String())
	if uIntCell != 604189371209351167 {
		t.Error("Wrong cell uInt")
	}

	if !cell.IsValid() {
		t.Error("Invalid Cell")
	}

	icosahedrons := cell.IcosahedronFaces()
	if len(icosahedrons) != 1 && icosahedrons[0] != 7 {
		t.Error("Invalid icosahedrons faces")
	}

	neighbors2 := cell.GridDisk(2)
	if len(neighbors2) != 19 {
		t.Error("Invalid level 2 neightbors count")
	}

	distances2 := cell.GridDiskDistances(2)
	if len(distances2) != 3 {
		t.Error("Invalid level 2 grid disk distances count")
	}
	if len(distances2[0]) != 1 || len(distances2[1]) != 6 || len(distances2[2]) != 12 {
		t.Error("Invalid level 2 grid disk distances count")
	}
}
