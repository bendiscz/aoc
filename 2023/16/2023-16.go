package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 16
	example = `

.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....

`
)

func main() {
	Run(year, day, example, solve)
}

type dir int

const (
	right dir = iota
	left
	down
	up
)

var dirs = [...]XY{
	right: PosX,
	left:  NegX,
	down:  PosY,
	up:    NegY,
}

type cell struct {
	value  byte
	energy [4]bool
}

type grid struct {
	*Matrix[cell]
}

func (g grid) clear() {
	g.Dim.ForEach(func(xy XY) {
		g.At(xy).energy = [4]bool{}
	})
}

type ray struct {
	o XY
	d dir
}

func solve(p *Problem) {
	g := grid{}
	for p.NextLine() {
		if g.Matrix == nil {
			g.Matrix = NewMatrix[cell](Rectangle(len(p.Line()), 0))
		}
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{value: b} })
	}

	p.PartOne(trace(g, ray{XY{-1, 0}, right}))

	best := math.MinInt
	for x := 0; x < g.Dim.X; x++ {
		g.clear()
		best = max(best, trace(g, ray{XY{x, -1}, down}))
		g.clear()
		best = max(best, trace(g, ray{XY{x, g.Dim.Y}, up}))
	}
	for y := 0; y < g.Dim.Y; y++ {
		g.clear()
		best = max(best, trace(g, ray{XY{-1, y}, right}))
		g.clear()
		best = max(best, trace(g, ray{XY{g.Dim.X, y}, left}))
	}
	p.PartTwo(best)
}

func trace(g grid, r ray) int {
	sum := 0
	splits := []ray(nil)
	pos, d := r.o, r.d
	for {
		pos = pos.Add(dirs[d])
		if !g.Dim.HasInside(pos) {
			break
		}

		c := g.At(pos)
		if c.energy[d] {
			break
		}

		energized := true
		for _, e := range c.energy {
			if e {
				energized = false
			}
		}
		if energized {
			sum++
		}

		c.energy[d] = true

		switch c.value {
		case '/':
			d = 3 - d

		case '\\':
			d = (d + 2) % 4

		case '-':
			if d == down || d == up {
				d = left
				splits = append(splits, ray{pos, right})
			}

		case '|':
			if d == right || d == left {
				d = up
				splits = append(splits, ray{pos, down})
			}
		}
	}

	for _, r2 := range splits {
		sum += trace(g, r2)
	}

	return sum
}
