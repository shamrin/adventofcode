package main

import (
	"fmt"
	"os"
	"strings"
)

var boxchars = map[rune]rune{
	'|': '│',
	'-': '─',
	'L': '└',
	'J': '┘',
	'7': '┐',
	'F': '┌',
	'.': '.',
	'S': 'S',
}

// var boldchars = map[rune]rune{
// 	'│': '┃',
// 	'─': '━',
// 	'└': '┗',
// 	'┘': '┛',
// 	'┐': '┓',
// 	'┌': '┏',
// 	'S': '┏', // hardcoded
// }

type xy struct {
	x int
	y int
}

func size(field map[xy]rune) (nrows, ncols int) {
	ncols, nrows = 0, 0
	for xy := range field {
		if xy.x+1 > ncols {
			ncols = xy.x + 1
		}
		if xy.y+1 > nrows {
			nrows = xy.y + 1
		}
	}
	return
}

func print(field map[xy]rune) {
	nrows, ncols := size(field)
	fmt.Println("    01234567890")
	for y := 0; y < nrows; y++ {
		fmt.Printf("%3d ", y)
		for x := 0; x < ncols; x++ {
			fmt.Print(string(field[xy{x, y}]))
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.ReadFile("10.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	// 	input = `..F7.
	// .FJ|.
	// SJ.L7
	// |F--J
	// LJ...`
	// 	input = `...........
	// .S-------7.
	// .|F-----7|.
	// .||.....||.
	// .||.....||.
	// .|L-7.F-J|.
	// .|..|.|..|.
	// .L--J.L--J.
	// ...........`

	pairs := []string{}
	for from, to := range boxchars {
		pairs = append(pairs, string(from), string(to))
	}
	r := strings.NewReplacer(pairs...)
	input = r.Replace(input)

	lines := strings.Split(input, "\n")
	field := map[xy]rune{}
	start := xy{}
	for y, line := range lines {
		x := 0
		for _, c := range line {
			field[xy{x, y}] = c
			if c == 'S' {
				start = xy{x, y}
			}
			x++
		}
	}
	print(field)

	steps := 0
	loop := map[xy]bool{}
	prev := start
	loop[start] = true
	pos := xy{start.x + 1, start.y} // hardcoded!
	for pos != start {
		nextPrev := pos
		loop[pos] = true
		switch field[pos] {
		case 'S':
			panic("oops")
		case '│':
			if (xy{pos.x, pos.y - 1}) == prev {
				pos.y++
			} else {
				pos.y--
			}
		case '─':
			if (xy{pos.x - 1, pos.y}) == prev {
				pos.x++
			} else {
				pos.x--
			}
		case '└':
			if (xy{pos.x, pos.y - 1}) == prev {
				pos.x++
			} else {
				pos.y--
			}
		case '┘':
			if (xy{pos.x - 1, pos.y}) == prev {
				pos.y--
			} else {
				pos.x--
			}
		case '┐':
			if (xy{pos.x - 1, pos.y}) == prev {
				pos.y++
			} else {
				pos.x--
			}
		case '┌':
			if (xy{pos.x, pos.y + 1}) == prev {
				pos.x++
			} else {
				pos.y++
			}
		default:
			panic("oops")
		}
		steps++
		prev = nextPrev
	}
	fmt.Println(steps/2 + 1)

	nrows, ncols := size(field)
	ins := 0
	for y := 0; y < nrows; y++ {
		up, do := false, false
		for x := 0; x < ncols; x++ {
			c := field[xy{x, y}]
			if !loop[xy{x, y}] {
				if up != do {
					panic("oops")
				}
				if up {
					field[xy{x, y}] = 'i'
					ins++
				} else {
					field[xy{x, y}] = 'o'
				}
				continue
			}
			switch c {
			case '│':
				if up != do {
					panic("oops")
				}
				up = !up
				do = !do
			case '─':
				if up == do {
					panic("oops")
				}
			case '└':
				if up != do {
					panic("oops")
				}
				up = !up
			case '┘':
				if up == do {
					panic("oops")
				}
				up = !up
			case '┐':
				if up == do {
					panic("oops")
				}
				do = !do
			case '┌', 'S': // hardcoded
				if up != do {
					panic("oops")
				}
				do = !do
			default:
				panic("oops")
			}
		}
	}
	fmt.Println(ins)
}
