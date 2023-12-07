package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type Hand struct {
	cards [5]int
	bid   int
}

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

var cards = "  23456789TJQKA"
var cardsWithJoker = " J23456789TQKA"

func strengthen(hand Hand) Hand {
	labels := map[int]int{}
	joker := strings.Index(cardsWithJoker, "J")
	maxCount, maxCard := 0, 0
	for _, c := range hand.cards {
		if c != joker {
			labels[c]++
			if labels[c] > maxCount {
				maxCount, maxCard = labels[c], c
			}
		}
	}
	for i, c := range hand.cards {
		if c == joker {
			hand.cards[i] = maxCard
		}
	}
	return hand
}

func solve(lines []string, cards string, handType func(hand Hand) int) int {
	hands := []Hand{}
	for _, line := range lines {
		var hand Hand
		var err  error
		spec, bid, _ := strings.Cut(line, " ")
		if hand.bid, err = strconv.Atoi(bid); err != nil {
			panic(err)
		}
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
	return sum
}

func main() {
	// file, err := os.ReadFile("7ex.txt")
	file, err := os.ReadFile("7.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	lines := strings.Split(input, "\n")
	fmt.Println(solve(lines, cards, handType))
	fmt.Println(solve(lines, cardsWithJoker, func(hand Hand) int { return handType(strengthen(hand)) }))
}
