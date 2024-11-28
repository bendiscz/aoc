package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 1
	example = `

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	sum1, sum2 := 0, 0
	for p.NextLine() {
		sum1 += parsePart1(p.Line())
		sum2 += parsePart2(p.Line())
	}
	p.PartOne(sum1)
	p.PartTwo(sum2)
}

func parsePart1(s string) int {
	d1, d2, i := 0, 0, 0
	for i < len(s) {
		if ch := s[i]; ch > '0' && ch <= '9' {
			d1 = int(ch - '0')
			d2 = d1
			break
		}
		i++
	}

	for j := len(s) - 1; j > i; j-- {
		if ch := s[j]; ch > '0' && ch <= '9' {
			d2 = int(ch - '0')
			break
		}
	}

	return 10*d1 + d2
}

func parsePart2(s string) int {
	d1, d2, i := 0, 0, 0
	for i < len(s) {
		if d, ok := digit(s[i:]); ok {
			d1 = d
			d2 = d
			break
		}
		i++
	}

	for j := len(s) - 1; j > i; j-- {
		if d, ok := digit(s[j:]); ok {
			d2 = d
			break
		}
	}

	return 10*d1 + d2
}

var digits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func digit(s string) (int, bool) {
	if ch := s[0]; ch > '0' && ch <= '9' {
		return int(ch - '0'), true
	}
	for d, digit := range digits {
		if strings.HasPrefix(s, digit) {
			return d + 1, true
		}
	}
	return 0, false
}
