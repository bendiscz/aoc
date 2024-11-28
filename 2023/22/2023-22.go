package main

import (
	"cmp"
	"slices"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 22
	example = `

1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9

`
)

func main() {
	Run(year, day, example, solve)
}

type brick struct {
	id int

	x0, y0, z0 int
	x1, y1, z1 int

	supportedBy []*brick
	supports    []*brick
}

func (b *brick) above(b2 *brick) (bool, int) {
	if b.z0 <= b2.z0 {
		return false, 0
	}
	return overlaps(b.x0, b.x1, b2.x0, b2.x1) && overlaps(b.y0, b.y1, b2.y0, b2.y1), b.z0 - b2.z1 - 1
}

func overlaps(a0, a1, b0, b1 int) bool {
	return max(a0, b0) <= min(a1, b1)
}

func solve(p *Problem) {
	bricks := []*brick(nil)
	for i := 1; p.NextLine(); i++ {
		f := ParseInts(p.Line())
		bricks = append(bricks, &brick{id: i, x0: f[0], y0: f[1], z0: f[2], x1: f[3], y1: f[4], z1: f[5]})
	}

	settle(bricks)

	count1 := 0
	for _, b := range bricks {
		supports := false
		for _, b2 := range b.supports {
			if len(b2.supportedBy) == 1 {
				supports = true
				break
			}
		}

		if !supports {
			count1++
		}
	}
	p.PartOne(count1)

	count2 := 0
	for _, b := range bricks {
		count2 += countAll(b, make(map[int]bool))
	}
	p.PartTwo(count2)
}

func settle(bricks []*brick) {
	slices.SortFunc(bricks, func(a, b *brick) int { return cmp.Compare(a.z0, b.z0) })
	settled := []*brick(nil)

	for _, b := range bricks {
		gap, support := b.z0-1, []*brick(nil)
		for _, b2 := range settled {
			if b2.z1+gap < b.z0-1 {
				break
			}
			if ok, d := b.above(b2); ok {
				if d < gap {
					gap = d
					support = []*brick{b2}
				} else if d == gap {
					support = append(support, b2)
				}
			}
		}

		b.z0 -= gap
		b.z1 -= gap
		b.supportedBy = append(b.supportedBy, support...)
		for _, b2 := range support {
			b2.supports = append(b2.supports, b)
		}

		index, _ := slices.BinarySearchFunc(settled, b, func(b1 *brick, b2 *brick) int {
			return cmp.Compare(b2.z1, b1.z1)
		})
		settled = slices.Insert(settled, index, b)
	}
}

func countAll(b *brick, gone map[int]bool) int {
	gone[b.id] = true
	if len(b.supports) == 0 {
		return 0
	}

	count := 0
	for _, b2 := range b.supports {
		free := true
		for _, b3 := range b2.supportedBy {
			if !gone[b3.id] {
				free = false
				break
			}
		}
		if free {
			count += 1 + countAll(b2, gone)
		}
	}
	return count
}
