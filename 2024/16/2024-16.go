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

func (g grid) state(k key) *state {
	c := g.At(k.p)
	s := c.s[k.d]
	if s == nil {
		s = &state{w: math.MaxInt}
		c.s[k.d] = s
	}
	return s
}

var dirs = [...]XY{PosX, PosY, NegX, NegY}

type key struct {
	p XY
	d int
}

func (k key) left() key    { return key{k.p, (k.d + 3) % 4} }
func (k key) right() key   { return key{k.p, (k.d + 1) % 4} }
func (k key) forward() key { return key{k.p.Add(dirs[k.d]), k.d} }

type state struct {
	w   int
	inc []key
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

	for d := 1; d <= 3; d++ {
		backtrack(g, key{end, d})
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
	k := key{p: start}
	g.state(k).w = 0

	q := NewHeap[key](func(k1, k2 key) bool {
		return g.state(k1).w < g.state(k2).w
	})
	q.Push(k)

	best, fix := math.MaxInt, false
	checkState := func(k, next key, w int) {
		if g.At(next.p).ch == '#' {
			return
		}

		s := g.state(next)
		if w > s.w || w > best {
			return
		}
		if w == s.w {
			s.inc = append(s.inc, k)
			return
		}

		fix = s.w < math.MaxInt
		s.w = w
		s.inc = []key{k}
		if next.p == end {
			best = min(best, w)
		} else {
			q.Push(next)
		}
	}

	for q.Len() > 0 {
		if fix {
			q.Fix()
			fix = false
		}

		k = q.Pop()
		if k.p == end {
			continue
		}

		s := g.state(k)
		checkState(k, k.forward(), s.w+1)
		checkState(k, k.left(), s.w+1000)
		checkState(k, k.right(), s.w+1000)
	}

	return best
}

func backtrack(g grid, k key) {
	g.At(k.p).v = true
	for _, prev := range g.state(k).inc {
		backtrack(g, prev)
	}
}
