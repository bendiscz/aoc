package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 2
	example = `

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

`
)

func main() {
	Run(year, day, example, solve)
}

type pull struct {
	r, g, b int
}

func parseGame(s string) (id int, p pull) {
	s = strings.TrimPrefix(s, "Game ")
	i := strings.IndexByte(s, ':')
	id = ParseInt(s[:i])
	for _, s1 := range strings.FieldsFunc(s[i+1:], func(r rune) bool { return r == ',' || r == ';' }) {
		if x, ok := parseColor(s1, "red"); ok {
			p.r = max(p.r, x)
		}
		if x, ok := parseColor(s1, "green"); ok {
			p.g = max(p.g, x)
		}
		if x, ok := parseColor(s1, "blue"); ok {
			p.b = max(p.b, x)
		}
	}
	return
}

func parseColor(s, c string) (int, bool) {
	if !strings.HasSuffix(s, c) {
		return 0, false
	}
	return ParseInt(strings.TrimSpace(s[:len(s)-len(c)])), true
}

func solve(p *Problem) {
	sum1, sum2 := 0, 0
	for p.NextLine() {
		id, p := parseGame(p.Line())
		if p.r <= 12 && p.g <= 13 && p.b <= 14 {
			sum1 += id
		}
		sum2 += p.r * p.g * p.b
	}
	p.PartOne(sum1)
	p.PartOne(sum2)
}
