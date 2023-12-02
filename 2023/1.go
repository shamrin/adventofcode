package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	def := []struct {
		n int
		s string
	}{
		{n: 1, s: "one"},
		{n: 2, s: "two"},
		{n: 3, s: "three"},
		{n: 4, s: "four"},
		{n: 5, s: "five"},
		{n: 6, s: "six"},
		{n: 7, s: "seven"},
		{n: 8, s: "eight"},
		{n: 9, s: "nine"},
	}

	char := regexp.MustCompile("[a-zA-Z]+")
	lines := strings.Split(input, "\n")
	sum := 0
	for _, originalLine := range lines {
		line := strings.Clone(originalLine)
		for _, d := range def {
			line = strings.ReplaceAll(line, d.s, d.s+strconv.Itoa(d.n)+d.s)
		}
		onlyDigits := char.ReplaceAllString(line, "")
		var number int
		number, err := strconv.Atoi(string(onlyDigits[0]) + string(onlyDigits[len(onlyDigits)-1]))
		if err != nil {
			panic(err)
		}
		sum += number
	}
	fmt.Println(sum)
}
