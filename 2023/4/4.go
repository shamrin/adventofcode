package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func IntPow(n, m int) int {
	if m < 0 {
		panic("negative power")
	}
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func Ints(text string) []int {
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

func main() {
	var input string

	file, err := os.ReadFile("4.txt")
	if err != nil {
		panic(err)
	}
	input = string(file)

	lines := strings.Split(input, "\n")
	counts := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		counts[i] = 1
	}

	sum := 0
	for i, line := range lines {
		win, have, _ := strings.Cut(line, "|")
		winning := map[int]bool{}
		for _, w := range Ints(win)[1:] {
			winning[w] = true
		}
		won := 0
		for _, h := range Ints(have) {
			if winning[h] {
				won++
			}
		}
		var points int
		if won == 0 {
			points = 0
		} else {
			points = IntPow(2, won-1)
		}
		sum += points
		for j := 0; j < won && i+j+1 < len(counts); j++ {
			counts[i+j+1] += counts[i]
		}
	}
	fmt.Println(sum)

	sum = 0
	for _, count := range counts {
		sum += count
	}
	fmt.Println(sum)
}
