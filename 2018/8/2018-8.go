package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 8
	example = `

2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	n := ParseInts(p.ReadLine())

	sum1, _ := part1(n)
	p.PartOne(sum1)

	sum2, _ := part2(n)
	p.PartTwo(sum2)
}

func part1(n []int) (sum, end int) {
	nc, nm := n[0], n[1]
	p := 2
	for i := 0; i < nc; i++ {
		s, e := part1(n[p:])
		sum += s
		p += e
	}
	for i := 0; i < nm; i++ {
		sum += n[p+i]
	}
	return sum, p + nm
}

func part2(n []int) (sum, end int) {
	nc, nm := n[0], n[1]
	p, c := 2, make([]int, nc)
	for i := 0; i < nc; i++ {
		s, e := part2(n[p:])
		c[i] = s
		p += e
	}
	if nc == 0 {
		for i := 0; i < nm; i++ {
			sum += n[p+i]
		}
	} else {
		for i := 0; i < nm; i++ {
			if j := n[p+i] - 1; j >= 0 && j < nc {
				sum += c[j]
			}
		}
	}
	return sum, p + nm
}
