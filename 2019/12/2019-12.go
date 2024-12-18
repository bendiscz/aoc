package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 12
	example = `

<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>

`
)

//<x=-8, y=-10, z=0>
//<x=5, y=5, z=10>
//<x=2, y=-7, z=3>
//<x=9, y=-8, z=-3>

func main() {
	Run(year, day, example, solve)
}

type xyz [3]int

type moon struct {
	p, v xyz
}

func solve(p *Problem) {
	s1 := 0

	init := []xyz(nil)
	moons := []*moon(nil)
	for i := 0; p.NextLine(); i++ {
		f := ParseInts(p.Line())
		m := moon{p: [...]int{f[0], f[1], f[2]}}
		init = append(init, m.p)
		moons = append(moons, &m)
	}

	step, periods := 0, [3]int{}
	for {
		for _, m1 := range moons {
			for _, m2 := range moons {
				if m1 == m2 {
					continue
				}
				for k := 0; k < 3; k++ {
					m1.v[k] += Sign(m2.p[k] - m1.p[k])
				}
			}
		}

		for _, m := range moons {
			for k := 0; k < 3; k++ {
				m.p[k] += m.v[k]
			}
		}

		step++

		done := true
	loop:
		for k := 0; k < 3; k++ {
			if periods[k] > 0 {
				continue
			}
			done = false
			for i, m := range moons {
				if m.v[k] != 0 || m.p[k] != init[i][k] {
					continue loop
				}
			}
			periods[k] = step
		}

		if step == 1000 {
			for _, m := range moons {
				ep := Abs(m.p[0]) + Abs(m.p[1]) + Abs(m.p[2])
				ek := Abs(m.v[0]) + Abs(m.v[1]) + Abs(m.v[2])
				s1 += ep * ek
			}

			p.PartOne(s1)
		}

		if step >= 1000 && done {
			break
		}
	}
	p.PartTwo(LCM(LCM(periods[0], periods[1]), periods[2]))
}
