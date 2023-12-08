package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	left  string
	right string
}

func words(text string) []string {
	ss := []string{}
	for _, m := range regexp.MustCompile(`[0-9a-zA-Z]+`).FindAllString(text, -1) {
		ss = append(ss, m)
	}
	return ss
}

func main() {
	// file, err := os.ReadFile("7ex.txt")
	file, err := os.ReadFile("8.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	lines := strings.Split(input, "\n")
	instructions := lines[0]
	fmt.Println(instructions)
	net := map[string]Node{}
	for _, line := range lines[2:] {
		if len(line) == 0 { continue }
		ws := words(line)
		net[ws[0]] = Node{ws[1], ws[2]}
	}
	fmt.Println(net)

	node := "AAA"
	steps := 0
	for ; node != "ZZZ"; steps++ {
		dir := instructions[steps % len(instructions)]
		// fmt.Println(node, string(dir))
		switch dir {
		case 'L':
			node = net[node].left
		case 'R':
			node = net[node].right
		default:
			panic("oops")
		}
	}
	fmt.Println(steps)
}
