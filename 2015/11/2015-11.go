package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2015
	day  = 11
)

func main() {
	Run(year, day, "", solve)
}

func inc(s []byte) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'z' {
			s[i] = 'a'
		} else {
			s[i]++
			break
		}
	}
}

func validate(s []byte) bool {
	triplet := false
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' || s[i] == 'o' || s[i] == 'l' {
			return false
		}
		if i >= 2 && s[i-2] == s[i]-2 && s[i-1] == s[i]-1 {
			triplet = true
		}
	}

	if !triplet {
		return false
	}

	pairs := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			pairs++
			i++
		}
	}

	if pairs < 2 {
		return false
	}

	return true
}

func solve(p *Problem) {
	p.NextLine()
	s := p.LineBytes()

	for {
		inc(s)
		if validate(s) {
			break
		}
	}
	p.PartOne(string(s))

	for {
		inc(s)
		if validate(s) {
			break
		}
	}
	p.PartTwo(string(s))
}
