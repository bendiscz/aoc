package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 9
	example = `

London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141

`
)

func main() {
	Run(year, day, example, solve)
}

type world struct {
	names map[string]int
	dists [][]int
}

func (w *world) index(n string) int {
	if i, ok := w.names[n]; ok {
		return i
	}
	i := len(w.names)
	w.names[n] = i
	return i
}

func (w *world) add(n1, n2 string, d int) {
	i1, i2 := w.index(n1), w.index(n2)
	if i1 > i2 {
		i1, i2 = i2, i1
	}
	i2 = i2 - i1 - 1
	for i1 >= len(w.dists) {
		w.dists = append(w.dists, nil)
	}
	for i2 >= len(w.dists[i1]) {
		w.dists[i1] = append(w.dists[i1], 0)
	}
	w.dists[i1][i2] = d
}

func (w *world) dist(i1, i2 int) int {
	if i1 > i2 {
		i1, i2 = i2, i1
	}
	return w.dists[i1][i2-i1-1]
}

func (w *world) search(v map[int]struct{}, s, p int, max bool) int {
	v[s] = struct{}{}
	defer delete(v, s)

	if len(v) == len(w.names) {
		return p
	}

	best := 0
	if !max {
		best = math.MaxInt
	}

	for t := 0; t < len(w.names); t++ {
		if _, ok := v[t]; ok {
			continue
		}

		x := w.dist(s, t)
		if x == 0 {
			continue
		}

		c := w.search(v, t, p+x, max)
		if !max && c < best || max && c > best {
			best = c
		}
	}

	return best
}

func solve(p *Problem) {
	w := world{names: map[string]int{}}
	for p.NextLine() {
		m := p.Parse(`^(\w+) to (\w+) = (\d+)$`)
		w.add(m[1], m[2], ParseInt(m[3]))
	}

	x1, x2 := math.MaxInt, 0
	for s := 0; s < len(w.names); s++ {
		x1 = min(w.search(map[int]struct{}{}, s, 0, false), x1)
		x2 = max(w.search(map[int]struct{}{}, s, 0, true), x2)
	}

	p.PartOne(x1)
	p.PartTwo(x2)
}
