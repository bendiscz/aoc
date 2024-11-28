package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 23
	example = `

#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	g := grid{}
	for p.NextLine() {
		if g.Matrix == nil {
			g.Matrix = NewMatrix[cell](Rectangle(len(p.Line()), 0))
		}
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{ch: b} })
	}

	start, end := XY{1, 0}, XY{g.Dim.X - 2, g.Dim.Y - 1}

	r1, _ := search1(g, start.Add(PosY), end, 1, map[XY]bool{start: true})
	p.PartOne(r1)

	n := g.makeNet(start, end)
	r2, _ := search2(n.start, n.end, 0, map[XY]bool{})
	p.PartTwo(r2)
}

func search1(g grid, xy, end XY, length int, visited map[XY]bool) (int, bool) {
	if xy == end {
		return length, true
	}

	visited[xy] = true
	defer delete(visited, xy)

	var dirs []XY
	switch g.At(xy).ch {
	case '>':
		dirs = []XY{PosX}
	case '<':
		dirs = []XY{NegX}
	case 'v':
		dirs = []XY{PosY}
	case '^':
		dirs = []XY{NegY}
	default:
		dirs = []XY{PosX, NegX, PosY, NegY}
	}

	best, found := 0, false
	for _, d := range dirs {
		xy2 := xy.Add(d)
		if !visited[xy2] && g.At(xy2).ch != '#' {
			l, ok := search1(g, xy2, end, length+1, visited)
			if ok {
				best = max(best, l)
				found = true
			}
		}
	}

	return best, found
}

type vertex struct {
	xy    XY
	edges []edge
}

type edge struct {
	v *vertex
	d int
}

type net struct {
	vertices   map[XY]*vertex
	start, end *vertex
}

func (n *net) vertex(xy XY) *vertex {
	if v, ok := n.vertices[xy]; ok {
		return v
	}
	v := &vertex{xy: xy}
	n.vertices[xy] = v
	return v
}

func (v *vertex) connect(v2 *vertex, length int) {
	for _, e := range v.edges {
		if e.v == v2 {
			e.d = max(e.d, length)
			return
		}
	}
	v.edges = append(v.edges, edge{v2, length})
}

func (g grid) dirs(xy, dir XY) []XY {
	dirs := make([]XY, 0, 4)
	for _, d := range HVDirs {
		if dir.Add(d) != (XY{0, 0}) && g.Dim.HasInside(xy.Add(d)) && g.At(xy.Add(d)).ch != '#' {
			dirs = append(dirs, d)
		}
	}
	return dirs
}

func (g grid) travel(xy, dir XY) (XY, XY, int) {
	for length := 1; ; length++ {
		xy = xy.Add(dir)
		dirs := g.dirs(xy, dir)
		if len(dirs) != 1 {
			return xy, dir, length
		}
		dir = dirs[0]
	}
}

func (g grid) makeNet(start, end XY) *net {
	n := &net{
		vertices: map[XY]*vertex{},
	}
	n.start = n.vertex(start)
	n.end = n.vertex(end)

	type path struct {
		xy, dir XY
		origin  *vertex
	}

	q, visited := Queue[path]{}, map[XY]bool{start: true}
	q.Push(path{start, PosY, n.start})

	for q.Len() > 0 {
		p := q.Pop()
		for _, dir := range g.dirs(p.xy, p.dir) {
			xy2, dir2, length := g.travel(p.xy, dir)
			v := n.vertex(xy2)
			v.connect(p.origin, length)
			p.origin.connect(v, length)

			if !visited[v.xy] {
				visited[v.xy] = true
				q.Push(path{v.xy, dir2, v})
			}
		}
	}

	return n
}

func search2(v, end *vertex, length int, visited map[XY]bool) (int, bool) {
	if v == end {
		return length, true
	}

	visited[v.xy] = true
	defer delete(visited, v.xy)

	best, found := 0, false
	for _, e := range v.edges {
		if !visited[e.v.xy] {
			l, ok := search2(e.v, end, length+e.d, visited)
			if ok {
				best = max(best, l)
				found = true
			}
		}
	}

	return best, found
}
