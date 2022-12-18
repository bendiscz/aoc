package aoc

type Vertex interface {
	Key() any
	Edges() []Edge
}

type Edge struct {
	V Vertex
	W int
}
