package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/bendiscz/aoc/2021/22/alt"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2021
	day  = 22

	alternative = false
)

//go:embed example
var example string

func main() {
	solve := solve
	if alternative {
		solve = alt.Solve
	}
	Run(year, day, example, solve)
}

type axis struct {
	p, q int
}

type box struct {
	a      [3]axis
	weight int
}

func (b box) volume() int { return b.a[0].size() * b.a[1].size() * b.a[2].size() }

func (a axis) size() int              { return a.q - a.p + 1 }
func (a axis) intersects(b axis) bool { return a.q >= b.p && a.p <= b.q }
func (a axis) small() bool            { return a.p >= -50 && a.q <= 50 }

func intersect(c1, c2 box) (box, bool) {
	for i := 0; i < 3; i++ {
		if !c1.a[i].intersects(c2.a[i]) {
			return box{}, false
		}
	}

	return box{
		a: [3]axis{
			{Max(c1.a[0].p, c2.a[0].p), Min(c1.a[0].q, c2.a[0].q)},
			{Max(c1.a[1].p, c2.a[1].p), Min(c1.a[1].q, c2.a[1].q)},
			{Max(c1.a[2].p, c2.a[2].p), Min(c1.a[2].q, c2.a[2].q)},
		},
		weight: 0,
	}, true
}

func count(boxes []box) (volume int) {
	for _, b := range boxes {
		volume += b.weight * b.volume()
	}
	return
}

func solve(p *Problem) {
	t0 := time.Now()
	partOneDone := false
	var boxes []box
	for p.NextLine() {
		m := p.Parse(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)
		x1, x2 := ParseInt(m[2]), ParseInt(m[3])
		y1, y2 := ParseInt(m[4]), ParseInt(m[5])
		z1, z2 := ParseInt(m[6]), ParseInt(m[7])
		bc := box{a: [...]axis{{x1, x2}, {y1, y2}, {z1, z2}}, weight: 1}

		if !partOneDone && (!bc.a[0].small() || !bc.a[1].small() || !bc.a[2].small()) {
			p.PartOne(fmt.Sprintf("%d (time %v)", count(boxes), time.Since(t0)))
			partOneDone = true
		}

		for _, b := range boxes[:] {
			if bi, ok := intersect(bc, b); ok {
				bi.weight = -1 * b.weight
				boxes = append(boxes, bi)
			}
		}

		if m[1] == "on" {
			boxes = append(boxes, bc)
		}
	}

	p.PartTwo(count(boxes))
}
