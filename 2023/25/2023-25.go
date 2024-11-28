package main

import (
	"container/heap"
	"fmt"
	"slices"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 25
	example = `

jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr

`
)

// https://dl.acm.org/doi/pdf/10.1145/263867.263872

func main() {
	Run(year, day, example, solve)
}

type vertex struct {
	id    string
	edges []*edge

	count int

	conn  int
	index int
}

func (v *vertex) String() string { return v.id }

func (v *vertex) connect(v2 *vertex, w int) {
	for _, e := range v.edges {
		if e.v == v2 {
			e.w += w
			return
		}
	}
	v.edges = append(v.edges, &edge{v: v2, w: w})
}

type edge struct {
	v *vertex
	w int
}

func (e *edge) String() string { return fmt.Sprintf("%s %d", e.v.id, e.w) }

type graph struct {
	vertices map[string]*vertex
}

func newGraph() *graph {
	return &graph{
		vertices: map[string]*vertex{},
	}
}

func (g *graph) vertex(id string) *vertex {
	if v, ok := g.vertices[id]; ok {
		return v
	}
	v := &vertex{id: id, count: 1}
	g.vertices[id] = v
	return v
}

func (g *graph) merge(v1, v2 *vertex) {
	for _, e := range v2.edges {
		if e.v != v1 {
			v1.connect(e.v, e.w)
			e.v.connect(v1, e.w)
		}
		e.v.edges = slices.DeleteFunc(e.v.edges, func(e *edge) bool { return e.v == v2 })
	}
	v1.count += v2.count
	delete(g.vertices, v2.id)
}

func solve(p *Problem) {
	g := newGraph()
	for p.NextLine() {
		f := SplitFields(p.Line())
		v1 := g.vertex(f[0][:3])
		for _, id := range f[1:] {
			v2 := g.vertex(id)
			v1.connect(v2, 1)
			v2.connect(v1, 1)
		}
	}

	n := len(g.vertices)
	for {
		c, l := minCut(g)
		if c == 3 {
			p.PartOne((n - l) * l)
			break
		}
	}

	p.PartTwo("Merry Christmas!")
}

type remaining struct {
	v []*vertex
}

func (r *remaining) Len() int {
	return len(r.v)
}

func (r *remaining) Less(i, j int) bool {
	return r.v[i].conn > r.v[j].conn
}

func (r *remaining) Swap(i, j int) {
	r.v[i], r.v[j] = r.v[j], r.v[i]
	r.v[i].index = i
	r.v[j].index = j
}

func (r *remaining) Push(e any) {
	v := e.(*vertex)
	v.index = len(r.v)
	r.v = append(r.v, e.(*vertex))
}

func (r *remaining) Pop() any {
	v := r.v[len(r.v)-1]
	r.v = r.v[:len(r.v)-1]
	return v
}

func minCut(g *graph) (int, int) {
	a := map[string]*vertex{}
	r := &remaining{make([]*vertex, len(g.vertices)-1)}
	last := (*vertex)(nil)

	i := 0
	for id, v := range g.vertices {
		if len(a) == 0 {
			a[id] = v
			last = v
		} else {
			v.conn = 0
			v.index = i
			r.v[i] = v
			i++
		}
	}
	for _, e := range last.edges {
		e.v.conn = e.w
	}
	heap.Init(r)

	for len(r.v) > 1 {
		v := heap.Pop(r).(*vertex)
		for _, e := range v.edges {
			if _, ok := a[e.v.id]; !ok {
				e.v.conn += e.w
				heap.Fix(r, e.v.index)
			}
		}

		a[v.id] = v
		last = v
	}

	w0 := 0
	for _, e := range r.v[0].edges {
		w0 += e.w
	}

	g.merge(last, r.v[0])

	return w0, r.v[0].count
}
