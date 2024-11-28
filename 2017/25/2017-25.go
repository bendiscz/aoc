package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 25
	example = `

Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.

`
)

func main() {
	Run(year, day, example, solve)
}

type rule struct {
	input  byte
	output byte
	move   int
	next   int
}

type state struct {
	id int
	r  [2]rule
}

type program map[int]*state

func parseRule(p *Problem) rule {
	r := rule{}
	r.input = byte(ParseUints(p.ReadLine())[0])
	r.output = byte(ParseUints(p.ReadLine())[0])
	if strings.HasSuffix(p.ReadLine(), "right.") {
		r.move = 1
	} else {
		r.move = -1
	}
	p.NextLine()
	p.Scanf(" - Continue with state %c.", &r.next)
	return r
}

func solve(p *Problem) {
	start, steps := int('A'), 0
	p.NextLine()
	p.Scanf("Begin in state %c.", &start)
	p.NextLine()
	p.Scanf("Perform a diagnostic checksum after %d steps.", &steps)

	prog := program{}
	for p.NextLine() {
		s := state{}
		p.NextLine()
		p.Scanf("In state %c:", &s.id)
		r := parseRule(p)
		s.r[r.input] = r
		r = parseRule(p)
		s.r[r.input] = r
		prog[s.id] = &s
	}

	tape, pos, s := map[int]byte{}, 0, prog[start]
	for i := 0; i < steps; i++ {
		r := s.r[tape[pos]]
		if r.output != 0 {
			tape[pos] = r.output
		} else {
			delete(tape, pos)
		}

		pos += r.move
		s = prog[r.next]
	}
	p.PartOne(len(tape))
}
