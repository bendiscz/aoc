package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2015
	day  = 2
)

func main() {
	Run(year, day, "", solve)
}

func parse(line string) (int, int, int) {
	i1, i2 := strings.IndexByte(line, 'x'), strings.LastIndexByte(line, 'x')
	return ParseInt(line[:i1]), ParseInt(line[i1+1 : i2]), ParseInt(line[i2+1:])
}

func solve(p *Problem) {
	paper, ribbon := 0, 0
	for p.NextLine() {
		l, w, h := parse(p.Line())
		a, b, c := l*w, w*h, h*l
		paper += a*2 + b*2 + c*2 + min(a, b, c)

		x, y, z := 2*(l+w), 2*(w+h), 2*(h+l)
		ribbon += min(x, y, z) + l*w*h
	}

	p.PartOne(paper)
	p.PartTwo(ribbon)
}
