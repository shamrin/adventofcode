package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cards = "  23456789TJQKA"

type Hand struct {
	cards [5]int
	bid   int
}

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func handType(hand Hand) int {
	labels := map[int]int{}
	for _, c := range hand.cards {
		labels[c]++
	}
	counts := []int{}
	for _, count := range labels {
		counts = append(counts, count)
	}
	switch len(counts) {
	case 1:
		return fiveOfAKind
	case 2:
		switch {
		case slices.Contains(counts, 4):
			return fourOfAKind
		case slices.Contains(counts, 3):
			return fullHouse
		default:
			panic("oops")
		}
	case 3:
		switch {
		case slices.Contains(counts, 3):
			return threeOfAKind
		case slices.Contains(counts, 2):
			return twoPair
		default:
			panic("oops")
		}
	case 4:
		switch {
		case slices.Contains(counts, 2):
			return onePair
		default:
			panic("oops")
		}
	case 5:
		return highCard
	default:
		panic("oops")
	}
}

func main() {
	// file, err := os.ReadFile("7ex.txt")
	file, err := os.ReadFile("7.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	lines := strings.Split(input, "\n")
	hands := []Hand{}
	for _, line := range lines {
		spec, bidS, _ := strings.Cut(line, " ")
		bid, err := strconv.Atoi(bidS)
		if err != nil {
			panic(err)
		}
		hand := Hand{bid: bid}
		for i, r := range spec {
			hand.cards[i] = strings.IndexRune(cards, r)
		}
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		c := cmp.Compare(handType(a), handType(b))
		if c != 0 {
			return c
		}
		for i := range a.cards {
			if a.cards[i] != b.cards[i] {
				return cmp.Compare(a.cards[i], b.cards[i])
			}
		}
		panic("oops")
	})

	sum := 0
	for rank, hand := range hands {
		sum += (rank + 1) * hand.bid
	}

	fmt.Println(sum)
}
