package main

import (
	"math"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 6
	example = `

Time:      7  15   30
Distance:  9  40  200

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	ts := ParseInts(p.ReadLine())
	ds := ParseInts(p.ReadLine())

	r1 := 1
	for i := 0; i < len(ts); i++ {
		r1 *= count2(ts[i], ds[i])
	}
	p.PartOne(r1)

	p.Reset()
	t := ParseInts(strings.ReplaceAll(p.ReadLine(), " ", ""))[0]
	d := ParseInts(strings.ReplaceAll(p.ReadLine(), " ", ""))[0]
	p.PartTwo(count2(t, d))
}

func count2(t, d int) int {
	// x(t - x) > d
	// -x^2 + tx - d > 0

	b := float64(t)
	disc := math.Sqrt(b*b - 4.0*float64(d))
	t0, t1 := (b-disc)/2.0, (b+disc)/2.0
	return int(math.Ceil(t1)) - int(math.Floor(t0)) - 1
}
