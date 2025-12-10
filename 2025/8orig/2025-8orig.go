package main

import (
	"cmp"
	"fmt"
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

type Box struct {
	x, y, z int
	marked  bool
	conn    []int
}

func (b *Box) String() string {
	return fmt.Sprintf("%d,%d,%d", b.x, b.y, b.z)
}

type Pair struct {
	d      int
	b1, b2 int
}

func solve(p *Problem) {
	s1, s2 := 1, 0

	boxes := []*Box(nil)
	for p.NextLine() {
		f := ParseInts(p.Line())
		boxes = append(boxes, &Box{f[0], f[1], f[2], false, nil})
	}

	pairs := []Pair(nil)
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			x, y, z := boxes[i].x-boxes[j].x, boxes[i].y-boxes[j].y, boxes[i].z-boxes[j].z
			d := x*x + y*y + z*z
			pairs = append(pairs, Pair{d, i, j})
		}
	}

	slices.SortFunc(pairs, func(p1, p2 Pair) int {
		return cmp.Compare(p1.d, p2.d)
	})

	p.Printf("%d", len(boxes))
	for i := 0; i < len(boxes); i++ {
		b1 := boxes[pairs[i].b1]
		b2 := boxes[pairs[i].b2]
		//p.Printf("conn %v, %v", b1, b2)
		//p.Printf("conn %v, %v", pairs[i].b1, pairs[i].b2)

		b1.conn = append(b1.conn, pairs[i].b2)
		b2.conn = append(b2.conn, pairs[i].b1)
	}

	sizes := []int(nil)
	for i := 0; i < len(boxes); i++ {
		if boxes[i].marked {
			continue
		}

		x := 0
		s := &Stack[int]{}
		s.Push(i)
		for s.Len() > 0 {
			j := s.Pop()
			b := boxes[j]

			if b.marked {
				continue
			}

			b.marked = true
			x++

			for _, k := range b.conn {
				if !boxes[k].marked {
					s.Push(k)
				}
			}
		}
		//p.Printf("%d", x)
		sizes = append(sizes, x)
	}

	slices.Sort(sizes)
	slices.Reverse(sizes)
	s1 = sizes[0] * sizes[1] * sizes[2]

	p.PartOne(s1)

	for i := len(boxes); i < len(pairs); i++ {
		b1 := boxes[pairs[i].b1]
		b2 := boxes[pairs[i].b2]
		//p.Printf("conn %v, %v", b1, b2)
		//p.Printf("conn %v, %v", pairs[i].b1, pairs[i].b2)

		b1.conn = append(b1.conn, pairs[i].b2)
		b2.conn = append(b2.conn, pairs[i].b1)

		for _, b := range boxes {
			b.marked = false
		}

		x := 0
		s := &Stack[int]{}
		s.Push(0)
		for s.Len() > 0 {
			j := s.Pop()
			b := boxes[j]

			if b.marked {
				continue
			}

			b.marked = true
			x++

			for _, k := range b.conn {
				if !boxes[k].marked {
					s.Push(k)
				}
			}
		}

		if x == len(boxes) {
			s2 = b1.x * b2.x
			break
		}
	}

	p.PartTwo(s2)
}
