package main

import (
	"cmp"
	. "github.com/bendiscz/aoc/aoc"
	"slices"
)

const (
	year    = 2018
	day     = 20
	example = `

^WNE$
^ENWWW(NEEE|SSE(EE|N))$
^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$
^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$
^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$

`
)

func main() {
	Run(year, day, example, solve)
}

var (
	compass  = [...]int{'N': 0, 'S': 1, 'E': 2, 'W': 3}
	opposite = [...]int{0: 1, 1: 0, 2: 3, 3: 2}
	dir      = [...]XY{NegY, PosY, PosX, NegX}
)

func cmpXY(a, b XY) int {
	r := cmp.Compare(a.Y, b.Y)
	if r == 0 {
		r = cmp.Compare(a.X, b.X)
	}
	return r
}

type room struct {
	doors [4]bool
}

type grid map[XY]*room

func (g grid) room(xy XY) *room {
	r := g[xy]
	if r == nil {
		r = &room{}
		g[xy] = r
	}
	return r
}

func (g grid) pass(xy XY, d int) XY {
	g.room(xy).doors[d] = true
	xy = xy.Add(dir[d])
	g.room(xy).doors[opposite[d]] = true
	return xy
}

func (g grid) explore(regex string, index int, origin []XY) (int, []XY) {
	reached := []XY(nil)
	at := slices.Clone(origin)

loop:
	for index < len(regex) {
		ch := regex[index]
		switch ch {
		case ')':
			reached = append(reached, at...)
			break loop

		case '(':
			index, at = g.explore(regex, index+1, at)

		case '|':
			reached = append(reached, at...)
			at = slices.Clone(origin)

		default:
			for i := range at {
				at[i] = g.pass(at[i], compass[ch])
			}
		}
		index++
	}

	slices.SortFunc(reached, cmpXY)
	reached = slices.Compact(reached)

	return index, reached
}

func (g grid) findPaths(limit int) (int, int) {
	v, lMax, lLim := map[XY]int{}, 0, 0
	v[XY0] = 0
	q := Queue[XY]{}
	q.Push(XY0)
	for q.Len() > 0 {
		p := q.Pop()
		lMax = max(lMax, v[p])
		if v[p] >= limit {
			lLim++
		}

		r := g[p]
		for d := 0; d < 4; d++ {
			if !r.doors[d] {
				continue
			}
			p2 := p.Add(dir[d])
			if r2 := g[p2]; r2 != nil && v[p2] == 0 {
				v[p2] = v[p] + 1
				q.Push(p2)
			}
		}
	}
	return lMax, lLim
}

func solve(p *Problem) {
	for p.NextLine() {
		regex := p.Line()
		regex = regex[1 : len(regex)-1]

		g := grid{}
		g.room(XY0)
		g.explore(regex, 0, []XY{XY0})
		s1, s2 := g.findPaths(1000)
		p.PartOne(s1)
		p.PartTwo(s2)
	}
}
