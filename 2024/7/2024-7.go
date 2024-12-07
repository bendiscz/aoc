package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 7
	example = `

190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	for p.NextLine() {
		f := ParseInts(p.Line())
		if check(f[0], f[1:], false) {
			s1 += f[0]
			s2 += f[0]
		} else if check(f[0], f[1:], true) {
			s2 += f[0]
		}
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}

func check(sum int, nums []int, part2 bool) (result bool) {
	if len(nums) == 1 {
		return sum == nums[0]
	}

	l, r := nums[:len(nums)-1], nums[len(nums)-1]
	if r > sum {
		return false
	}

	if sum > r && check(sum-r, l, part2) {
		return true
	}

	if sum%r == 0 && check(sum/r, l, part2) {
		return true
	}

	if part2 {
		for d := 10; d < sum; d *= 10 {
			if sum%d == r {
				return check(sum/d, l, part2)
			}
		}
	}

	return false
}
