package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"sync/atomic"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 5
	example = `

seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	in := parse(p)
	p.PartOne(part1(in))
	p.PartTwo(part2(in))
	//p.PartTwo(part2BruteForce2(in))
}

type mapping struct {
	d, s, n int
}

type input struct {
	seeds    []int
	mappings [][]mapping
}

func parse(p *Problem) *input {
	in := &input{}
	in.seeds = ParseInts(strings.TrimPrefix(p.ReadLine(), "seeds: "))
	p.NextLine()

	for p.NextLine() {
		m := []mapping(nil)
		for p.NextLine() {
			if p.Line() == "" {
				break
			}

			t := ParseInts(p.Line())
			m = append(m, mapping{t[0], t[1], t[2]})
		}
		in.mappings = append(in.mappings, m)
	}

	return in
}

func part1(in *input) int {
	src := make([]int, len(in.seeds))
	dst := make([]int, len(in.seeds))
	copy(src, in.seeds)

	for _, table := range in.mappings {
	loop:
		for i, x := range src {
			for _, m := range table {
				s, e, d := m.s, m.s+m.n, m.d-m.s
				if x >= s && x < e {
					dst[i] = x + d
					continue loop
				}
				dst[i] = src[i]
			}
		}
		src, dst = dst, src
	}

	r := math.MaxInt
	for _, x := range src {
		r = min(r, x)
	}
	return r
}

type interval struct {
	s, e int
}

func (i interval) String() string { return fmt.Sprintf("<%d;%d)", i.s, i.e) }

func (i interval) transform(m mapping) ([]interval, []interval) {
	if m.s+m.n <= i.s || m.s >= i.e {
		return []interval{i}, nil
	}

	r := []interval(nil)
	if m.s > i.s {
		r = append(r, interval{i.s, m.s})
	}
	if m.s+m.n < i.e {
		r = append(r, interval{m.s + m.n, i.e})
	}

	d := m.d - m.s
	return r, []interval{{max(i.s, m.s) + d, min(i.e, m.s+m.n) + d}}
}

func part2(in *input) int {
	src := []interval(nil)
	for i := 0; i < len(in.seeds); i += 2 {
		src = append(src, interval{in.seeds[i], in.seeds[i] + in.seeds[i+1]})
	}

	for _, table := range in.mappings {
		dst := []interval(nil)
		for _, m := range table {
			next := []interval(nil)
			for _, i := range src {
				r, t := i.transform(m)
				next = append(next, r...)
				dst = append(dst, t...)
			}
			src = next
		}
		src = append(src, dst...)
	}

	r := math.MaxInt
	for _, i := range src {
		r = min(r, i.s)
	}
	return r
}

func part2BruteForce2(in *input) int {
	r := atomic.Int64{}
	r.Store(math.MaxInt64)

	wg := &sync.WaitGroup{}
	for i := 0; i < len(in.seeds); i += 2 {
		const block = 100000000
		s, n := in.seeds[i], in.seeds[i+1]
		for n > 0 {
			b := min(n, block)

			wg.Add(1)
			go func(s, n int) {
				seeds := make([]int, n)
				for k := 0; k < n; k++ {
					seeds[k] = s + k
				}

				x := part1(&input{
					seeds:    seeds,
					mappings: in.mappings,
				})

				for {
					rx := r.Load()
					if r.CompareAndSwap(rx, min(rx, int64(x))) {
						break
					}
				}
				wg.Done()
			}(s, b)

			s += b
			n -= b
		}
	}

	wg.Wait()
	return int(r.Load())
}
