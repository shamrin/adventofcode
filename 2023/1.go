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

	numbers := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}
	lines := strings.Split(input, "\n")
	sum := 0
	for _, originalLine := range lines {
		line := strings.Clone(originalLine)
		for s, n := range numbers {
			line = strings.ReplaceAll(line, s, s+strconv.Itoa(n)+s)
		}
		digits := regexp.MustCompile("[a-zA-Z]+").ReplaceAllString(line, "")
		number, err := strconv.Atoi(digits[:1] + digits[len(digits)-1:])
		if err != nil {
			panic(err)
		}
		sum += number
	}
	fmt.Println(sum)
}
