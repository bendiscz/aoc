package main

import (
	_ "embed"
	. "github.com/bendiscz/aoc/aoc"
)

//go:embed Day18challengeM.txt
var input string

func main() {
	Example(input).Run(solve)
}

type cell struct {
	blocked bool
	visited bool
	touched bool
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	const N = 213
	g := grid{NewMatrix[cell](Square(N))}
	bs := []XY(nil)
	for p.NextLine() {
		f := ParseInts(p.Line())
		xy := XY{X: f[0], Y: f[1]}
		bs = append(bs, xy)
		g.At(xy).blocked = true
	}

	i, xy, f := len(bs), XY0, true
	for i >= 0 {
		if f && fill(g, xy) {
			if i < len(bs) {
				p.Printf("%d,%d", bs[i].X, bs[i].Y)
			} else {
				p.Printf("path not blocked")
			}
			break
		}

		i--
		c := g.At(bs[i])
		c.blocked = false
		if c.touched {
			f = true
			xy = bs[i]
		} else {
			f = false
		}
	}
}

func fill(g grid, xy XY) bool {
	q := Queue[XY]{}
	q.Push(xy)
	for q.Len() > 0 {
		xy = q.Pop()
		for _, d := range HVDirs {
			n := xy.Add(d)
			if !g.Dim.HasInside(n) {
				continue
			}
			c := g.At(n)
			if c.blocked {
				c.touched = true
				continue
			}
			if c.visited {
				continue
			}
			if n.X == g.Dim.X-1 && n.Y == g.Dim.Y-1 {
				return true
			}
			q.Push(n)
			c.visited = true
		}
	}
	return false
}
