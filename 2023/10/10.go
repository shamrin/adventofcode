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
	for y := 0; y < nrows; y++ {
		for x := 0; x < ncols; x++ {
			fmt.Print(string(field[xy{x, y}]))
		}
		fmt.Println()
	}
}

func next(field map[xy]rune, pos xy, prev xy) xy {
	switch field[pos] {
	case '│':
		if pos.y-1 == prev.y {
			pos.y++
		} else {
			pos.y--
		}
	case '─':
		if pos.x-1 == prev.x {
			pos.x++
		} else {
			pos.x--
		}
	case '└':
		if pos.y-1 == prev.y {
			pos.x++
		} else {
			pos.y--
		}
	case '┘':
		if pos.x-1 == prev.x {
			pos.y--
		} else {
			pos.x--
		}
	case '┐':
		if pos.x-1 == prev.x {
			pos.y++
		} else {
			pos.x--
		}
	case '┌':
		if pos.y+1 == prev.y {
			pos.x++
		} else {
			pos.y++
		}
	}
	return pos
}

func main() {
	file, err := os.ReadFile("10.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	lines := strings.Split(input, "\n")
	field := map[xy]rune{}
	start := xy{}
	for y, line := range lines {
		x := 0
		for _, c := range line {
			field[xy{x, y}] = boxchars[c]
			if c == 'S' {
				start = xy{x, y}
			}
			x++
		}
	}

	// set start element
	west := strings.ContainsRune("─└┌", field[xy{start.x - 1, start.y}])
	east := strings.ContainsRune("┘┐─", field[xy{start.x + 1, start.y}])
	north := strings.ContainsRune("│┐┌", field[xy{start.x, start.y - 1}])
	south := strings.ContainsRune("│┘└", field[xy{start.x, start.y + 1}])
	switch {
	case west && north:
		field[start] = '─'
	case west && east:
		field[start] = '┘'
	case west && south:
		field[start] = '┐'
	case north && east:
		field[start] = '└'
	case north && south:
		field[start] = '│'
	case east && south:
		field[start] = '┌'
	}

	// walk the loop and count steps (part 1)
	steps := 0
	loop := map[xy]bool{start: true}
	for prev, pos := start, next(field, start, xy{}); pos != start; prev, pos = pos, next(field, pos, prev) {
		loop[pos] = true
		steps++
	}

	// erase all non-loop elements
	for pos := range field {
		if !loop[pos] {
			field[pos] = '.'
		}
	}

	// count inner elements (part 2)
	nrows, ncols := size(field)
	ins := 0
	for y := 0; y < nrows; y++ {
		topIn, bottomIn := false, false
		for x := 0; x < ncols; x++ {
			switch field[xy{x, y}] {
			case '.':
				if topIn {
					ins++
					field[xy{x, y}] = 'I'
				}
			case '│':
				topIn = !topIn
				bottomIn = !bottomIn
			case '└', '┘':
				topIn = !topIn
			case '┐', '┌':
				bottomIn = !bottomIn
			}
		}
	}

	print(field)
	fmt.Println(steps/2 + 1)
	fmt.Println(ins)
}
