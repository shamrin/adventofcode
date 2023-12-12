package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
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

func powerOfTwo(n int) int {
	return 1 << n
}

func main() {
	file, err := os.ReadFile("12.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
// 	input = `???.### 1,1,3
// .??..??...?##. 1,1,3
// ?#?#?#?#?#?#?#? 1,3,1,6
// ????.#...#... 4,1,1
// ????.######..#####. 1,6,5
// ?###???????? 3,2,1`

	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row, d, _ := strings.Cut(line, " ")
		damaged := ints(d)
		qn := strings.Count(row, "?")
		n := powerOfTwo(qn)
		s := []byte(row)
		for i := 0; i < n; i++ {
			j := i
			for k := 0; k < len(row); k++ {
				if row[k] == '?' {
					if j%2 == 0 {
						s[k] = '.'
					} else {
						s[k] = '#'
					}
					j /= 2
				}
			}
			cc := make([]int, 0, 32)
			counting := false
			for k := 0; k < len(s); k++ {
				switch s[k] {
				case '.':
					counting = false
				case '#':
					if !counting {
						cc = append(cc, 1)
					} else {
						cc[len(cc)-1]++
					}
					counting = true
				}
			}
			if slices.Equal(cc, damaged) {
				sum += 1
			}
		}
	}
	fmt.Println(sum,"maxQn", maxQn, "maxDamagedSum", maxDamagedSum)
}
