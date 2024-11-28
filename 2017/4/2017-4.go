package main

import (
	"sort"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 4
	example = `

aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	sum1, sum2 := 0, 0
	for p.NextLine() {
		if solveLine(p.Line(), false) {
			sum1++
		}
		if solveLine(p.Line(), true) {
			sum2++
		}
	}
	p.PartOne(sum1)
	p.PartTwo(sum2)
}

func solveLine(s string, anagram bool) bool {
	words := map[string]bool{}
	for _, w := range SplitFields(s) {
		if anagram {
			w = sortWord(w)
		}
		if words[w] {
			return false
		}
		words[w] = true
	}
	return true
}

func sortWord(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}
