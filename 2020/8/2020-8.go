package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 8
	example = `

nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6

`
)

func main() {
	Run(year, day, example, solve)
}

type cpu struct {
	acc  int
	pc   int
	prog []inst
}

type inst struct {
	code string
	arg  int
}

func (i inst) flip() inst {
	switch i.code {
	case "nop":
		return inst{"jmp", i.arg}
	case "jmp":
		return inst{"nop", i.arg}
	default:
		return i
	}
}

func solve(p *Problem) {
	s2 := 0
	c := &cpu{}
	for p.NextLine() {
		f := SplitFields(p.Line())
		c.prog = append(c.prog, inst{f[0], ParseInt(f[1])})
	}

	c.run()
	p.PartOne(c.acc)

	for pc, i := range c.prog {
		flipped := i.flip()
		if flipped != i {
			c.prog[pc] = flipped
			if c.run() {
				p.PartTwo(c.acc)
				break
			}
			c.prog[pc] = i
		}
	}
	p.PartTwo(s2)
}

func (c *cpu) run() bool {
	c.acc = 0
	c.pc = 0
	v := map[int]bool{}
	for {
		if c.pc >= len(c.prog) {
			return true
		}
		if v[c.pc] {
			return false
		}
		v[c.pc] = true

		i := c.prog[c.pc]
		switch i.code {
		case "acc":
			c.acc += i.arg
			c.pc++
		case "jmp":
			c.pc += i.arg
		default:
			c.pc++
		}
	}
}
