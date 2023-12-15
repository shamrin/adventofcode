package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("15.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	// input = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	sum := 0
	for _, step := range strings.Split(input, ",") {
		hash := 0
		for _, c := range step {
			hash += int(c)
			hash *= 17
			hash %= 256
		}
		sum += hash
	}
	fmt.Println(sum)
}
