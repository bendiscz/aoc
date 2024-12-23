package main

import (
	"cmp"
	"fmt"
	. "github.com/bendiscz/aoc/aoc"
	"slices"
)

const (
	year    = 2018
	day     = 13
	example = `

/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

func (c cell) String() string { return string([]byte{c.ch}) }

type grid struct {
	*Matrix[cell]
}

type cart struct {
	at    XY
	dir   XY
	phase int
}

func (c *cart) move(g grid) {
	c.at = c.at.Add(c.dir)
	switch g.At(c.at).ch {
	case '/':
		if c.dir.Y == 0 {
			c.dir = left(c.dir)
		} else {
			c.dir = right(c.dir)
		}
	case '\\':
		if c.dir.X == 0 {
			c.dir = left(c.dir)
		} else {
			c.dir = right(c.dir)
		}
	case '+':
		switch c.phase {
		case 0:
			c.dir = left(c.dir)
		case 2:
			c.dir = right(c.dir)
		}
		c.phase = (c.phase + 1) % 3
	}
}

func compareCarts(c1, c2 *cart) int {
	c := cmp.Compare(c1.at.Y, c2.at.Y)
	if c == 0 {
		c = cmp.Compare(c1.at.X, c2.at.X)
	}
	return c
}

var dirs = [128]XY{
	'>': PosX,
	'<': NegX,
	'v': PosY,
	'^': NegY,
}

func left(d XY) XY  { return XY{X: d.Y, Y: -d.X} }
func right(d XY) XY { return XY{X: -d.Y, Y: d.X} }

func solve(p *Problem) {
	g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	carts := []*cart(nil)
	for p.NextLine() {
		s, y := p.Line(), g.Dim.Y
		ParseVectorFunc(g.AppendRow(), s, func(b byte) cell { return cell{ch: b} })
		for x := 0; x < g.Dim.X; x++ {
			d := dirs[s[x]]
			if d != XY0 {
				carts = append(carts, &cart{at: XY{X: x, Y: y}, dir: d})
				if d.X == 0 {
					g.AtXY(x, y).ch = '|'
				} else {
					g.AtXY(x, y).ch = '-'
				}
			}
		}
	}

	for collision := false; !collision; {
		slices.SortFunc(carts, compareCarts)
		for i, c := range carts {
			if c == nil {
				continue
			}
			c.move(g)
			for j, c2 := range carts {
				if i != j && c2 != nil && c.at == c2.at {
					if !collision {
						p.PartOne(fmt.Sprintf("%d,%d", c.at.X, c.at.Y))
						collision = true
					}
					carts[i], carts[j] = nil, nil
				}
			}
		}
	}

	for {
		carts = slices.DeleteFunc(carts, func(c *cart) bool { return c == nil })
		if len(carts) <= 1 {
			break
		}
		slices.SortFunc(carts, compareCarts)
		for i, c := range carts {
			if c == nil {
				continue
			}
			c.move(g)
			for j, c2 := range carts {
				if i != j && c2 != nil && c.at == c2.at {
					carts[i], carts[j] = nil, nil
				}
			}
		}
	}

	for _, c := range carts {
		if c != nil {
			p.PartTwo(fmt.Sprintf("%d,%d", c.at.X, c.at.Y))
			break
		}
	}
}
