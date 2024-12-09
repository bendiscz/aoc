package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"slices"
)

const (
	year    = 2020
	day     = 10
	example = `

28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	jolts := []int{0}
	for p.NextLine() {
		jolts = append(jolts, ParseInt(p.Line()))
	}

	slices.Sort(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	j1, j3 := 0, 0
	last := 0
	for _, j := range jolts {
		d := j - last
		last = j

		switch d {
		case 1:
			j1++
		case 3:
			j3++
		}
	}
	p.PartOne(j1 * j3)

	n := make([]int, len(jolts))
	n[0] = 1

	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n) && jolts[j]-jolts[i] <= 3; j++ {
			n[j] += n[i]
		}
	}
	p.PartTwo(n[len(jolts)-1])
}
