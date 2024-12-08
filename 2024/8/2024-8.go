package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 8
	example = `

............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	dim := XY{X: len(p.PeekLine())}
	ant := map[byte][]XY{}
	for p.NextLine() {
		s := p.Line()
		for x := 0; x < len(s); x++ {
			if s[x] != '.' {
				ant[s[x]] = append(ant[s[x]], XY{x, dim.Y})
			}
		}
		dim.Y++
	}

	anti1, anti2 := map[XY]struct{}{}, map[XY]struct{}{}
	for _, v := range ant {
		for i := 0; i < len(v); i++ {
			for j := 0; j < len(v); j++ {
				if i == j {
					continue
				}
				d := v[i].Sub(v[j])
				for k, xy := 0, v[i]; dim.HasInside(xy); k, xy = k+1, xy.Add(d) {
					if k == 1 {
						anti1[xy] = struct{}{}
					}
					anti2[xy] = struct{}{}
				}
			}
		}
	}

	p.PartOne(len(anti1))
	p.PartTwo(len(anti2))
}
