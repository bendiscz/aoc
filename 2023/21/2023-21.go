package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 21
	example = `

...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	plot    bool
	reached bool
	dist    int
}

type grid struct {
	*Matrix[cell]
}

type path struct {
	XY
	d int
}

func solve(p *Problem) {
	g, start := grid{}, XY{}
	for p.NextLine() {
		if g.Matrix == nil {
			g.Matrix = NewMatrix[cell](Rectangle(len(p.Line()), 0))
		}

		i := strings.IndexByte(p.Line(), 'S')
		if i >= 0 {
			start = XY{i, g.Dim.Y}
		}

		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{plot: b != '#'} })
	}

	dist := 64
	if p.Example() {
		dist = 6
	}

	g.fill(start)
	p.PartOne(g.count(dist))

	if p.Example() {
		return
	}

	dist = 26501365

	xm, ym := g.Dim.X-1, g.Dim.Y-1
	offset, step := g.Dim.X/2+1, g.Dim.X

	part2 := g.count(dist)                                         // origin
	part2 += g.compute(XY{start.X, 0}, offset, step, dist, false)  // down
	part2 += g.compute(XY{start.X, ym}, offset, step, dist, false) // up
	part2 += g.compute(XY{0, start.Y}, offset, step, dist, false)  // right
	part2 += g.compute(XY{xm, start.Y}, offset, step, dist, false) // left
	part2 += g.compute(XY{0, 0}, 2*offset, step, dist, true)       // right down
	part2 += g.compute(XY{xm, 0}, 2*offset, step, dist, true)      // left down
	part2 += g.compute(XY{0, ym}, 2*offset, step, dist, true)      // right up
	part2 += g.compute(XY{xm, ym}, 2*offset, step, dist, true)     // left up

	p.PartTwo(part2)
}

func (g grid) reset() {
	g.Dim.ForEach(func(xy XY) {
		c := g.At(xy)
		c.reached = false
		c.dist = 0
	})
}

func (g grid) fill(start XY) {
	q := Queue[path]{}
	q.Push(path{start, 0})

	for q.Len() > 0 {
		p := q.Pop()
		c := g.At(p.XY)
		if c.reached {
			continue
		}

		c.reached = true
		c.dist = p.d

		for _, dir := range HVDirs {
			xy := p.Add(dir)
			if g.Dim.HasInside(xy) && g.At(xy).plot {
				q.Push(path{xy, p.d + 1})
			}
		}
	}
}

func (g grid) count(dist int) int {
	count := 0
	g.Dim.ForEach(func(xy XY) {
		c := g.At(xy)
		if c.reached && c.dist <= dist && (dist-c.dist)%2 == 0 {
			count++
		}
	})
	return count
}

func (g grid) compute(s XY, offset, step, dist int, diag bool) int {
	g.reset()
	g.fill(s)

	h, countEvenOdd := 0, [2]int{}
	g.Dim.ForEach(func(xy XY) {
		c := g.At(xy)
		if c.reached {
			h = max(h, c.dist)
			countEvenOdd[c.dist%2]++
		}
	})

	count, steps := 0, 1
	for d0 := offset; d0 <= dist; d0 += step {
		tile := 0
		d := dist - d0
		if d > h {
			tile = countEvenOdd[d%2]
		} else {
			tile = g.count(d)
		}
		if diag {
			tile *= steps
		}
		count += tile
		steps++
	}
	return count
}
