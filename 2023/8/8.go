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

func arrived(nodes []string) bool {
	for _, node := range nodes {
		if node[2] != 'Z' {
			return false
		}
	}
	return true
}

func main() {
	// file, err := os.ReadFile("7ex.txt")
	file, err := os.ReadFile("8.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	input = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	lines := strings.Split(input, "\n")
	instructions := lines[0]
	net := map[string]Node{}
	for _, line := range lines[2:] {
		if len(line) == 0 { continue }
		ws := words(line)
		net[ws[0]] = Node{ws[1], ws[2]}
	}

	nodes := []string{}
	for node := range net {
		if node[2] == 'A' {
			nodes = append(nodes, node)
		}
	}
	steps := 0
	for ; !arrived(nodes); steps++ {
		for i, node := range nodes {
			dir := instructions[steps % len(instructions)]
			switch dir {
			case 'L':
				nodes[i] = net[node].left
			case 'R':
				nodes[i] = net[node].right
			default:
				panic("oops")
			}
		}
	}
	fmt.Println(steps)
}
