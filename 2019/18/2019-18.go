package main

import (
	"fmt"
	. "github.com/bendiscz/aoc/aoc"
	"math"
	"math/bits"
	"slices"
)

const (
	year    = 2019
	day     = 18
	example = `

#############
#DcBa.#.GhKl#
#.###...#I###
#e#d#.@.#j#k#
###C#...###J#
#fEbA.#.FgHi#
#############

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
	v  byte
}

func (c cell) String() string { return string([]byte{c.ch}) }

type grid struct {
	*Matrix[cell]
}

func isStart(ch byte) bool { return ch == '@' }
func isKey(ch byte) bool   { return ch >= 'a' && ch <= 'z' }
func isDoor(ch byte) bool  { return ch >= 'A' && ch <= 'Z' }

type keybits uint

func keybit(ch byte) keybits {
	switch {
	case isStart(ch):
		return 1
	case isKey(ch):
		return 2 << (ch - 'a')
	case isDoor(ch):
		return 2 << (ch - 'A')
	default:
		panic("invalid key char")
	}
}

func (k keybits) unlocks(doors keybits) bool {
	return doors & ^k == 0
}

func (k keybits) count() int {
	return bits.OnesCount(uint(k))
}

type vertex struct {
	net   *net
	key   byte
	pos   XY
	edges []*edge
}

func (v *vertex) String() string {
	return fmt.Sprintf("%c %v -> %d", v.key, v.pos, len(v.edges))
}

type edge struct {
	key   byte
	len   int
	doors keybits
}

type net struct {
	v    map[byte]*vertex
	keys keybits
}

func makeNet(g Grid[cell]) *net {
	n := &net{v: map[byte]*vertex{}}
	for xy, c := range g.All() {
		if isKey(c.ch) || isStart(c.ch) {
			v := &vertex{net: n, key: c.ch, pos: xy}
			n.v[c.ch] = v
			n.keys |= keybit(v.key)
			findEdges(g, v)
		}
	}

	return n
}

func (n net) start() *vertex {
	return n.v['@']
}

func findEdges(g Grid[cell], v *vertex) {
	type qs struct {
		at    XY
		len   int
		doors keybits
	}

	q := Queue[qs]{}
	q.Push(qs{at: v.pos})
	for q.Len() > 0 {
		s := q.Pop()
		c := g.At(s.at)
		c.v = v.key

		if isKey(c.ch) && s.at != v.pos {
			v.edges = append(v.edges, &edge{
				key:   c.ch,
				len:   s.len,
				doors: s.doors,
			})
			continue
		}

		if isDoor(c.ch) {
			s.doors |= keybit(c.ch)
		}

		for _, d := range HVDirs {
			at := s.at.Add(d)
			if !g.Size().HasInside(at) {
				continue
			}
			c2 := g.At(at)
			if c2.v != v.key && c2.ch != '#' {
				q.Push(qs{at: at, len: s.len + 1, doors: s.doors})
			}
		}
	}
}

func solve(p *Problem) {
	g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	for p.NextLine() {
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{ch: b} })
	}

	// possible improvement:
	// - use tree instead of generic graph
	// - construct it by flood fill
	// - tree nodes are either keys or maze crossroads
	// - trim branches without keys
	// - every node knows which doors are in its subtree
	// - when visiting subtree that is completely unlocked, visit it entirely (do not return until all keys are collected)
	//
	// - but the current solution is good enough :-) (~150ms)

	solvePartOne(p, g)

	if p.Example() {
		//return
	}

	solvePartTwo(p, g)
}

func solvePartOne(p *Problem, g grid) {
	n := makeNet(g)
	s := n.start()
	t := traveler{
		n:      n,
		best:   path{len: math.MaxInt},
		limits: map[snapshot]int{},
	}
	t.findPath(path{steps: []*vertex{s}, keys: keybit(s.key)})
	p.PartOne(t.best.len)
}

func solvePartTwo(p *Problem, g grid) {
	o := XY0
	for xy, c := range g.All() {
		c.v = 0
		if isStart(c.ch) {
			o = xy
		}
	}

	g.At(o).ch = '#'
	for _, d := range HVDirs {
		g.At(o.Add(d)).ch = '#'
	}
	for _, d := range DiagDirs {
		g.At(o.Add(d)).ch = '@'
	}

	ns := [4]*net{}
	ns[0] = makeNet(g.SubGrid(XY0, o.Add(Square(1))))
	ns[1] = makeNet(g.SubGrid(XY{o.X, 0}, XY{g.Dim.X - o.X, o.Y + 1}))
	ns[2] = makeNet(g.SubGrid(XY{0, o.Y}, XY{o.X + 1, g.Dim.Y - o.Y}))
	ns[3] = makeNet(g.SubGrid(o, g.Dim.Sub(o)))

	t := traveler2{
		nets:    ns,
		best:    path2{len: math.MaxInt},
		allKeys: 0,
		limits:  map[snapshot2]int{},
	}
	for i := 0; i < 4; i++ {
		t.allKeys |= ns[i].keys
	}
	s0, s1, s2, s3 := ns[0].start(), ns[1].start(), ns[2].start(), ns[3].start()
	t.findPath(path2{v: [4]*vertex{s0, s1, s2, s3}, steps: []*vertex{s0, s1, s2, s3}, keys: keybit(s0.key)})
	p.PartTwo(t.best.len)
}

type path struct {
	steps []*vertex
	len   int
	keys  keybits
}

func (p path) step(s step) path {
	return path{
		steps: append(slices.Clone(p.steps), s.v),
		len:   p.len + s.len,
		keys:  p.keys | keybit(s.v.key),
	}
}

type snapshot struct {
	key  byte
	keys keybits
}

type traveler struct {
	n      *net
	best   path
	limits map[snapshot]int
}

func (t *traveler) findPath(p path) {
	if t.n.keys == p.keys {
		if p.len < t.best.len {
			t.best = p
		}
		return
	}

	v := p.steps[len(p.steps)-1]
	k := snapshot{key: v.key, keys: p.keys}
	if limit, ok := t.limits[k]; ok && p.len >= limit {
		return
	}
	t.limits[k] = p.len

	steps := findNextSteps(v, p.keys)
	for _, s := range steps {
		t.findPath(p.step(s))
	}
}

type step struct {
	v   *vertex
	len int
}

func findNextSteps(v *vertex, keys keybits) []step {
	type qs struct {
		v   *vertex
		len int
	}

	steps := []step(nil)
	q := Queue[qs]{}
	q.Push(qs{v: v})
	seen := keybit(v.key)

	for q.Len() > 0 {
		s := q.Pop()
		for _, e := range s.v.edges {
			next := v.net.v[e.key]
			if seen&keybit(next.key) != 0 || !keys.unlocks(e.doors) {
				continue
			}

			seen |= keybit(next.key)

			if keys&keybit(e.key) == 0 {
				steps = append(steps, step{v: next, len: s.len + e.len})
			} else {
				q.Push(qs{v: next, len: s.len + e.len})
			}
		}
	}

	return steps
}

type snapshot2 struct {
	at   [4]byte
	keys keybits
}

type traveler2 struct {
	nets    [4]*net
	best    path2
	allKeys keybits
	limits  map[snapshot2]int
}

type path2 struct {
	v     [4]*vertex
	steps []*vertex
	len   int
	keys  keybits
}

func (t *traveler2) step(p path2, s step) path2 {
	n := path2{
		v:     p.v,
		steps: append(slices.Clone(p.steps), s.v),
		len:   p.len + s.len,
		keys:  p.keys | keybit(s.v.key),
	}
	for i := 0; i < 4; i++ {
		if s.v.net == t.nets[i] {
			n.v[i] = s.v
		}
	}
	return n
}

func (t *traveler2) findPath(p path2) {
	if p.keys == t.allKeys {
		if p.len < t.best.len {
			t.best = p
		}
		return
	}

	k := snapshot2{at: [4]byte{p.v[0].key, p.v[1].key, p.v[2].key, p.v[3].key}, keys: p.keys}
	if limit, ok := t.limits[k]; ok && p.len >= limit {
		return
	}
	t.limits[k] = p.len

	for i := 0; i < 4; i++ {
		steps := findNextSteps(p.v[i], p.keys)
		for _, s := range steps {
			t.findPath(t.step(p, s))
		}
	}
}
