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
	for i := 0; i < len(s); i++ {
		n = n*5 + strings.IndexByte(digits, s[i]) - 2
	}
	return n
}

func encode(n int) string {
	if n == 0 {
		return "0"
	}

	s := [30]byte{}
	i := len(s) - 1
	for ; n > 0; i-- {
		s[i] = digits[(n+2)%5]
		n = (n + 2) / 5
	}
	return string(s[i:len(s)])
}

func solve(p *Problem) {
	sum := 0
	for p.NextLine() {
		sum += decode(p.Line())
	}

	p.PartOne(encode(sum))
	p.PartTwo("Merry Christmas!")
}
