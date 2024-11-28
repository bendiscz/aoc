package main

import (
	"fmt"
	"math/bits"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 14
	example = `

flqrgnkx

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	used bool
}

func (c cell) String() string {
	if c.used {
		return "#"
	} else {
		return "."
	}
}

func solve(p *Problem) {
	seed := strings.TrimSpace(p.ReadAll())

	g := NewMatrix[cell](XY{128, 128})
	sum1 := 0
	for i := 0; i < 128; i++ {
		row := hash(fmt.Sprintf("%s-%d", seed, i))
		for j, b := range row {
			sum1 += bits.OnesCount8(b)
			for k := 0; k < 8; k++ {
				g.AtXY(8*j+k, i).used = (b & (0b10000000 >> k)) != 0
			}
		}
	}
	p.PartOne(sum1)

	sum2 := 0
	for y := 0; y < g.Dim.Y; y++ {
		for x := 0; x < g.Dim.X; x++ {
			if g.AtXY(x, y).used {
				fill(g, x, y)
				sum2++
			}
		}
	}
	p.PartTwo(sum2)
}

func hash(s string) []byte {
	const n = 256
	input := append([]byte(s), 17, 31, 73, 47, 23)

	d := make([]byte, n)
	for i := 0; i < n; i++ {
		d[i] = byte(i)
	}

	pos, skip := 0, 0
	for k := 0; k < 64; k++ {
		for _, b := range input {
			l := int(b)
			for i := 0; i < l/2; i++ {
				i1, i2 := (pos+i)%n, (pos+l-i-1)%n
				d[i1], d[i2] = d[i2], d[i1]
			}
			pos = (pos + l + skip) % n
			skip++
		}
	}

	h := make([]byte, 16)
	for i := 0; i < 16; i++ {
		x := byte(0)
		for j := 0; j < 16; j++ {
			x = x ^ d[16*i+j]
		}
		h[i] = x
	}

	return h
}

func fill(g *Matrix[cell], x, y int) {
	q := Queue[XY]{}
	q.Push(XY{x, y})
	for q.Len() > 0 {
		xy := q.Pop()
		g.At(xy).used = false
		for _, d := range HVDirs {
			a := xy.Add(d)
			if g.Dim.HasInside(a) && g.At(a).used {
				q.Push(a)
			}
		}
	}
}
