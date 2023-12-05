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
type Map struct {
	dst uint32
	src R
}

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

func part1(sections [][]Map, seeds []uint32) uint32 {
	nw := 4
	for _, maps := range sections {
		var wg sync.WaitGroup
		for w := 0; w < nw; w++ {
			wg.Add(1)
			go func(w int, maps []Map) {
				defer wg.Done()
				for i := w; i < len(seeds); i += nw {
					for _, m := range maps {
						if m.src.start <= seeds[i] && seeds[i] < m.src.end {
							seeds[i] += m.dst - m.src.start
							break
						}
					}
				}
			}(w, maps)
		}
		wg.Wait()
	}
	return slices.Min(seeds)
}

func part2Brute(sections [][]Map, specs []uint32) uint32 {
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
	return part1(sections, seeds)
}

func part2(sections [][]Map, specs []uint32) uint32 {
	rs := []R{}
	for i := 0; i < len(specs)-1; i += 2 {
		rs = append(rs, R{specs[i], specs[i] + specs[i+1]})
	}
	for _, maps := range sections {
		next := make([]R, 0)
		for len(rs) > 0 {
			handled := false
			r := rs[0]
			rs = rs[1:]
			for _, m := range maps {
				if inter := intersection(r, m.src); len(inter) > 0 {
					next = append(next, R{
						start: inter[0].start + m.dst - m.src.start,
						end:   inter[0].end + m.dst - m.src.start,
					})
					rs = append(rs, difference(r, m.src)...)
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
	in := ints[uint32](lines[0])

	sections := [][]Map{}
	for i, section := range lines[1:] {
		ranges := strings.Split(section, "\n")
		sections = append(sections, []Map{})
		for _, r := range ranges[1:] {
			m := ints[uint32](r)
			dst, src, length := m[0], m[1], m[2]
			sections[i] = append(sections[i], Map{dst, R{src, src + length}})
		}
	}

	fmt.Println(part1(sections, slices.Clone(in)))
	fmt.Println(part2(sections, in))
	// fmt.Println(part2Brute(sections, in))
}
