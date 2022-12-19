package main

import (
	"math"

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
	//Run(year, day, "", solve)
}

const types = 4

const (
	ore = iota
	clay
	obsidian
	geode
)

type cost [types]int16

type blueprint struct {
	id     int
	costs  [types]cost
	limits [types]int16
}

type state struct {
	minutes   int16
	materials [types]int16
	robots    [types]int16
}

func initState(minutes int) state {
	s := state{minutes: int16(minutes)}
	s.robots[ore] = 1
	return s
}

func (s state) build(b *blueprint, rt int, skip int16) state {
	for i := 0; i < types; i++ {
		s.materials[i] += (skip+1)*s.robots[i] - b.costs[rt][i]
	}
	s.minutes -= skip + 1
	s.robots[rt]++
	return s
}

func (b *blueprint) search(s state, max *int16) {
	built := false
	for rt := 0; rt < types; rt++ {
		if s.robots[rt] == b.limits[rt] {
			continue
		}

		var skip int16
		for i := 0; i < types-1 && skip <= s.minutes; i++ {
			missing := b.costs[rt][i] - s.materials[i]
			switch {
			case missing <= 0:
				break
			case s.robots[i] == 0:
				skip = math.MaxInt16
			default:
				skip = Max(skip, (missing-1)/s.robots[i]+1)
			}
		}

		if skip >= s.minutes {
			continue
		}

		next := s.build(b, rt, skip)

		// sum:
		// - current geodes
		// - product of already built geode robots for remaining time
		// - product of new geode robots if we build one for every remaining minute
		potential := next.materials[geode] + next.minutes*next.robots[geode] + (next.minutes-1)*next.minutes/2
		if potential <= *max {
			continue
		}

		built = true
		b.search(next, max)
	}

	if built {
		return
	}

	// no more robots can be built, compute geodes
	*max = Max(*max, s.materials[geode]+s.robots[geode]*s.minutes)
}

func (b *blueprint) simulate(minutes int) int {
	max := int16(0)
	b.search(initState(minutes), &max)
	return int(max)
}

func solve(p *Problem) {
	// optimization
	// - always go for some robot (eliminate "idle" states)
	// - try to build more "basic" robots first
	// - do not build more robots than needed
	// - compute upper bound for every state (what we got if a new geode robot is built for every remaining minute)
	//   and prune the state if it is not better than the current maximum
	blueprints := []*blueprint(nil)
	for p.NextLine() {
		b := blueprint{}
		p.Scanf("Blueprint %d: Each ore robot costs %d ore."+
			" Each clay robot costs %d ore."+
			" Each obsidian robot costs %d ore and %d clay."+
			" Each geode robot costs %d ore and %d obsidian.",
			&b.id, &b.costs[ore][ore],
			&b.costs[clay][ore],
			&b.costs[obsidian][ore], &b.costs[obsidian][clay],
			&b.costs[geode][ore], &b.costs[geode][obsidian])

		for i := 0; i < types; i++ {
			b.limits[ore] = Max(b.limits[ore], b.costs[i][ore])
		}
		b.limits[clay] = b.costs[obsidian][clay]
		b.limits[obsidian] = b.costs[geode][obsidian]
		b.limits[geode] = math.MaxInt16

		blueprints = append(blueprints, &b)
	}

	sum := 0
	for _, b := range blueprints {
		q := b.simulate(24)
		//p.Printf("%d: %d", b.id, q)
		sum += b.id * q
	}
	p.PartOne(sum)

	if len(blueprints) > 3 {
		blueprints = blueprints[:3]
	}

	sum = 1
	for _, b := range blueprints {
		q := b.simulate(32)
		//p.Printf("%d: %d", b.id, q)
		sum *= q
	}
	p.PartTwo(sum)
}
