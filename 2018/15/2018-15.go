package main

import (
	"cmp"
	"fmt"
	. "github.com/bendiscz/aoc/aoc"
	"slices"
	"sort"
)

const (
	year    = 2018
	day     = 15
	example = `

#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######

`
)

//#######
//#E..EG#
//#.#G.E#
//#E.##E#
//#G..#.#
//#..E#.#
//#######

func main() {
	Run(year, day, example, solve)
}

var dirs = [...]XY{NegY, NegX, PosX, PosY}

func cmpXY(a, b XY) int {
	r := cmp.Compare(a.Y, b.Y)
	if r == 0 {
		r = cmp.Compare(a.X, b.X)
	}
	return r
}

type cell struct {
	wall bool
	mob  *mob
}

func (c cell) String() string {
	switch {
	case c.mob != nil:
		return string([]byte{c.mob.kind})
	case c.wall:
		return "#"
	default:
		return "."
	}
}

func (c cell) free() bool {
	return !c.wall && c.mob == nil
}

type mob struct {
	g *grid

	kind byte
	pos  XY
	hp   int
	ap   int
}

func (m *mob) String() string {
	return fmt.Sprintf("%c%v %d", m.kind, m.pos, m.hp)
}

func (m *mob) clone(g *grid) *mob {
	m2 := *m
	m2.g = g
	return &m2
}

func (m *mob) dead() bool {
	return m.hp == 0
}

func (m *mob) hit(ap int) {
	m.hp = max(0, m.hp-ap)
}

func (m *mob) moveTo(pos XY) {
	if pos == m.pos {
		return
	}

	dst := m.g.At(pos)
	if !dst.free() {
		panic(fmt.Sprintf("cannot move to %v", pos))
	}

	m.g.At(m.pos).mob = nil
	m.pos = pos
	dst.mob = m
}

func (m *mob) hasTarget() bool {
	for _, n := range m.g.mobs {
		if n != nil && n.kind != m.kind {
			return true
		}
	}
	return false
}

func (m *mob) findNextStep() (XY, bool) {
	visited := map[XY]XY{}
	visited[m.pos] = m.pos

	q := Queue[XY]{}
	q.Push(m.pos)

	for q.Len() > 0 {
		xy := q.Pop()
		for _, d := range dirs {
			next := xy.Add(d)
			if _, ok := visited[next]; ok {
				continue
			}
			visited[next] = xy

			c := m.g.At(next)
			switch {
			case c.mob != nil && c.mob.kind != m.kind:
				for prev := xy; prev != m.pos; prev = visited[prev] {
					xy = prev
				}
				return xy, xy != m.pos

			case c.free():
				q.Push(next)
			}
		}
	}
	return XY0, false
}

func (m *mob) fight() *mob {
	var target *mob
	for _, d := range dirs {
		n := m.g.At(m.pos.Add(d)).mob
		if n != nil && n.kind != m.kind && !n.dead() && (target == nil || n.hp < target.hp) {
			target = n
		}
	}

	if target != nil {
		target.hit(m.ap)
	}

	return target
}

type grid struct {
	*Matrix[cell]
	round   int
	mobs    []*mob
	corpses []*mob
}

func (g *grid) clone() *grid {
	g2 := &grid{
		Matrix: NewMatrix[cell](g.Dim),
		round:  g.round,
	}
	for xy, c := range g.All() {
		c2 := g2.At(xy)
		c2.wall = c.wall
	}
	for _, m := range g.mobs {
		m2 := m.clone(g2)
		g2.mobs = append(g2.mobs, m2)
		g2.At(m2.pos).mob = m2
	}
	for _, m := range g.corpses {
		m2 := m.clone(g2)
		g2.corpses = append(g2.corpses, m2)
	}
	return g2
}

func (g *grid) parseRow(line string) {
	y := g.Dim.Y
	row := g.AppendRow()
	for x := 0; x < g.Dim.X; x++ {
		switch {
		case line[x] == '#':
			row.At(x).wall = true
		case line[x] >= 'A' && line[x] <= 'Z':
			m := &mob{
				g:    g,
				kind: line[x],
				pos:  XY{X: x, Y: y},
				hp:   200,
				ap:   3,
			}
			g.mobs = append(g.mobs, m)
			row.At(x).mob = m
		}
	}
}

func (g *grid) playRound() bool {
	slices.SortFunc(g.mobs, func(a, b *mob) int { return cmpXY(a.pos, b.pos) })
	defer func() {
		g.mobs = slices.DeleteFunc(g.mobs, func(m *mob) bool { return m == nil })
	}()
	for _, m := range g.mobs {
		if m == nil {
			continue
		}
		if !m.hasTarget() {
			return false
		}
		if xy, ok := m.findNextStep(); ok {
			m.moveTo(xy)
		}
		t := m.fight()
		if t != nil && t.dead() {
			j := slices.Index(g.mobs, t)
			g.mobs[j] = nil
			g.corpses = append(g.corpses, t)
			g.At(t.pos).mob = nil
		}
	}
	return true
}

func (g *grid) playPartOne() {
	for g.playRound() {
		g.round++
	}
}

func (g *grid) playPartTwo(ap int) bool {
	g.boost('E', ap)
	for g.playRound() {
		for _, m := range g.corpses {
			if m.kind == 'E' {
				return false
			}
		}
		g.round++
	}
	return true
}

func (g *grid) outcome() int {
	hp := 0
	for _, m := range g.mobs {
		hp += m.hp
	}
	return hp * g.round
}

func (g *grid) boost(kind byte, ap int) {
	for _, m := range g.mobs {
		if m.kind == kind {
			m.ap = ap
		}
	}
}

func solve(p *Problem) {
	g0 := &grid{Matrix: NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	for p.NextLine() {
		g0.parseRow(p.Line())
	}

	g := g0.clone()
	g.playPartOne()
	p.PartOne(g.outcome())

	ap := 4
	for !g0.clone().playPartTwo(ap) {
		ap *= 2
	}
	ap = sort.Search(ap, func(ap int) bool { return g0.clone().playPartTwo(ap) })

	g = g0.clone()
	g.playPartTwo(ap)
	p.PartTwo(g.outcome())
}
