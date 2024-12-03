package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"regexp"
)

const (
	year    = 2024
	day     = 3
	example = `

xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	s1, s2 := 0, 0
	s, enabled := p.ReadAll(), true

	for {
		i := re.FindStringSubmatchIndex(s)
		if i == nil {
			break
		}

		switch {
		case s[i[0]:i[1]] == "do()":
			enabled = true

		case s[i[0]:i[1]] == "don't()":
			enabled = false

		default:
			x := ParseInt(s[i[2]:i[3]]) * ParseInt(s[i[4]:i[5]])
			s1 += x
			if enabled {
				s2 += x
			}
		}

		s = s[i[1]:]
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}
