package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Spec struct {
	n int
	color string
}

type Bag struct {
	red int
	green int
	blue int
}

func possible(bag Bag, game [][]Spec) bool {
	for _, draw := range(game) {
		for _, spec := range(draw) {
			switch spec.color {
			case "red":
				if spec.n > bag.red {
					return false
				}
			case "green":
				if spec.n > bag.green {
					return false
				}
			case "blue":
				if spec.n > bag.blue {
					return false
				}
			default:
				panic("unknown color")
			}
		}
	}
	return true
}

func minBag(game [][]Spec) Bag {
	var bag Bag
	for _, draw := range(game) {
		for _, spec := range(draw) {
			switch spec.color {
			case "red":
				bag.red = max(bag.red, spec.n)
			case "green":
				bag.green = max(bag.green, spec.n)
			case "blue":
				bag.blue = max(bag.blue, spec.n)
			default:
				panic("unknown color")
			}
		}
	}
	return bag
}

func power(bag Bag) int {
	return bag.red * bag.blue * bag.green
}

func main() {
	file, err := os.ReadFile("2.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	bag := Bag{red: 12, green: 13, blue: 14}

	games := strings.Split(input, "\n")
	sum := 0
	sum2 := 0
	for _, game := range games {
		name, drawsStr, found := strings.Cut(game, ": ")
		if !found {
			panic(fmt.Sprintf("not found in %q", game))
		}

		_, idStr, found := strings.Cut(name, " ")
		if !found {
			panic(fmt.Sprintf("not found in %q", name))
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}

		draws := strings.Split(drawsStr, "; ")

		var game [][]Spec
		for _, draw := range draws {
			specs := strings.Split(draw, ", ")
			game = append(game, make([]Spec, 0))
			for _, spec := range specs {
				nStr, color, found := strings.Cut(spec, " ")
				if !found {
					panic(fmt.Sprintf("not found in %q", spec))
				}
				n, err := strconv.Atoi(nStr)
				if err != nil {
					panic(err)
				}
				game[len(game) - 1] = append(game[len(game) - 1], Spec{n: n, color: color})
			}
		}
		if possible(bag, game) {
			sum += id
		}
		sum2 += power(minBag(game))
	}
	fmt.Println(sum)	
	fmt.Println(sum2)	
}