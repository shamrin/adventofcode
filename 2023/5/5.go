package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Integer interface {
	int | uint32
}

func ints[T Integer](text string) []T {
	ints := []T{}
	for _, m := range regexp.MustCompile(`-?[0-9]+`).FindAllString(text, -1) {
		n, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
		ints = append(ints, T(n))
	}
	return ints
}

type R struct{ start, end uint32 }

func intersection(a, b R) []R {
	start := max(a.start, b.start)
	end := min(a.end, b.end)
	if start < end {
		return []R{{start, end}}
	}
	return []R{}
}

func difference(a, b R) []R {
	c := intersection(a, b)
	r := []R{}
	if len(c) == 0 {
		r = append(r, a)
	} else {
		left := R{a.start, c[0].start}
		if left.start < left.end {
			r = append(r, left)
		}
		right := R{c[0].end, a.end}
		if right.start < right.end {
			r = append(r, right)
		}
	}
	return r
}

func main() {
	var input string

	// file, err := os.ReadFile("5ex.txt")
	file, err := os.ReadFile("5.txt")
	if err != nil {
		panic(err)
	}
	input = string(file)

	lines := strings.Split(input, "\n\n")
	specsInput := ints[uint32](lines[0])
	rs := map[R]bool{}

	sum := uint32(0)
	for j := 0; j < len(specsInput)-1; j += 2 {
		sum += specsInput[j+1]
	}
	seeds := make([]uint32, 0, sum)
	for j := 0; j < len(specsInput)-1; j += 2 {
		rs[R{specsInput[j], specsInput[j] + specsInput[j+1]}] = true
		for s := specsInput[j]; s < specsInput[j]+specsInput[j+1]; s++ {
			seeds = append(seeds, s)
		}
	}

	M := [][][]uint32{}
	for i, block := range lines[1:] {
		ranges := strings.Split(block, "\n")
		M = append(M, [][]uint32{})
		for _, range_ := range ranges[1:] {
			M[i] = append(M[i], ints[uint32](range_))
		}
	}

	nw := 4
	for _, block := range M {
		newRs := []R{}
		oldRs := map[R]bool{}
		for _, m := range block {
			for r := range rs {
				dst, src, length := m[0], m[1], m[2]
				srcR := R{src, src + length}
				if inter := intersection(r, srcR); len(inter) > 0 {
					delete(rs, r)
					newRs = append(newRs, R{inter[0].start + dst - src, inter[0].end + dst - src})
					for _, rest := range difference(r, srcR) {
						oldRs[rest] = true
					}
				}
			}
		}
		for _, m := range block {
			for r := range oldRs {
				dst, src, length := m[0], m[1], m[2]
				srcR := R{src, src + length}
				if inter := intersection(r, srcR); len(inter) > 0 {
					delete(oldRs, r)
					newRs = append(newRs, R{inter[0].start + dst - src, inter[0].end + dst - src})
					for _, rest := range difference(r, srcR) {
						oldRs[rest] = true
					}
				}
			}
		}
		for _, r := range newRs {
			rs[r] = true
		}
		for r := range oldRs {
			rs[r] = true
		}

		var wg sync.WaitGroup
		for w := 0; w < nw; w++ {
			wg.Add(1)
			go func(w int, block [][]uint32) {
				defer wg.Done()
				for i := w; i < len(seeds); i += nw {
					for _, m := range block {
						dst, src, length := m[0], m[1], m[2]
						if src <= seeds[i] && seeds[i] < src+length {
							seeds[i] += dst - src
							break
						}
					}
				}
			}(w, block)
		}
		wg.Wait()
	}
	fmt.Println("brute", slices.Min(seeds))
	starts := []uint32{}
	for r := range rs {
		starts = append(starts, r.start)
	}
	fmt.Println("fast", slices.Min(starts))
}
