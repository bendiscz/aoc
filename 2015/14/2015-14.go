package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 14
	example = `

Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.

`
)

func main() {
	Run(year, day, example, solve)
}

type deer struct {
	speed   int
	flying  int
	resting int

	distance int
	score    int
}

func (d *deer) flyingAt(t int) bool {
	return t%(d.flying+d.resting) < d.flying
}

func distance(time, speed, flying, resting int) int {
	d := time / (flying + resting)
	r := time % (flying + resting)
	if r > flying {
		r = flying
	}
	return speed * (flying*d + r)
}

func solve(p *Problem) {
	const TIME = 2503
	var deers []*deer
	best := 0
	for p.NextLine() {
		m := p.Parse(`^\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.$`)
		speed, flying, resting := ParseInt(m[1]), ParseInt(m[2]), ParseInt(m[3])
		deers = append(deers, &deer{
			speed:   speed,
			flying:  flying,
			resting: resting,
		})
		best = max(distance(TIME, speed, flying, resting), best)
	}
	p.PartOne(best)

	for t := 0; t < TIME; t++ {
		m := 0
		for _, d := range deers {
			if d.flyingAt(t) {
				d.distance += d.speed
			}
			m = max(d.distance, m)
		}
		for _, d := range deers {
			if d.distance == m {
				d.score++
			}
		}
	}

	best = 0
	for _, d := range deers {
		best = max(d.score, best)
	}
	p.PartTwo(best)
}
