package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 19
	example = `

Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.

`
)

func main() {
	Run(year, day, example, solve)
}

type blueprint struct {
	id          int
	oreRobotOre int16
	claRobotOre int16
	obsRobotOre int16
	obsRobotCla int16
	geoRobotOre int16
	geoRobotObs int16
}

func (b *blueprint) evaluate(s state) int {
	if s.minute == 0 {
		return int(s.geo)
	}

	if s.ore >= b.geoRobotOre && s.obs >= b.geoRobotObs {
		s2 := s.next()
		s2.geoRobots++
		s2.ore -= b.geoRobotOre
		s2.obs -= b.geoRobotObs
		return b.evaluate(s2)
	}

	max := 0
	obsBuilt := false

	if s.ore >= b.obsRobotOre && s.cla >= b.obsRobotCla && s.obsRobots < b.geoRobotObs {
		s2 := s.next()
		s2.obsRobots++
		s2.ore -= b.obsRobotOre
		s2.cla -= b.obsRobotCla
		max = Max(max, b.evaluate(s2))
		obsBuilt = true
	}

	if s.ore >= b.claRobotOre && s.claRobots < b.obsRobotCla {
		s2 := s.next()
		s2.claRobots++
		s2.ore -= b.claRobotOre
		max = Max(max, b.evaluate(s2))
		//obsBuilt = true
	}

	if s.ore >= b.oreRobotOre && s.oreRobots < 4 {
		s2 := s.next()
		s2.oreRobots++
		s2.ore -= b.oreRobotOre
		max = Max(max, b.evaluate(s2))
		//obsBuilt = true
	}

	if !obsBuilt {
		max = Max(max, b.evaluate(s.next()))
	}

	return max
}

type state struct {
	minute             int16
	ore, cla, obs, geo int16
	oreRobots          int16
	claRobots          int16
	obsRobots          int16
	geoRobots          int16
}

func (s *state) next() state {
	s2 := *s
	s2.minute--
	s2.ore += s.oreRobots
	s2.cla += s.claRobots
	s2.obs += s.obsRobots
	s2.geo += s.geoRobots
	return s2
}

func solve(p *Problem) {
	blueprints := []*blueprint(nil)
	for p.NextLine() {
		b := blueprint{}
		p.Scanf("Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. "+
			"Each obsidian robot costs %d ore and %d clay. "+
			"Each geode robot costs %d ore and %d obsidian.",
			&b.id, &b.oreRobotOre, &b.claRobotOre, &b.obsRobotOre, &b.obsRobotCla, &b.geoRobotOre, &b.geoRobotObs)
		blueprints = append(blueprints, &b)
	}

	sum := 0
	for _, b := range blueprints {
		q := b.evaluate(state{minute: 24, oreRobots: 1})
		//p.Printf("%d: %d", b.id, q)
		sum += b.id * q
	}
	p.PartOne(sum)

	if len(blueprints) > 3 {
		blueprints = blueprints[:3]
	}

	prod := 1
	for _, b := range blueprints {
		q := b.evaluate(state{minute: 32, oreRobots: 1})
		//p.Printf("%d: %d", b.id, q)
		prod *= q
	}
	p.PartTwo(prod)
}
