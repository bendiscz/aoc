package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 18
	example = `

.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

func (c cell) String() string { return string([]byte{c.ch}) }

type grid struct {
	*Matrix[cell]
}

func (g grid) count(p XY) (int, int) {
	wood, yard := 0, 0
	for _, d := range AllDirs {
		p2 := p.Add(d)
		if !g.Dim.HasInside(p2) {
			continue
		}
		switch g.At(p2).ch {
		case '|':
			wood++
		case '#':
			yard++
		}
	}
	return wood, yard
}

func (g grid) resourceValue() int {
	wood, yard := 0, 0
	for _, c := range g.All() {
		switch c.ch {
		case '|':
			wood++
		case '#':
			yard++
		}
	}
	return wood * yard
}

func (g grid) step() grid {
	g2 := grid{NewMatrix[cell](g.Dim)}
	for p, c2 := range g2.All() {
		wood, yard := g.count(p)
		c := g.At(p)
		switch c.ch {
		case '.':
			if wood >= 3 {
				c2.ch = '|'
			} else {
				c2.ch = '.'
			}
		case '|':
			if yard >= 3 {
				c2.ch = '#'
			} else {
				c2.ch = '|'
			}
		case '#':
			if wood >= 1 && yard >= 1 {
				c2.ch = '#'
			} else {
				c2.ch = '.'
			}
		}
	}
	return g2
}

func (g grid) isEqual(g2 grid) bool {
	for xy, c := range g.All() {
		if g2.At(xy).ch != c.ch {
			return false
		}
	}
	return true
}

type snapshot struct {
	t int
	g grid
}

func solve(p *Problem) {
	g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	for p.NextLine() {
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{b} })
	}

	history := []int(nil)
	snapshots := map[int][]snapshot{}

	for t := 0; t < 100000; t++ {
		value := g.resourceValue()
		history = append(history, value)

		if t == 10 {
			p.PartOne(value)
		}

		for _, s := range snapshots[value] {
			if s.g.isEqual(g) {
				p.PartTwo(history[s.t+(1_000_000_000-t)%(t-s.t)])
				return
			}
		}
		snapshots[value] = append(snapshots[value], snapshot{t, g})

		g = g.step()
	}
}
