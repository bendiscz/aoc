package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 25
	example = `

1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122

`
)

func main() {
	Run(year, day, example, solve)
}

const digits = "=-012"

func decode(s string) int {
	n := 0
	for i, b := len(s)-1, 1; i >= 0; i, b = i-1, b*5 {
		d := strings.IndexByte(digits, s[i]) - 2
		n += b * d
	}
	return n
}

func encode(n int) string {
	var s []byte
	for b := 1; n > 0; b *= 5 {
		d := (n + 2) % 5
		s = append(s, digits[d])
		n = (n + 2) / 5
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	if len(s) == 0 {
		s = append(s, '0')
	}
	return string(s)
}

func solve(p *Problem) {
	sum := 0
	for p.NextLine() {
		sum += decode(p.Line())
	}

	p.PartOne(encode(sum))
	p.PartTwo("Merry Christmas!")
}
