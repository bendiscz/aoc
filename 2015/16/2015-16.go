package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2015
	day  = 16
)

func main() {
	Run(year, day, "", solve)
}

var analysis = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func checkAnalysis(m []string) bool {
	for i := 2; i < len(m); i += 2 {
		if x, ok := analysis[m[i]]; !ok || x != ParseInt(m[i+1]) {
			return false
		}
	}
	return true
}

func checkAnalysisErrata(m []string) bool {
	for i := 2; i < len(m); i += 2 {
		x, ok := analysis[m[i]]
		if !ok {
			return false
		}

		y := ParseInt(m[i+1])
		switch m[i] {
		case "cats", "trees":
			if y <= x {
				return false
			}
		case "pomeranians", "goldfish":
			if y >= x {
				return false
			}
		default:
			if y != x {
				return false
			}
		}
	}
	return true
}

func solve(p *Problem) {
	var sue1, sue2 int
	for p.NextLine() {
		m := p.Parse(`^Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)$`)
		sue := ParseInt(m[1])
		if checkAnalysis(m) {
			sue1 = sue
		}
		if checkAnalysisErrata(m) {
			sue2 = sue
		}
	}

	p.PartOne(sue1)
	p.PartTwo(sue2)
}
