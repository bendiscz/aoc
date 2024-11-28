package main

import (
	"encoding/json"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2015
	day  = 12
)

func main() {
	Run(year, day, "", solve)
}

func count(n any, skipRed bool) int {
	if x, ok := n.(float64); ok {
		return int(x)
	} else if x, ok := n.([]any); ok {
		sum := 0
		for _, e := range x {
			sum += count(e, skipRed)
		}
		return sum
	} else if x, ok := n.(map[string]any); ok {
		sum := 0
		for _, e := range x {
			if skipRed {
				if s, ok := e.(string); ok && s == "red" {
					return 0
				}
			}
			sum += count(e, skipRed)
		}
		return sum
	}
	return 0
}

func solve(p *Problem) {
	p.NextLine()
	var doc any
	_ = json.Unmarshal(p.LineBytes(), &doc)
	p.PartOne(count(doc, false))
	p.PartTwo(count(doc, true))
}
