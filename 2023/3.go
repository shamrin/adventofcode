package main

import (
	"fmt"
	"strings"
)

type Cell struct{
	r int
	c int
}

type S map[Cell]rune

// type Connection struct {
// 	number Cell
// 	symbol Cell
// }

func isSymbol(r rune) bool {
	result := !isDigit(r) && r != '.'
	// fmt.Printf("  isSymbol(%s) => %t\n", string(r), result)
	return result
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func adjacentToSymbol(s S, r, c int) bool {
	for _, i := range []int{-1, 0, +1} {
		for _, j := range []int{-1, 0, +1} {
			if i != 0 || j != 0 {
				if isSymbol(s[Cell{r+i, c+j}]) {
					// fmt.Printf("adjacentToSymbol(%d, %d) => true\n", r, c)
					return true
				}
			}
		}
	}
	// fmt.Printf("adjacentToSymbol(%d, %d) => false\n", r, c)
	return false
}

func main() {
	// file, err := os.ReadFile("3.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// input := string(file)
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	// connections := map[Connection]bool{}

	sum := 0
	nrows := 0
	ncols := 0
	s := S{}
	for r, line := range strings.Split(input, "\n") {
		nrows = max(r+1, nrows)
		for c, ch := range line {
			ncols = max(c+1, ncols)
			s[Cell{r, c}] = ch
		}
		s[Cell{r, -1}] = '.'
		s[Cell{r, ncols}] = '.'
	}
	for c := -1; c <= ncols; c++ {
		s[Cell{-1, c}] = '.'
		s[Cell{nrows, c}] = '.'
	}

	for r := 0; r < nrows; r++ {
		n := -1
		// var number Cell
		keep := false
		for c := 0; c < ncols; c++ {
			ch := s[Cell{r, c}]
			if isDigit(ch) {
				if adjacentToSymbol(s, r, c) {
					// connections[Connection{number, Cell{r,c}}]
					keep = true
				}
				if n == -1 {
					// number = Cell{r, c}
					n = 0
				}
				n = n*10 + int(ch - '0')
			}
			if !isDigit(ch) || c == ncols - 1 {
				if n != -1 {
					if keep {
						sum += n
					}
					n = -1
				}
				keep = false
			}
		}
	}

	fmt.Println(sum)
}
