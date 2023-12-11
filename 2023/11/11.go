package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type G struct{ row, col int }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(gs []G, k int) int {
	gs = slices.Clone(gs)
	nr, nc := 0, 0
	for _, g := range gs {
        nc = max(g.col+1, nc)
        nr = max(g.row+1, nr)
	}
	grows, gcols := make([]bool, nr), make([]bool, nc)
	for _, g := range gs {
		grows[g.row] = true
		gcols[g.col] = true
	}
	for row := len(grows) - 1; row >= 0; row-- {
		if !grows[row] {
			for i := range gs {
				if gs[i].row > row {
					gs[i].row += k - 1
				}
			}
		}
	}
	for col := len(gcols) - 1; col >= 0; col-- {
		if !gcols[col] {
			for i := range gs {
				if gs[i].col > col {
					gs[i].col += k - 1
				}
			}
		}
	}
	s := 0
	for i, g1 := range gs {
		for _, g2 := range gs[i+1:] {
			s += abs(g1.row-g2.row) + abs(g1.col-g2.col)
		}
	}
	return s
}

func main() {
	file, err := os.ReadFile("11.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	lines := strings.Split(input, "\n")
	gs := []G{}
	for r, line := range lines {
		for c, ch := range line {
			if ch == '#' && !slices.Contains(gs, G{r, c}) {
				gs = append(gs, G{r, c})
			}
		}
	}
	fmt.Println(solve(gs, 2))
	fmt.Println(solve(gs, 1000000))
}
