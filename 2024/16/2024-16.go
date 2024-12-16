package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"math"
	"strings"
)

const (
	year    = 2024
	day     = 16
	example = `

###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############

`
)

type cell struct {
	ch byte
	v  bool
	s  [4]*state
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

func (g grid) state(r reindeer) *state {
	c := g.At(r.p)
	s := c.s[r.d]
	if s == nil {
		s = &state{w: math.MaxInt}
		c.s[r.d] = s
	}
	return s
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

type state struct {
	w    int
	back []reindeer
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

	p.PartOne(findPath(g, start, end))

	for d := 0; d < 4; d++ {
		backtrack(g, reindeer{end, d})
	}
	//PrintGrid(g)

	s2 := 0
	for _, c := range g.All() {
		if c.v {
			s2++
		}
	}
	p.PartTwo(s2)
}

func findPath(g grid, start, end XY) int {
	r := reindeer{p: start}
	g.state(r).w = 0

	q := NewHeap[reindeer](func(r1, r2 reindeer) bool {
		return g.state(r1).w < g.state(r2).w
	})
	q.Push(r)

	best := math.MaxInt
	for q.Len() > 0 {
		r = q.Pop()
		s, fix := g.state(r), false

		for _, n := range r.nextSteps() {
			if g.At(n.r.p).ch == '#' {
				continue
			}
			w, ns := s.w+n.w, g.state(n.r)
			if w > ns.w || w > best {
				continue
			}
			if w == ns.w {
				ns.back = append(ns.back, r)
				continue
			}

			if ns.w < math.MaxInt {
				fix = true
			}
			ns.w = w
			ns.back = []reindeer{r}

			if n.r.p == end {
				best = min(best, w)
			} else {
				q.Push(n.r)
			}
		}
		if fix {
			q.Fix()
		}
	}
	return best
}

func backtrack(g grid, r reindeer) {
	g.At(r.p).v = true
	for _, b := range g.state(r).back {
		backtrack(g, b)
	}
}
