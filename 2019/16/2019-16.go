package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 16
	example = `

03036732577212944063491565474664

`
)

func main() {
	Run(year, day, example, solve)
}

var pattern = []int{1, 0, -1, 0}

type signal struct {
	v []int
	s []int
	o int
}

func parseSignal1(line string) *signal {
	s := signal{v: make([]int, len(line))}
	for i := 0; i < len(line); i++ {
		s.v[i] = int(line[i] - '0')
	}
	return &s
}

func parseSignal2(line string) *signal {
	s := signal{v: make([]int, len(line)*10_000)}
	s.o = ParseInt(line[:7])

	for i := 0; i < len(line); i++ {
		s.v[i] = int(line[i] - '0')
	}

	for i := s.o / len(line) * len(line); i < len(s.v); i += len(line) {
		copy(s.v[i:], s.v[:len(line)])
	}

	return &s
}

func (s *signal) sum(a, b int) int {
	a, b = min(a, len(s.v)), min(b, len(s.v))
	return s.s[b] - s.s[a]
}

func (s *signal) phase() {
	if s.o > len(s.v)/2 {
		x := 0
		for i := len(s.v) - 1; i >= s.o; i-- {
			x += s.v[i]
			s.v[i] = x % 10
		}
	} else {
		s.s = make([]int, len(s.v)+1)
		for i, sum := s.o, 0; i < len(s.v); i++ {
			sum += s.v[i]
			s.s[i+1] = sum
		}

		for i := s.o; i < len(s.v); i++ {
			x := 0
			for j, k := i, 0; j < len(s.v); j, k = j+i+1, (k+1)%len(pattern) {
				q := pattern[k]
				if q != 0 {
					x += q * s.sum(j, j+i+1)
				}

			}
			s.v[i] = Abs(x) % 10
		}
	}
}

func (s *signal) eight() string {
	b := [8]byte{}
	for i := 0; i < len(b); i++ {
		b[i] = byte('0' + s.v[i+s.o])
	}
	return string(b[:])
}

func solve(p *Problem) {
	for p.NextLine() {
		s1 := parseSignal1(p.Line())
		s2 := parseSignal2(p.Line())
		for n := 0; n < 100; n++ {
			s1.phase()
			s2.phase()
		}
		p.PartOne(s1.eight())
		p.PartTwo(s2.eight())
	}
}
