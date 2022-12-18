package aoc

import (
	"reflect"
	"testing"
)

type cell struct {
	c bool
	w int
}

type matrix struct {
	d int
	w [][]cell
}

func newMatrix(d int) *matrix {
	m := matrix{
		d: d,
		w: make([][]cell, d),
	}

	for i := range m.w {
		m.w[i] = make([]cell, d)
	}

	return &m
}

func (m *matrix) vertex(i int) matrixVertex {
	return matrixVertex{m, i}
}

func (m *matrix) connect(i, j, w int) {
	m.w[i][j] = cell{true, w}
}

func (m *matrix) connectBoth(i, j, w int) {
	m.connect(i, j, w)
	m.connect(j, i, w)
}

type matrixVertex struct {
	m *matrix
	i int
}

func (v matrixVertex) Key() any {
	return v.i
}

func (v matrixVertex) Edges() []Edge {
	var edges []Edge
	for j := 0; j < v.m.d; j++ {
		if c := v.m.w[v.i][j]; c.c {
			edges = append(edges, Edge{V: v.m.vertex(j), W: c.w})
		}
	}
	return edges
}

func TestDijkstra(t *testing.T) {
	m := newMatrix(9)
	m.connectBoth(0, 1, 4)
	m.connectBoth(0, 7, 8)

	m.connectBoth(1, 2, 8)
	m.connectBoth(1, 7, 11)

	m.connectBoth(2, 3, 7)
	m.connectBoth(2, 5, 4)
	m.connectBoth(2, 8, 2)

	m.connectBoth(3, 4, 9)
	m.connectBoth(3, 5, 14)

	m.connectBoth(4, 5, 10)

	m.connectBoth(5, 6, 2)

	m.connectBoth(6, 7, 1)
	m.connectBoth(6, 8, 6)

	m.connectBoth(7, 8, 7)

	m.connect(0, 0, 0)
	m.connect(2, 2, 0)
	m.connect(4, 4, 0)
	m.connect(6, 6, 0)

	path := ShortestPath(m.vertex(0), m.vertex(4))
	if want, got := 21, path.Cost; got != want {
		t.Errorf("Cost want: %v, got: %v", want, got)
	}

	steps := func(p Path[matrixVertex]) []int {
		var steps []int
		for _, s := range p.Steps {
			steps = append(steps, s.i)
		}
		return steps
	}

	if want, got := []int{0, 7, 6, 5, 4}, steps(path); !reflect.DeepEqual(got, want) {
		t.Errorf("Path want: %+v, got: %+v", want, got)
	}

	paths := ShortestPaths(m.vertex(0), func(v matrixVertex) (bool, bool) { return v.i != 0, false })
	wantPaths := []struct {
		cost  int
		steps []int
	}{
		{4, []int{0, 1}},
		{8, []int{0, 1, 2, 8}},
		{9, []int{0, 7, 6}},
		{11, []int{0, 7, 6, 5}},
		{12, []int{0, 1, 2}},
		{14, []int{0, 1, 2, 8}},
		{19, []int{0, 1, 2, 3}},
		{21, []int{0, 7, 6, 5, 4}},
	}

	if want, got := len(wantPaths), len(paths); got != want {
		t.Errorf("len(paths) want: %d, got: %d", want, got)
	}
	for i := range paths {
		if want, got := wantPaths[i].cost, paths[i].Cost; got != want {
			t.Errorf("path #%d: Cost want: %d, got: %d", i, want, got)
		}
	}
}

func TestDijkstraWithNegativeCycle(t *testing.T) {
	m := newMatrix(4)
	m.connect(0, 1, 1)
	m.connect(1, 2, 1)
	m.connect(2, 1, -2)
	m.connect(2, 3, 1)

	assertPanic(t, "edge from 2 to 1 has negative weight: -2", func() {
		ShortestPath(m.vertex(0), m.vertex(3))
	})
}
