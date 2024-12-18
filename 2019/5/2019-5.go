package main

import (
	"github.com/bendiscz/aoc/2019/intcode"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 5
	example = `

3,0,4,0,99

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	prog := intcode.Parse(p)
	p.PartOne(run(prog, 1))
	p.PartOne(run(prog, 5))
}

func run(prog *intcode.Program, input int) int {
	output := prog.Exec().ReadWrite(input)
	return output[len(output)-1]
}
