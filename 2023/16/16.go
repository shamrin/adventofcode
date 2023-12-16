package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
)

type Tile struct{ x, y int }
type Vec struct{ x, y int }
type Beam struct {
	tile Tile
	dir  Vec
}

var right = Vec{+1, 0}
var left = Vec{-1, 0}
var down = Vec{0, +1}
var up = Vec{0, -1}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func beam(pass map[Beam]bool, grid map[Tile]rune, tile Tile, dir Vec) {
	tile = Tile{tile.x + dir.x, tile.y + dir.y}
	if _, exists := grid[tile]; !exists || pass[Beam{tile, dir}] {
		return
	}
	pass[Beam{tile, dir}] = true
	switch grid[tile] {
	case '.':
		beam(pass, grid, tile, dir)
	case '/':
		switch dir {
		case right:
			dir = up
		case left:
			dir = down
		case down:
			dir = left
		case up:
			dir = right
		}
		beam(pass, grid, tile, dir)
	case '\\':
		switch dir {
		case right:
			dir = down
		case left:
			dir = up
		case down:
			dir = right
		case up:
			dir = left
		}
		beam(pass, grid, tile, dir)
	case '-':
		switch dir {
		case right, left:
			beam(pass, grid, tile, dir)
		case down, up:
			beam(pass, grid, tile, left)
			beam(pass, grid, tile, right)
		}
	case '|':
		switch dir {
		case up, down:
			beam(pass, grid, tile, dir)
		case left, right:
			beam(pass, grid, tile, up)
			beam(pass, grid, tile, down)
		}
	}
}

func energy(grid map[Tile]rune, tile Tile, dir Vec) int {
	pass := map[Beam]bool{}
	beam(pass, grid, tile, dir)
	energized := map[Tile]bool{}
	for e := range pass {
		energized[e.tile] = true
	}
	return len(energized)
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	file, err := os.ReadFile("16.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	grid := map[Tile]rune{}
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, c := range line {
			grid[Tile{x, y}] = c
		}
	}
	fmt.Println(energy(grid, Tile{-1, 0}, right))

	ny, nx := len(lines), len(lines[0])
	maxEnergy := 0
	for y := 0; y < ny; y++ {
		maxEnergy = max(maxEnergy, energy(grid, Tile{-1, y}, right))
		maxEnergy = max(maxEnergy, energy(grid, Tile{nx, y}, left))
	}
	for x := 0; x < nx; x++ {
		maxEnergy = max(maxEnergy, energy(grid, Tile{x, -1}, down))
		maxEnergy = max(maxEnergy, energy(grid, Tile{x, ny}, up))
	}
	fmt.Println(maxEnergy)
}
