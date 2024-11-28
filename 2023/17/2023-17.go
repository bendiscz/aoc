package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 17
	example = `

2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	v int
}

type grid struct {
	*Matrix[cell]
}

type state struct {
	pos  XY
	dir  XY
	line int
}

type path struct {
	state
	cost int
}

func search(g grid, minLine, maxLine int) int {
	h := NewHeap[*path](func(p1 *path, p2 *path) bool {
		return p1.cost < p2.cost
	})

	h.Push(&path{
		state: state{
			dir:  PosX,
			line: 1,
		},
	})
	h.Push(&path{
		state: state{
			dir:  PosY,
			line: 1,
		},
	})

	visited := map[state]bool{}
	for h.Len() > 0 {
		pt := h.Pop()

		if pt.line >= minLine && pt.pos == g.Dim.Sub(XY{1, 1}) {
			return pt.cost
		}

		if pos2 := pt.pos.Add(pt.dir); g.Dim.HasInside(pos2) && pt.line < maxLine {
			next := path{
				state: state{
					pos:  pos2,
					dir:  pt.dir,
					line: pt.line + 1,
				},
				cost: pt.cost + g.At(pos2).v,
			}
			if !visited[next.state] {
				visited[next.state] = true
				h.Push(&next)
			}
		}

		if pt.line < minLine {
			continue
		}

		left := XY{pt.dir.Y, -pt.dir.X}
		if pos2 := pt.pos.Add(left); g.Dim.HasInside(pos2) {
			next := path{
				state: state{
					pos:  pos2,
					dir:  left,
					line: 1,
				},
				cost: pt.cost + g.At(pos2).v,
			}
			if !visited[next.state] {
				visited[next.state] = true
				h.Push(&next)
			}
		}

		right := XY{-pt.dir.Y, pt.dir.X}
		if pos2 := pt.pos.Add(right); g.Dim.HasInside(pos2) {
			next := path{
				state: state{
					pos:  pos2,
					dir:  right,
					line: 1,
				},
				cost: pt.cost + g.At(pos2).v,
			}
			if !visited[next.state] {
				visited[next.state] = true
				h.Push(&next)
			}
		}
	}
	return -1
}

func solve(p *Problem) {
	g := grid{}
	for p.NextLine() {
		if g.Matrix == nil {
			g.Matrix = NewMatrix[cell](Rectangle(len(p.Line()), 0))
		}
		ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{int(b - '0')} })
	}
	p.PartOne(search(g, 0, 3))
	p.PartTwo(search(g, 4, 10))
}
