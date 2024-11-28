package main

import (
	"regexp"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 9
	example = `

(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN

`
)

func main() {
	Run(year, day, example, solve)
}

var pattern = regexp.MustCompile(`\((\d+)x(\d+)\)`)

func decompress(part []byte, v2 bool) int {
	count := 0
	for len(part) > 0 {
		m := pattern.FindSubmatchIndex(part)

		if m == nil {
			count += len(part)
			break
		}

		l := ParseInt(string(part[m[2]:m[3]]))
		t := ParseInt(string(part[m[4]:m[5]]))

		d := l
		if v2 {
			d = decompress(part[m[5]+1:m[1]+l], true)
		}

		count += m[0] + d*t
		part = part[m[1]+l:]
	}
	return count
}

func solve(p *Problem) {
	text := []byte(strings.TrimSpace(p.ReadAll()))
	p.PartOne(decompress(text, false))
	p.PartTwo(decompress(text, true))
}
