package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 3
	example = `

00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	var nums []uint
	var width int
	for p.NextLine() {
		if width == 0 {
			width = len(p.Line())
		}
		nums = append(nums, parseBinary(p.Line()))
	}

	g, e := 0, 0
	for b := width - 1; b >= 0; b-- {
		count := 0
		for _, x := range nums {
			if bit(x, b) {
				count++
			} else {
				count--
			}
		}

		g, e = 2*g, 2*e
		if count >= 0 {
			g++
		} else {
			e++
		}
	}

	p.PartOne(g * e)

	p.PartTwo(search(nums, width, false) * search(nums, width, true))
}

func parseBinary(s string) uint {
	x := uint(0)
	for _, ch := range s {
		x *= 2
		if ch == '1' {
			x += 1
		}
	}
	return x
}

func search(nums []uint, width int, inv bool) uint {
	n := len(nums)
	a := make([]uint, n)
	copy(a, nums)

	for b := width - 1; b >= 0; b-- {
		count := 0
		for i := 0; i < n; i++ {
			if bit(a[i], b) {
				count++
			} else {
				count--
			}
		}

		x, j := count >= 0 != inv, 0
		for i := 0; i < n; i++ {
			if bit(a[i], b) == x {
				a[j] = a[i]
				j++
			}
		}

		if j == 1 {
			return a[0]
		}

		n = j
	}

	return 0
}

func bit(x uint, b int) bool {
	return x&uint(1<<b) != 0
}
