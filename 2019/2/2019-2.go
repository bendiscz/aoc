package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"slices"
)

const (
	year    = 2019
	day     = 2
	example = `

1,9,10,3,2,3,11,0,99,30,40,50

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	p.NextLine()
	prog := ParseInts(p.Line())

	a, b := 0, 0
	if !p.Example() {
		a, b = 12, 2
	}
	p.PartOne(run(prog, a, b))

	if p.Example() {
		return
	}

	x := 19690720

	x0 := run(prog, 0, 0)
	d := run(prog, 1, 0) - x0
	a = (x - x0) / d
	b = x - (a*d + x0)
	p.PartTwo(a*100 + b)
}

func run(prog []int, a, b int) int {
	prog = slices.Clone(prog)
	prog[1] = a
	prog[2] = b
	pc := 0
	for {
		switch prog[pc] {
		case 1:
			prog[prog[pc+3]] = prog[prog[pc+1]] + prog[prog[pc+2]]
			pc += 4
		case 2:
			prog[prog[pc+3]] = prog[prog[pc+1]] * prog[prog[pc+2]]
			pc += 4
		case 99:
			return prog[0]
		default:
			panic("invalid opcode")
		}
	}
}
