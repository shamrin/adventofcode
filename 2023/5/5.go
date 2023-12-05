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

func part1(M [][][]uint32, seeds []uint32) uint32 {
	x := slices.Clone(seeds)
	nw := 4
	for _, section := range M {
		var wg sync.WaitGroup
		for w := 0; w < nw; w++ {
			wg.Add(1)
			go func(w int, section [][]uint32) {
				defer wg.Done()
				for i := w; i < len(x); i += nw {
					for _, m := range section {
						dst, src, length := m[0], m[1], m[2]
						if src <= x[i] && x[i] < src+length {
							x[i] += dst - src
							break
						}
					}
				}
			}(w, section)
		}
		wg.Wait()
	}
	return slices.Min(x)
}

func part2Brute(maps [][][]uint32, specs []uint32) uint32 {
	sum := uint32(0)
	for i := 0; i < len(specs)-1; i += 2 {
		sum += specs[i+1]
	}
	seeds := make([]uint32, 0, sum)
	for i := 0; i < len(specs)-1; i += 2 {
		for s := specs[i]; s < specs[i]+specs[i+1]; s++ {
			seeds = append(seeds, s)
		}
	}
	return part1(maps, seeds)
}

func part2(maps [][][]uint32, specs []uint32) uint32 {
	rs := []R{}
	for i := 0; i < len(specs)-1; i += 2 {
		rs = append(rs, R{specs[i], specs[i] + specs[i+1]})
	}
	for _, section := range maps {
		next := make([]R, 0)
		for len(rs) > 0 {
			handled := false
			var r R
			r, rs = rs[0], rs[1:]
			for _, m := range section {
				dst, src, length := m[0], m[1], m[2]
				srcR := R{src, src + length}
				if inter := intersection(r, srcR); len(inter) > 0 {
					next = append(next, R{inter[0].start + dst - src, inter[0].end + dst - src})
					rs = append(rs, difference(r, srcR)...)
					handled = true
					break
				}
			}
			if !handled {
				next = append(next, r)
			}
		}
		rs = next
	}
	starts := []uint32{}
	for _, r := range rs {
		starts = append(starts, r.start)
	}
	return slices.Min(starts)
}

func main() {
	// file, err := os.ReadFile("5ex.txt")
	file, err := os.ReadFile("5.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	lines := strings.Split(input, "\n\n")
	specs := ints[uint32](lines[0])

	maps := [][][]uint32{}
	for i, section := range lines[1:] {
		ranges := strings.Split(section, "\n")
		maps = append(maps, [][]uint32{})
		for _, range_ := range ranges[1:] {
			maps[i] = append(maps[i], ints[uint32](range_))
		}
	}

	fmt.Println(part1(maps, specs))
	fmt.Println(part2(maps, specs))
	// fmt.Println(part2Brute(maps, specs))
}
