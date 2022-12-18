package main

import (
	"unicode"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 6
	example = `

3,4,3,1,2

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	count80, count256 := 0, 0
	for _, c := range p.ReadAll() {
		if unicode.IsDigit(c) {
			count80 += dayCount[80-ParseInt(string(c))]
			count256 += dayCount[256-ParseInt(string(c))]
		}
	}

	p.PartOne(count80)
	p.PartTwo(count256)
}

var dayCount = simFish(256)

func simFish(days int) []int {
	day := make([]int, days+1)
	fish := []int{0, 1, 0, 0, 0, 0, 0, 0, 0}

	for d := 0; d <= days; d++ {
		born := fish[0]
		count := 2 * born
		for i := 1; i < len(fish); i++ {
			count += fish[i]
			fish[i-1] = fish[i]
		}
		fish[6] += born
		fish[8] = born
		day[d] = count
	}

	return day
}
