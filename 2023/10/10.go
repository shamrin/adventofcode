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

var boldchars = map[rune]rune{
	'│': '┃',
	'─': '━',
	'└': '┗',
	'┘': '┛',
	'┐': '┓',
	'┌': '┏',
	'S': '┏', // hardcoded
}

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

	steps := 0
	loop := map[xy]bool{}
	prev := start
	loop[start] = true
	x, y := start.x+1, start.y // hardcoded
	for (xy{x, y} != start) {
		pos := xy{x, y}
		loop[pos] = true
		switch field[pos] {
		case 'S':
			panic("oops")
		case '│':
			if y-1 == prev.y {
				y++
			} else {
				y--
			}
		case '─':
			if x-1 == prev.x {
				x++
			} else {
				x--
			}
		case '└':
			if y-1 == prev.y {
				x++
			} else {
				y--
			}
		case '┘':
			if x-1 == prev.x {
				y--
			} else {
				x--
			}
		case '┐':
			if x-1 == prev.x {
				y++
			} else {
				x--
			}
		case '┌':
			if y+1 == prev.y {
				x++
			} else {
				y++
			}
		default:
			panic("oops")
		}
		steps++
		prev = pos
	}
	fmt.Println(steps/2 + 1)

	nrows, ncols := size(field)
	ins := 0
	for y := 0; y < nrows; y++ {
		up, do := false, false
		for x := 0; x < ncols; x++ {
			c := field[xy{x, y}]
			switch {
			case !loop[xy{x, y}]:
				if up {
					ins++
				}
			case c == '│':
				up = !up
				do = !do
			case c == '─':
			case c == '└':
				up = !up
			case c == '┘':
				up = !up
			case c == '┐':
				do = !do
			case c == '┌' || c == 'S': // hardcoded
				do = !do
			default:
				panic("oops")
			}
		}
	}
	fmt.Println(ins)
}
