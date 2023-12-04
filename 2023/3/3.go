package main

import (
	"fmt"
	"os"
	"strings"
)

type Cell struct {
	r int
	c int
}

type Schematic map[Cell]rune

type Connection struct {
	number Cell
	symbol Cell
}

func isSymbol(r rune) bool {
	return !isDigit(r) && r != '.' && r != 0
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func adjacentSymbols(s Schematic, r, c int) []Cell {
	symbols := []Cell{}
	for _, i := range []int{-1, 0, +1} {
		for _, j := range []int{-1, 0, +1} {
			if i != 0 || j != 0 {
				if isSymbol(s[Cell{r + i, c + j}]) {
					symbols = append(symbols, Cell{r + i, c + j})
				}
			}
		}
	}
	return symbols
}

func main() {
	file, err := os.ReadFile("3.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	s := Schematic{}
	rows := strings.Split(input, "\n")
	ncols := len(rows[0])
	for r, line := range rows {
		for c, ch := range line {
			s[Cell{r, c}] = ch
		}
	}

	connections := map[Connection]bool{}
	numbers := map[Cell]int{}
	for r := 0; r < len(rows); r++ {
		reading, cell := false, Cell{}
		for c := 0; c < ncols; c++ {
			ch := s[Cell{r, c}]
			if isDigit(ch) {
				if !reading {
					reading, cell = true, Cell{r, c}
				}
				for _, symbol := range adjacentSymbols(s, r, c) {
					connections[Connection{cell, symbol}] = true
				}
				numbers[cell] = numbers[cell]*10 + int(ch-'0')
			} else {
				reading = false
			}
		}
	}

	partNumbers := map[Cell]bool{}
	gears := map[Cell][]Cell{}
	for conn := range connections {
		partNumbers[conn.number] = true
		if s[conn.symbol] == '*' {
			gears[conn.symbol] = append(gears[conn.symbol], conn.number)
		}
	}

	sum := 0
	for cell := range partNumbers {
		sum += numbers[cell]
	}
	fmt.Println(sum)

	sum = 0
	for _, ns := range gears {
		if len(ns) == 2 {
			sum += numbers[ns[0]] * numbers[ns[1]]
		}
	}
	fmt.Println(sum)
}
