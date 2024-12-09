package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"strings"
)

const (
	year    = 2020
	day     = 18
	example = `

1 + 2 * 3 + 4 * 5 + 6
1 + (2 * 3) + (4 * (5 + 6))
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	for p.NextLine() {
		s := strings.ReplaceAll(p.Line(), " ", "")
		x, _ := eval(s)
		s1 += x
		x, _ = evalExpr(s)
		s2 += x
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}

func eval(s string) (int, string) {
	r, mul := 0, false
loop:
	for len(s) > 0 {
		ch := s[0]
		s = s[1:]

		var x int
		switch ch {
		case ')':
			break loop
		case '(':
			x, s = eval(s)
		case '+':
			mul = false
			continue loop
		case '*':
			mul = true
			continue loop
		default:
			x = int(ch - '0')
		}
		if mul {
			r *= x
		} else {
			r += x
		}
	}
	return r, s
}

func evalExpr(s string) (int, string) {
	var x int
	x, s = evalMul(s)
	if len(s) > 0 && s[0] == ')' {
		s = s[1:]
	}
	return x, s
}

func evalMul(s string) (int, string) {
	var x, y int
	x, s = evalAdd(s)
	for len(s) > 0 && s[0] == '*' {
		y, s = evalAdd(s[1:])
		x *= y
	}
	return x, s
}

func evalAdd(s string) (int, string) {
	var x, y int
	x, s = evalNum(s)
	for len(s) > 0 && s[0] == '+' {
		y, s = evalNum(s[1:])
		x += y
	}
	return x, s
}

func evalNum(s string) (int, string) {
	if s[0] == '(' {
		return evalExpr(s[1:])
	}
	return int(s[0] - '0'), s[1:]
}
