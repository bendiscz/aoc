package main

import (
	"github.com/bendiscz/aoc/2019/intcode"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 9
	example = `

109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	prog := intcode.Parse(p)

	out := prog.Exec().ReadWrite(1)
	p.PartOne(out[len(out)-1])

	out = prog.Exec().ReadWrite(2)
	p.PartTwo(out[len(out)-1])
}
