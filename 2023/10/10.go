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

	// walk the loop (part 1)
	steps := 0
	loop := map[xy]bool{}
	prev := xy{-2, -2}
	x, y := start.x, start.y
	for {
		pos := xy{x, y}
		loop[pos] = true
		switch field[pos] {
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
		}
		if start.x == x && start.y == y {
			break
		}
		steps++
		prev = pos
	}
	fmt.Println(steps/2 + 1)

	// erase all non-loop elements
	for pos := range field {
		if !loop[pos] {
			field[pos] = '.'
		}
	}
	// print(field)

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
	fmt.Println(ins)
}
