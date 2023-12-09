package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	left  string
	right string
}

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

func zeroes(history []int) bool {
	for _, n := range history {
		if n != 0 {
			return false
		}
	}
	return true
}

func solve(lines []string, future bool) int {
	sum := 0
	for _, line := range lines {
		history := [][]int{}
		ns := ints(line)
		if future {
			slices.Reverse(ns)
		}
		history = append(history, ns)
		for !zeroes(history[len(history)-1]) {
			prev := history[len(history)-1]
			next := make([]int, 0, len(prev)-1)
			for i := range prev[:len(prev)-1] {
				next = append(next, prev[i]-prev[i+1])
			}
			history = append(history, next)
		}
		extra := make([]int, len(history))
		extra[len(extra)-1] = 0
		for i := len(extra) - 2; i >= 0; i-- {
			extra[i] = history[i][0] + extra[i+1]
		}
		sum += extra[0]
	}
	return sum
}

func main() {
	file, err := os.ReadFile("9.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	lines := strings.Split(input, "\n")
	fmt.Println(solve(lines, true))
	fmt.Println(solve(lines, false))
}
