package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 13
	example = `

Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	for p.NextLine() {
		var ax, ay, bx, by, px, py int
		p.Scanf("Button A: X+%d, Y+%d", &ax, &ay)
		p.NextLine()
		p.Scanf("Button B: X+%d, Y+%d", &bx, &by)
		p.NextLine()
		p.Scanf("Prize: X=%d, Y=%d", &px, &py)
		p.NextLine()

		if a, b, ok := solveAB(ax, ay, bx, by, px, py); ok && a <= 100 && b <= 100 {
			s1 += 3*a + b
		}
		if a, b, ok := solveAB(ax, ay, bx, by, px+10000000000000, py+10000000000000); ok {
			s2 += 3*a + b
		}
	}
	p.PartOne(s1)
	p.PartTwo(s2)
}

func solveAB(ax, ay, bx, by, px, py int) (int, int, bool) {
	d := by*ax - ay*bx
	if d == 0 {
		return 0, 0, false
	}
	b := (py*ax - ay*px) / d
	a := (px - b*bx) / ax

	if ax*a+bx*b == px && ay*a+by*b == py {
		return a, b, true
	}
	return 0, 0, false
}
