package main

import (
	"slices"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 6
	example = `

123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	rows := [][]int(nil)
	lines := []string(nil)
	for p.NextLine() {
		f := ParseInts(p.Line())
		if len(f) == 0 {
			break
		}
		rows = append(rows, f)
		lines = append(lines, p.Line())
	}

	rows2 := [][]int(nil)
	nums := []int(nil)
	for i := len(lines[0]) - 1; i >= 0; i-- {
		x := 0
		ok := false
		for _, l := range lines {
			if l[i] == ' ' {
				continue
			}
			x = x*10 + int(l[i]-'0')
			ok = true
		}
		if !ok {
			rows2 = append(rows2, nums)
			nums = nil
		} else {
			nums = append(nums, x)
		}
	}
	rows2 = append(rows2, nums)
	slices.Reverse(rows2)

	ops := SplitFields(p.Line())
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		s := 0
		if op == "*" {
			s = 1
		}
		for _, r := range rows {
			if op == "+" {
				s = s + r[i]
			} else {
				s = s * r[i]
			}
		}
		s1 += s
	}

	for i := 0; i < len(ops); i++ {
		op := ops[i]
		s := 0
		if op == "*" {
			s = 1
		}
		for _, x := range rows2[i] {
			if op == "+" {
				s = s + x
			} else {
				s = s * x
			}
		}
		s2 += s
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}
