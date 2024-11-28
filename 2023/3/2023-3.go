package main

import (
	"fmt"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 3
	example = `

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	b byte
	n []int
}

func (c cell) String() string {
	return fmt.Sprintf("%c", c.b)
}

func parse(b byte) cell {
	return cell{b: b}
}

type grid = Matrix[cell]

func fill(n int) string {
	b := []byte(nil)
	for i := 0; i < n; i++ {
		b = append(b, '.')
	}
	return string(b)
}

func wrap(s string) string {
	b := make([]byte, len(s)+2)
	b[0] = '.'
	b[len(s)+1] = '.'
	copy(b[1:], s)
	return string(b)
}

func solve(p *Problem) {
	var g *grid
	for p.NextLine() {
		if g == nil {
			g = NewMatrix[cell](XY{len(p.Line()) + 2, 0})
			ParseVectorFunc(g.AppendRow(), fill(g.Dim.X), parse)
		}
		ParseVectorFunc(g.AppendRow(), wrap(p.Line()), parse)
	}
	if g == nil {
		return
	}
	ParseVectorFunc(g.AppendRow(), fill(g.Dim.X), parse)

	sum1 := 0
	for y := 1; y < g.Dim.Y-1; y++ {
		sum1 += countRow(g, y)
	}
	p.PartOne(sum1)

	sum2 := 0
	for x := 0; x < g.Dim.X; x++ {
		for y := 0; y < g.Dim.Y; y++ {
			n := g.AtXY(x, y).n
			if len(n) == 2 {
				sum2 += n[0] * n[1]
			}
		}
	}
	p.PartTwo(sum2)
}

func countRow(g *grid, y int) (sum int) {
	n := -1
	x0, x1 := 0, 0
	for x := 1; x < g.Dim.X; x++ {
		b := g.AtXY(x, y).b
		if b >= '0' && b <= '9' {
			d := int(b - '0')
			if n == -1 {
				n = 0
				x0 = x
			}
			n = n*10 + d
		} else if n >= 0 {
			x1 = x - 1
			if checkSymbols(g, n, x0, x1, y) {
				sum += n
			}
			n = -1
		}
	}
	return
}

func checkSymbols(g *grid, n, x0, x1, y int) (ok bool) {
	if checkSymbol(g, n, XY{x0 - 1, y}) {
		ok = true
	}
	if checkSymbol(g, n, XY{x1 + 1, y}) {
		ok = true
	}
	for x := x0 - 1; x <= x1+1; x++ {
		if checkSymbol(g, n, XY{x, y - 1}) {
			ok = true
		}
		if checkSymbol(g, n, XY{x, y + 1}) {
			ok = true
		}
	}
	return ok
}

func checkSymbol(g *grid, n int, xy XY) bool {
	b := g.At(xy).b
	if !isSymbol(b) {
		return false
	}

	if b == '*' {
		g.At(xy).n = append(g.At(xy).n, n)
	}
	return true
}

func isSymbol(b byte) bool {
	return b != '.' && (b < '0' || b > '9')
}
