package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Lens struct {
	label string
	focal int
}

func hash(s string) int {
	hash := 0
	for _, c := range s {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}

func main() {
	file, err := os.ReadFile("15.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	// input = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	re := regexp.MustCompile(`([a-zA-Z]+)([=-])([1-9])?`)
	hSum := 0
	boxes := [256][]Lens{}
	for _, step := range strings.Split(input, ",") {
		hSum += hash(step)
		r := re.FindStringSubmatch(step)
		label, op := r[1], r[2]
		h := hash(label)
		i := slices.IndexFunc(boxes[h], func(lens Lens) bool {
			return lens.label == label
		})
		switch op {
		case "-":
			if i != -1 {
				boxes[h] = slices.Delete(boxes[h], i, i+1)
			}
		case "=":
			var focal int
			if focal, err = strconv.Atoi(r[3]); err != nil {
				panic(err)
			}
			if i == -1 {
				boxes[h] = append(boxes[h], Lens{label, focal})
			} else {
				boxes[h][i] = Lens{label, focal}
			}
		}
	}

	sum := 0
	for i, box := range boxes {
		if len(box) != 0 {
			for j, lens := range box {
				sum += (i + 1) * (j + 1) * lens.focal
			}
		}
	}

	fmt.Println(hSum)
	fmt.Println(sum)
}
