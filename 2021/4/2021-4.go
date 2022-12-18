package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 4
	example = `

7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7

`
)

func main() {
	Run(year, day, example, solve)
}

type board struct {
	num  [5][5]int
	hit  [5][5]bool
	r, c [5]int
	done bool
}

func (b *board) mark(num int) bool {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if b.num[r][c] != num {
				continue
			}
			if !b.hit[r][c] {
				b.hit[r][c] = true
				b.r[r]++
				b.c[c]++

				if b.r[r] == 5 || b.c[c] == 5 {
					b.done = true
				}
			}
			return b.done
		}
	}
	return false
}

func (b *board) score() int {
	score := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !b.hit[r][c] {
				score += b.num[r][c]
			}
		}
	}
	return score
}

func solve(p *Problem) {
	var nums []int
	p.NextLine()
	for _, str := range strings.Split(p.Line(), ",") {
		nums = append(nums, ParseInt(str))
	}

	var boards []*board
	for p.NextLine() {
		b := &board{}
		for r := 0; r < 5; r++ {
			p.NextLine()
			f := strings.Fields(p.Line())
			for c := 0; c < 5; c++ {
				b.num[r][c] = ParseInt(f[c])
			}
		}

		boards = append(boards, b)
	}

	var fb, lb *board
	var fn, ln int
	for _, num := range nums {
		for _, b := range boards {
			if b.done {
				continue
			}

			if b.mark(num) {
				if fb == nil {
					fb = b
					fn = num
				}
				lb = b
				ln = num
			}
		}
	}

	p.PartOne(fb.score() * fn)
	p.PartTwo(lb.score() * ln)
}
