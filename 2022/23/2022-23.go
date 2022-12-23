package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 23
	example = `

..............
..............
.......#......
.....###.#....
...#...#.#....
....#...##....
...#.###......
...##.#.##....
....#..#......
..............
..............
..............

`
)

func main() {
	Run(year, day, example, solve)
}

func add(c XY, x, y int) XY {
	return c.Add(XY{x, y})
}

func iterate(elves map[XY]bool, dir int) bool {
	prop := make(map[XY]XY, len(elves))
loop:
	for c := range elves {
		n, s, w, e := elves[add(c, 0, -1)], elves[add(c, 0, 1)], elves[add(c, -1, 0)], elves[add(c, 1, 0)]
		ne, nw, se, sw := elves[add(c, 1, -1)], elves[add(c, -1, -1)], elves[add(c, 1, 1)], elves[add(c, -1, 1)]
		if !n && !s && !w && !e && !ne && !nw && !se && !sw {
			continue
		}
		for i := 0; i < 4; i++ {
			switch (dir + i) % 4 {
			case 0:
				if !n && !ne && !nw {
					c2 := add(c, 0, -1)
					if _, ok := prop[c2]; ok {
						delete(prop, c2)
					} else {
						prop[c2] = c
					}
					//prop[c2] = append(prop[c2], c)
					continue loop
				}
			case 1:
				if !s && !se && !sw {
					c2 := add(c, 0, 1)
					if _, ok := prop[c2]; ok {
						delete(prop, c2)
					} else {
						prop[c2] = c
					}
					//prop[c2] = append(prop[c2], c)
					continue loop
				}
			case 2:
				if !w && !nw && !sw {
					c2 := add(c, -1, 0)
					if _, ok := prop[c2]; ok {
						delete(prop, c2)
					} else {
						prop[c2] = c
					}
					//prop[c2] = append(prop[c2], c)
					continue loop
				}
			case 3:
				if !e && !ne && !se {
					c2 := add(c, 1, 0)
					if _, ok := prop[c2]; ok {
						delete(prop, c2)
					} else {
						prop[c2] = c
					}
					//prop[c2] = append(prop[c2], c)
					continue loop
				}
			}
		}
	}

	if len(prop) == 0 {
		return false
	}

	moved := false
	for c, p := range prop {
		//if len(p) > 1 {
		//	continue
		//}
		moved = true
		delete(elves, p /*[0]*/)
		elves[c] = true
	}
	return moved
}

func bounds(elves map[XY]bool) (min, max XY) {
	min, max = XY{math.MaxInt, math.MaxInt}, XY{math.MinInt, math.MinInt}
	for c := range elves {
		min = XY{Min(min.X, c.X), Min(min.Y, c.Y)}
		max = XY{Max(max.X, c.X), Max(max.Y, c.Y)}
	}
	return
}

func count(elves map[XY]bool) int {
	min, max := bounds(elves)
	d := max.Sub(min).Add(XY{1, 1})
	return d.X*d.Y - len(elves)
}

func solve(p *Problem) {
	elves := map[XY]bool{}
	for y := 0; p.NextLine(); y++ {
		for x := 0; x < len(p.Line()); x++ {
			if p.Line()[x] == '#' {
				elves[XY{x, y}] = true
			}
		}
	}

	n := 0
	for n < 10 {
		iterate(elves, n%4)
		n++
	}
	p.PartOne(count(elves))

	for iterate(elves, n%4) {
		n++
	}
	p.PartTwo(n + 1)
}
