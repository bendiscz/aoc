package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 12
	example = `

0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2

`
)

const N = 3

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	v byte
}

func (c cell) String() string {
	switch c.v {
	case 0:
		return "."
	default:
		return string([]byte{c.v})
	}
}

type grid struct {
	*Matrix[cell]
}

type shape struct {
	code     byte
	area     int
	variants []grid
}

func newShape(g grid, index int) shape {
	area := 0
	for _, c := range g.All() {
		if c.v == '#' {
			area++
		}
	}

	variants := []grid(nil)
	g2 := Grid[cell](g)
	for i := 0; i < 8; i++ {
		variants = append(variants, grid{CloneGrid(g2)})
		g2 = g2.RotateL()
		if i == 3 {
			g2 = g2.FlipX()
		}
	}

	f := []grid(nil)
	for i, v := range variants {
		unique := true
		for j := 0; j < i && unique; j++ {
			same := true
			for c := range v.Dim.All() {
				if v.At(c).v != variants[j].At(c).v {
					same = false
					break
				}
			}
			if same {
				unique = false
			}
		}

		if unique {
			f = append(f, v)
		}
	}

	return shape{
		code:     byte('A' + index),
		area:     area,
		variants: f,
	}
}

func solve(p *Problem) {
	areas := []int(nil)
	shapes := []shape(nil)
	for p.NextLine() {
		if !strings.HasSuffix(p.Line(), ":") {
			break
		}

		area := 0
		g := grid{NewMatrix[cell](Rectangle(N, 0))}
		for p.NextLine() && p.Line() != "" {
			for _, ch := range p.Line() {
				if ch == '#' {
					area++
				}
			}
			ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{v: b} })
		}
		areas = append(areas, area)
		shapes = append(shapes, newShape(g, len(shapes)))
	}

	s1, s2 := 0, 0
	for {
		//fmt.Printf("testing %s\n", p.Line())
		f := ParseInts(p.Line())
		area := 0
		for i, x := range f[2:] {
			area += x * areas[i]
		}
		if area <= f[0]*f[1] {
			s1++
		}

		if canFit(Rectangle(f[0], f[1]), f[2:], shapes) {
			s2++
		}

		if !p.NextLine() {
			break
		}
	}

	p.PartOne(s1)
	p.PartOne(s2)
}

func canFit(dim XY, counts []int, shapes []shape) bool {
	total := 0
	for i := range counts {
		total += counts[i] * shapes[i].area
	}
	if total > dim.X*dim.Y {
		return false
	}

	return fill(grid{NewMatrix[cell](dim)}, counts, make([]Stack[XY], len(counts)), shapes)
}

func fill(g grid, counts []int, placements []Stack[XY], shapes []shape) bool {
	if done(counts) {
		//PrintGrid(g)
		//fmt.Printf("---\n")
		return true
	}

	for i := range counts {
		if counts[i] == 0 {
			continue
		}

		counts[i]--

		s := shapes[i]
		orig := XY0
		if placements[i].Len() > 0 {
			orig = placements[i].Top()
		}

		for at := range g.Dim.Sub(Square(N - 1)).All() {
			if at.X <= orig.X && at.Y <= orig.Y {
				continue
			}

			for _, v := range s.variants {
				if !place(g, v, at, s.code) {
					continue
				}
				placements[i].Push(at)
				if fill(g, counts, placements, shapes) {
					return true
				}
				placements[i].Pop()
				remove(g, v, at)
			}
		}

		counts[i]++
		break
	}

	return false
}

func done(counts []int) bool {
	for _, x := range counts {
		if x != 0 {
			return false
		}
	}
	return true
}

func place(g, v grid, at XY, code byte) bool {
	for c := range v.Dim.All() {
		if v.At(c).v == '#' && g.At(c.Add(at)).v != 0 {
			return false
		}
	}

	for c := range v.Dim.All() {
		if v.At(c).v == '#' {
			g.At(c.Add(at)).v = code
		}
	}
	return true
}

func remove(g, v grid, at XY) {
	for c := range v.Dim.All() {
		if v.At(c).v == '#' {
			g.At(c.Add(at)).v = 0
		}
	}
}
