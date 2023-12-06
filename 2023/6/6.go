package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ints(text string) []int {
	ints := []int{}
	for _, m := range regexp.MustCompile(`-?[0-9]+`).FindAllString(text, -1) {
		n, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
		ints = append(ints, n)
	}
	return ints
}

func ways(time int, dist int) int {
	ways := 0
	for hold := 0; hold <= time; hold++ {
		if hold*(time-hold) > dist {
			ways++
		}
	}
	return ways
}

func main() {
	// file, err := os.ReadFile("6ex.txt")
	file, err := os.ReadFile("6.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	lines := strings.Split(input, "\n")

	times := ints(lines[0])
	dists := ints(lines[1])
	result := 1
	for i := 0; i < len(times); i++ {
		result *= ways(times[i], dists[i])
	}
	fmt.Println(result)

	time := ints(strings.ReplaceAll(lines[0], " ", ""))[0]
	dist := ints(strings.ReplaceAll(lines[0], " ", ""))[0]
	fmt.Println(ways(time, dist))
}
