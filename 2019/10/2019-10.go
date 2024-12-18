package main

import (
	"cmp"
	. "github.com/bendiscz/aoc/aoc"
	"math"
	"slices"
)

const (
	year    = 2019
	day     = 10
	example = `

.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
	v  int
}

type grid struct {
	*Matrix[cell]
	gen int
}

type ray struct {
	p, d XY
	inf  bool
}

func angle(d XY) float64 {
	return math.Atan2(-float64(d.X), float64(d.Y))
}

func solve(p *Problem) {
	g := &grid{Matrix: NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	for p.NextLine() {
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{ch: b} })
	}

	s1, rays := 0, []ray(nil)
	for xy := range g.Dim.All() {
		s, r := g.count(xy)
		if s > s1 {
			s1 = s
			rays = r
		}
	}
	p.PartOne(s1)

	slices.SortFunc(rays, func(a, b ray) int {
		return cmp.Compare(angle(a.d), angle(b.d))
	})

	r := -1
	for i := 0; i < 200; i++ {
	loop:
		for {
			r = (r + 1) % len(rays)
			if rays[r].inf {
				continue
			}

			for {
				rays[r].p = rays[r].p.Add(rays[r].d)
				if !g.Dim.HasInside(rays[r].p) {
					rays[r].inf = true
					continue loop
				}
				if g.At(rays[r].p).ch == '#' {
					break loop
				}
			}
		}
	}
	p.PartTwo(rays[r].p.X*100 + rays[r].p.Y)
}

func (g *grid) count(xy XY) (int, []ray) {
	if g.At(xy).ch == '.' {
		return 0, nil
	}

	g.gen++
	rays, s := []ray(nil), 0
	for p, c := range g.All() {
		if p == xy || c.ch == '.' || c.v == g.gen {
			continue
		}
		s++
		d := p.Sub(xy)
		n := GCD(Abs(d.X), Abs(d.Y))
		d.X /= n
		d.Y /= n
		rays = append(rays, ray{p: xy, d: d})
		for q := xy.Add(d); g.Dim.HasInside(q); q = q.Add(d) {
			g.At(q).v = g.gen
		}
	}
	return s, rays
}
