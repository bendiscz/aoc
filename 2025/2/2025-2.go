package main

import (
	"fmt"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 2
	example = `

11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	f := ParseInts(strings.ReplaceAll(p.ReadAll(), "-", "_"))
	for i := 0; i < len(f); i += 2 {
		for x := f[i]; x <= f[i+1]; x++ {
			if isInvalid1(x) {
				s1 += x
			}
			if isInvalid2(x) {
				s2 += x
			}
		}
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}

func isInvalid1(x int) bool {
	s := fmt.Sprintf("%d", x)
	h := len(s) / 2
	return s[:h] == s[h:]
}

func isInvalid2(x int) bool {
	s := fmt.Sprintf("%d", x)

loop:
	for l := 1; l <= len(s)/2; l++ {
		if len(s)%l != 0 {
			continue
		}
		sub := s[:l]
		for i := l; i+l <= len(s); i += l {
			if sub != s[i:i+l] {
				continue loop
			}
		}
		return true
	}
	return false
}
