package main

import (
	"sort"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 11
	example = `

...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....

`
)

func main() {
	Run(year, day, example, solve)
}

type galaxy struct {
	XY
}

func (g *galaxy) dist(g2 *galaxy) int {
	d := g.Sub(g2.XY)
	return Abs(d.X) + Abs(d.Y)
}

type line struct {
	offset int
	gs     []*galaxy
}

func solve(p *Problem) {
	galaxies := []*galaxy(nil)
	for y := 0; p.NextLine(); y++ {
		for x, ch := range p.LineBytes() {
			if ch == '#' {
				galaxies = append(galaxies, &galaxy{XY{x, y}})
			}

		}
	}

	hm, vm := map[int][]*galaxy{}, map[int][]*galaxy{}
	for _, g := range galaxies {
		hm[g.X] = append(hm[g.X], g)
		vm[g.Y] = append(vm[g.Y], g)
	}

	hl := makeLines(hm)
	vl := makeLines(vm)

	expand(hl, 2, true)
	expand(vl, 2, false)
	p.PartOne(sumDists(galaxies))

	expand(hl, 500_000, true)
	expand(vl, 500_000, false)
	p.PartTwo(sumDists(galaxies))
}

func makeLines(m map[int][]*galaxy) []*line {
	ls := []*line(nil)
	for offset, gs := range m {
		ls = append(ls, &line{offset, gs})
	}
	sort.Slice(ls, func(i, j int) bool {
		return ls[i].offset < ls[j].offset
	})
	return ls
}

func expand(ls []*line, f int, h bool) {
	x, d := 0, 0
	for _, l := range ls {
		d += (f - 1) * (l.offset - x - 1)
		x = l.offset

		l.offset += d
		for _, g := range l.gs {
			if h {
				g.X += d
			} else {
				g.Y += d
			}
		}
	}
}

func sumDists(gs []*galaxy) int {
	sum := 0
	for i := 0; i < len(gs); i++ {
		for j := i + 1; j < len(gs); j++ {
			sum += gs[i].dist(gs[j])
		}
	}
	return sum
}
