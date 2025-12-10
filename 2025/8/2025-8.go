package main

import (
	"cmp"
	"slices"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 8
	example = `

162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689

`
)

func main() {
	Run(year, day, example, solve)
}

type box struct {
	x, y, z int
	index   int
	comp    *box
	size    int
}

func newBox(x, y, z, index int) *box {
	return &box{
		x:     x,
		y:     y,
		z:     z,
		index: index,
		size:  1,
	}
}

func (b *box) pivot() *box {
	if b.comp == nil {
		return b
	}
	return b.comp.pivot()
}

func connect(b1, b2 *box) *box {
	p1, p2 := b1.pivot(), b2.pivot()
	if p1 == p2 {
		return p1
	}
	if p1.index > p2.index {
		p1, p2 = p2, p1
	}
	p1.size += p2.size
	p2.comp = p1
	return p1
}

type pair struct {
	b1, b2 *box
	d      int
}

func solve(p *Problem) {
	boxes := []*box(nil)
	for p.NextLine() {
		f := ParseInts(p.Line())
		boxes = append(boxes, newBox(f[0], f[1], f[2], len(boxes)))
	}

	pairs := []pair(nil)
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			b1, b2 := boxes[i], boxes[j]
			dx, dy, dz := b1.x-b2.x, b1.y-b2.y, b1.z-b2.z
			d := dx*dx + dy*dy + dz*dz
			pairs = append(pairs, pair{b1, b2, d})
		}
	}

	slices.SortFunc(pairs, func(p1, p2 pair) int { return cmp.Compare(p1.d, p2.d) })

	limit := 1000
	if p.Example() {
		limit = 10
	}

	for i := 0; i < len(pairs); i++ {
		if i == limit {
			comps := []int(nil)
			for _, b := range boxes {
				if b.comp == nil {
					comps = append(comps, b.pivot().size)
				}
			}
			slices.SortFunc(comps, func(x, y int) int { return cmp.Compare(y, x) })
			p.PartOne(comps[0] * comps[1] * comps[2])
		}

		bb := pairs[i]
		b := connect(bb.b1, bb.b2)
		if b.size == len(boxes) {
			p.PartTwo(bb.b1.x * bb.b2.x)
			break
		}
	}
}
