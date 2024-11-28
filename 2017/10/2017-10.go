package main

import (
	"regexp"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 10
	example = `

flqrgnkx-0
AoC 2017

1,2,3
1,2,4

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		p.PartOne(part1(p.Line()))
		p.PartTwo(part2(p.Line()))
	}
}

func part1(s string) int {
	if !regexp.MustCompile(`^[\d ,]+$`).MatchString(s) {
		return 0
	}

	list := makeList(256)
	pos, skip := 0, 0
	round(list, ParseInts(s), &pos, &skip)
	return list[0] * list[1]
}

func part2(s string) string {
	list := makeList(256)
	lengths := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		lengths[i] = int(s[i])
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	pos, skip := 0, 0
	for i := 0; i < 64; i++ {
		round(list, lengths, &pos, &skip)
	}

	r := [32]byte{}
	for i := 0; i < 16; i++ {
		x := 0
		for j := 0; j < 16; j++ {
			x = x ^ list[16*i+j]
		}
		r[2*i], r[2*i+1] = Hex[x/16], Hex[x%16]
	}

	return string(r[:])
}

func makeList(n int) []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i
	}
	return list
}

func round(list, lengths []int, pos, skip *int) {
	n := len(list)
	for _, l := range lengths {
		for i := 0; i < l/2; i++ {
			i1, i2 := (*pos+i)%n, (*pos+l-i-1)%n
			list[i1], list[i2] = list[i2], list[i1]
		}
		*pos = (*pos + l + *skip) % n
		*skip++
	}
}
