package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 19
	example = `

     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 
                
.
` // keep the empty line at the end
)

func main() {
	Run(year, day, example, solve)
}

type grid = *Matrix[byte]

func solve(p *Problem) {
	g := grid(nil)
	for p.NextLine() {
		if g == nil {
			g = NewMatrix[byte](XY{len(p.Line()), 0})
		}
		CopyVector(g.AppendRow(), p.Line())
	}

	xy, dir := XY{}, PosY
	for *g.At(xy) != '|' {
		xy.X++
	}

	r1, r2 := []byte(nil), 0
	for ; ; xy, r2 = xy.Add(dir), r2+1 {
		ch := *g.At(xy)

		if ch == ' ' {
			break
		}

		if ch == '-' || ch == '|' {
			continue
		}

		if ch == '+' {
			for _, d := range HVDirs {
				b := *g.At(xy.Add(d))
				s := byte('|')
				if d.X == 0 {
					s = '-'
				}
				if dir != d.Neg() && b != ' ' && b != s {
					dir = d
					break
				}
			}

			continue
		}

		r1 = append(r1, ch)
	}
	p.PartOne(string(r1))
	p.PartOne(r2)
}
