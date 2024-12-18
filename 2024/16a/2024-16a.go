package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"math"
	"slices"
	"strings"
)

const (
	year    = 2024
	day     = 16
	example = `

#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################

`
)

//###############
//#.......#....E#
//#.#.###.#.###.#
//#.....#.#...#.#
//#.###.#####.#.#
//#.#.#.......#.#
//#.#.#####.###.#
//#...........#.#
//###.#.#####.#.#
//#...#.....#.#.#
//#.#.#.###.#.#.#
//#.....#...#.#.#
//#.###.#.#.#.#.#
//#S..#.....#...#
//###############

//####
//##E#
//#..#
//#..#
//#S##
//####

type cell struct {
	ch byte
	v  bool
}

func (c cell) String() string {
	switch {
	case c.v:
		return "O"
	default:
		return string([]byte{c.ch})
	}
}

type grid struct {
	*Matrix[cell]
}

var dirs = [...]XY{PosX, PosY, NegX, NegY}

type reindeer struct {
	p XY
	d int
}

type step struct {
	r reindeer
	w int
}

func (r reindeer) nextSteps() []step {
	dl, dr := (r.d+3)%4, (r.d+1)%4
	return []step{
		{reindeer{r.p.Add(dirs[r.d]), r.d}, 1},
		{reindeer{r.p.Add(dirs[dl]), dl}, 1001},
		{reindeer{r.p.Add(dirs[dr]), dr}, 1001},
	}
}

type vertex struct {
	r reindeer
	w int
}

type path struct {
	r reindeer
	w int
	p []reindeer
}

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	start, end := XY0, XY0
	for p.NextLine() {
		if x := strings.IndexByte(p.Line(), 'S'); x >= 0 {
			start = XY{X: x, Y: g.Dim.Y}
		}
		if x := strings.IndexByte(p.Line(), 'E'); x >= 0 {
			end = XY{X: x, Y: g.Dim.Y}
		}
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell {
			return cell{ch: b}
		})
	}

	s1 := dijkstra(g, start, end)
	p.PartOne(s1)

	PrintGrid(g)

	s2 := 0
	for _, c := range g.All() {
		if c.v {
			s2++
		}
	}
	p.PartTwo(s2)
}

func dijkstra(g grid, start, end XY) int {
	vs := map[reindeer]*vertex{}
	r0 := reindeer{start, 0}
	vs[r0] = &vertex{r: r0}
	work := []path{{r: r0}}

	found, best := false, 0

	for len(work) > 0 {
		var p path
		work, p = extractBest(work)

		if found && p.w > best {
			break
		}

		if p.r.p == end {
			g.At(end).v = true
			for _, r := range p.p {
				g.At(r.p).v = true
			}

			found = true
			best = p.w
		}

		for _, s := range p.r.nextSteps() {
			if g.At(s.r.p).ch == '#' {
				continue
			}

			w := p.w + s.w
			pn := path{r: s.r, w: w, p: append(slices.Clone(p.p), p.r)}
			vn := vs[pn.r]

			if vn == nil {
				vn = &vertex{pn.r, w}
				vs[pn.r] = vn
				work = append(work, pn)
			} else {
				if w <= vn.w {
					vn.w = w
					work = append(work, pn)
				}
			}
		}
	}

	return best
}

func extractBest(work []path) ([]path, path) {
	w, i0 := math.MaxInt, 0
	for i, p := range work {
		if p.w < w {
			w, i0 = p.w, i
		}
	}
	p := work[i0]
	work = slices.Delete(work, i0, i0+1)
	return work, p
}
