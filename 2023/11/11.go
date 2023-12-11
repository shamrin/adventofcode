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
	grows, gcols := map[int]bool{}, map[int]bool{}
	for _, g := range gs {
		if g.col+1 > nc {
			nc = max(g.col+1, nc)
		}
		if g.row+1 > nr {
			nr = max(g.row+1, nr)
		}
		grows[g.row] = true
		gcols[g.col] = true
	}
	emptyrows, emptycols := []int{}, []int{}
	for r := 0; r < nr; r++ {
		if !grows[r] {
			emptyrows = append(emptyrows, r)
		}
	}
	for c := 0; c < nc; c++ {
		if !gcols[c] {
			emptycols = append(emptycols, c)
		}
	}
	slices.Sort(emptyrows)
	slices.Sort(emptycols)
	slices.Reverse(emptyrows)
	slices.Reverse(emptycols)

	for _, er := range emptyrows {
		for i := range gs {
			if gs[i].row > er {
				gs[i].row += k - 1
			}
		}
	}
	for _, ec := range emptycols {
		for i := range gs {
			if gs[i].col > ec {
				gs[i].col += k - 1
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
