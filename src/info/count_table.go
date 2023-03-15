package info

import (
	"fmt"

	"github.com/rgeoghegan/tabulate"
	"github.com/uber/h3-go/v4"
)

type ResolutionCounter struct {
	Res        int
	TotalCells int
	Hexagons   int
	Pentagons  int
}

func PrintCountTable() {
	counts := make([]*ResolutionCounter, 16)

	for res := 0; res < 16; res++ {
		counts[res] = &ResolutionCounter{int(res), numCells(res), numHexagons(res), numPentagons()}
	}

	tab, _ := tabulate.Tabulate(counts, &tabulate.Layout{Format: tabulate.PipeFormat})
	fmt.Println(tab)
}

func numPentagons() int {
	return 12
}

func numHexagons(res int) int {
	// Number of *hexagons* (excluding pentagons) at a resolution
	return h3.NumCells(res) - numPentagons()
}

func numCells(res int) int {
	// Number of *hexagons* (excluding pentagons) at a resolution
	return h3.NumCells(res)
}
