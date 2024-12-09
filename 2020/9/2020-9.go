package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"math"
)

const (
	year    = 2020
	day     = 9
	example = `

35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	n := 25
	if p.Example() {
		n = 5
	}

	nums := []int(nil)
	for p.NextLine() {
		nums = append(nums, ParseInt(p.Line()))
	}

	s1 := 0
	w := make([]int, 0, n)
	s := map[int]bool{}
	k := 0
	for _, x := range nums {
		if len(w) < n {
			s[x] = true
			w = append(w, x)
			continue
		}

		found := false
		for _, v := range w {
			if s[x-v] {
				found = true
				break
			}
		}

		if !found {
			s1 = x
			break
		}

		delete(s, w[k])
		s[x] = true
		w[k] = x
		k = (k + 1) % n
	}
	p.PartOne(s1)

	i, j := 0, 1
	sum := nums[i] + nums[j]
loop:
	for {
		switch {
		case sum == s1:
			break loop
		case sum < s1 || j-i < 2:
			j++
			sum += nums[j]
		case sum > s1:
			sum -= nums[i]
			i++
		}
	}

	a, b := math.MaxInt, math.MinInt
	for k = i; k <= j; k++ {
		a = min(a, nums[k])
		b = max(b, nums[k])
	}
	p.PartTwo(a + b)
}
