package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"slices"
)

const (
	year    = 2018
	day     = 14
	example = `

9

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	n := ParseInt(p.ReadLine())
	b := []byte{3, 7}
	e1, e2 := 0, 1

	for len(b) <= n+10 {
		b, e1, e2 = generateRecipes(b, e1, e2)
	}

	s1 := [10]byte{}
	for i := range s1 {
		s1[i] = '0' + b[n+i]
	}
	p.PartOne(string(s1[:]))

	pattern := []byte(nil)
	for x := n; x > 0; x /= 10 {
		pattern = append(pattern, byte(x%10))
	}
	slices.Reverse(pattern)

	for i, found := 0, false; !found; {
		for i+len(pattern) <= len(b) {
			if slices.Equal(pattern, b[i:i+len(pattern)]) {
				p.PartTwo(i)
				found = true
				break
			}
			i++
		}
		b, e1, e2 = generateRecipes(b, e1, e2)
	}
}

func generateRecipes(b []byte, e1, e2 int) ([]byte, int, int) {
	b1, b2 := b[e1], b[e2]
	r := b1 + b2
	if r > 9 {
		b = append(b, r/10)
	}
	b = append(b, r%10)
	e1 = (e1 + int(b1) + 1) % len(b)
	e2 = (e2 + int(b2) + 1) % len(b)
	return b, e1, e2
}
