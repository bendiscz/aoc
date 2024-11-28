package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 20
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	n := ParseInt(p.ReadLine())

	s1 := make([]int, n/10)
	s2 := make([]int, n/10)
	for i := 1; i < len(s1); i++ {
		count := 0
		for j := i; j < len(s1); j += i {
			s1[j] += 10 * i
			if count < 50 {
				s2[j] += 11 * i
				count++
			}
		}
	}

	for i := 1; i < len(s1); i++ {
		if s1[i] >= n {
			p.PartOne(i)
			break
		}
	}

	for i := 1; i < len(s1); i++ {
		if s2[i] >= n {
			p.PartTwo(i)
			break
		}
	}
}
