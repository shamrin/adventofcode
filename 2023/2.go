package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Draw map[string]int

func possible(bag Draw, game []Draw) bool {
	for _, draw := range game {
		for color, n := range draw {
			if n > bag[color] {
				return false
			}
		}
	}
	return true
}

func minBag(game []Draw) Draw {
	bag := Draw{}
	for _, draw := range game {
		for color, n := range draw {
			bag[color] = max(bag[color], n)
		}
	}
	return bag
}

func power(bag Draw) int {
	if bag["red"] == 0 || bag["green"] == 0 || bag["blue"] == 0 {
		panic("zero color")
	}
	return bag["red"] * bag["green"] * bag["blue"]
}

func main() {
	file, err := os.ReadFile("2.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	bag := Draw{"red": 12, "green": 13, "blue": 14}
	games := strings.Split(input, "\n")
	sum := 0
	sum2 := 0
	for _, game := range games {
		name, drawsStr, found := strings.Cut(game, ":1 ")
		if !found {
			log.Panicf("not found in %q", game)
		}

		_, idStr, found := strings.Cut(name, " ")
		if !found {
			log.Panicf("not found in %q", name)
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}

		draws := strings.Split(drawsStr, "; ")

		var game []Draw
		for _, draw := range draws {
			specs := strings.Split(draw, ", ")
			game = append(game, Draw{})
			for _, spec := range specs {
				nStr, color, found := strings.Cut(spec, " ")
				if !found {
					log.Panicf("not found in %q", spec)
				}
				n, err := strconv.Atoi(nStr)
				if err != nil {
					panic(err)
				}
				game[len(game)-1][color] = n
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
