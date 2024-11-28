package main

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 13
	example = `

6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5

`
)

func main() {
	Run(year, day, example, solve)
}

type sheet map[XY]struct{}

func (s sheet) foldX(x int) {
	for c := range s {
		if c.X > x {
			delete(s, c)
			c.X = 2*x - c.X
			s[c] = struct{}{}
		}
	}
}

func (s sheet) foldY(y int) {
	for c := range s {
		if c.Y > y {
			delete(s, c)
			c.Y = 2*y - c.Y
			s[c] = struct{}{}
		}
	}
}

func (s sheet) Print() {
	mx, my := 0, 0
	dots := make([]XY, 0, len(s))
	for c := range s {
		dots = append(dots, c)
		mx = max(mx, c.X+1)
		my = max(my, c.Y+1)
	}

	blank := make([]byte, mx)
	for i := range blank {
		blank[i] = '.'
	}

	sort.Slice(dots, func(i, j int) bool {
		ci, cj := dots[i], dots[j]
		return ci.Y < cj.Y || ci.Y == cj.Y && ci.X < cj.X
	})

	row := make([]byte, mx)
	for y, i := 0, 0; y < my; y++ {
		copy(row, blank)
		for i < len(dots) && dots[i].Y == y {
			row[dots[i].X] = '#'
			i++
		}
		fmt.Println(string(row))
	}
}

func solve(p *Problem) {
	s := sheet{}
	for p.NextLine() {
		if p.Line() == "" {
			break
		}
		x, y, _ := strings.Cut(p.Line(), ",")
		s[XY{ParseInt(x), ParseInt(y)}] = struct{}{}
	}

	for first := true; p.NextLine(); first = false {
		m := p.Parse(`^fold along ([xy])=(\d+)$`)
		switch m[1] {
		case "x":
			s.foldX(ParseInt(m[2]))
		case "y":
			s.foldY(ParseInt(m[2]))
		}
		if first {
			p.PartOne(len(s))
		}
	}

	p.PartTwo("")
	s.Print()
}
