package main

import (
	"fmt"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 22
	example = `

1
2
3
2024
`
)

func main() {
	Run(year, day, example, solve)
}

const M = 16777216

func next(x int) int {
	x = x ^ (x << 6)
	x %= M
	x = x ^ (x >> 5)
	x %= M
	x = x ^ (x << 11)
	x %= M
	return x
}

func extract(s []int, d1, d2, d3, d4 int) int {
	d := [4]int{s[1] - s[0], s[2] - s[1], s[3] - s[2], s[4] - s[3]}
	for i := 4; i < len(s); i++ {
		if i > 4 {
			copy(d[:], d[1:])
			d[3] = s[i] - s[i-1]
		}

		if d1 == d[0] && d2 == d[1] && d3 == d[2] && d4 == d[3] {
			return s[i]
		}
	}
	return 0
}

type diff [4]int

type diffs struct {
	m map[diff]int
}

func makeDiffs(s []int) diffs {
	ds := diffs{m: map[diff]int{}}
	d := diff{0, s[1] - s[0], s[2] - s[1], s[3] - s[2]}
	for i := 4; i < len(s); i++ {
		copy(d[:], d[1:])
		d[3] = s[i] - s[i-1]
		if _, ok := ds.m[d]; !ok {
			ds.m[d] = s[i]
		}
	}
	return ds
}

func (ds diffs) extract(d0, d1, d2, d3 int) int {
	return ds.m[diff{d0, d1, d2, d3}]
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	//g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}

	//s := p.ReadAll()

	secrets := [][]int(nil)
	for p.NextLine() {
		x := ParseInt(p.Line())

		s := []int{x % 10}
		for i := 0; i < 2000; i++ {
			x = next(x)
			s = append(s, x%10)

		}
		secrets = append(secrets, s)
		s1 += x
	}
	p.PartOne(s1)

	//testIt(secrets[1])

	ds := make([]diffs, len(secrets))
	for i, s := range secrets {
		ds[i] = makeDiffs(s)
	}

	best := 0
	for d1 := -9; d1 <= 9; d1++ {
		for d2 := -9; d2 <= 9; d2++ {
			for d3 := -9; d3 <= 9; d3++ {
				for d4 := -9; d4 <= 9; d4++ {
					sum := 0
					for _, d := range ds {
						sum += d.extract(d1, d2, d3, d4)
					}
					best = max(best, sum)
				}
			}
		}
	}

	s2 = best

	p.PartTwo(s2)

	//os.Exit(1)
}

func testIt(s []int) {
	fmt.Println(len(s))
	for i := 5; i < len(s); i++ {
		if s[i] == 7 && s[i-3]-s[i-4] == -2 {
			fmt.Printf("%v %d %d %d %d\n", s[i-4:i+1], s[i-3]-s[i-4], s[i-2]-s[i-3], s[i-1]-s[i-2], s[i-0]-s[i-1])
		}
	}
}
