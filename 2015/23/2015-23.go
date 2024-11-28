package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 23
	example = `

inc a
jio a, +2
tpl a
inc a

`
)

func main() {
	Run(year, day, example, solve)
}

type operand struct {
	direct bool
	reg    int
	value  int
}

func parseOperand(s string) operand {
	if s[0] >= 'a' && s[0] <= 'b' {
		return operand{reg: int(s[0] - 'a')}
	} else {
		return operand{direct: true, value: ParseInt(s)}
	}
}

type instruction struct {
	code     string
	op1, op2 operand
}

type program []instruction

type cpu struct {
	reg [2]int
	pc  int
}

func (c *cpu) lookup(op operand) int {
	if op.direct {
		return op.value
	} else {
		return c.reg[op.reg]
	}
}

func (c *cpu) step(prog program) bool {
	inst := prog[c.pc]
	switch inst.code {
	case "hlf":
		if !inst.op1.direct {
			c.reg[inst.op1.reg] /= 2
		}
		c.pc++

	case "tpl":
		if !inst.op1.direct {
			c.reg[inst.op1.reg] *= 3
		}
		c.pc++

	case "inc":
		if !inst.op1.direct {
			c.reg[inst.op1.reg]++
		}
		c.pc++

	case "jmp":
		c.pc += c.lookup(inst.op1)

	case "jie":
		if c.lookup(inst.op1)&1 == 0 {
			c.pc += c.lookup(inst.op2)
		} else {
			c.pc++
		}

	case "jio":
		if c.lookup(inst.op1) == 1 {
			c.pc += c.lookup(inst.op2)
		} else {
			c.pc++
		}
	}

	return c.pc >= 0 && c.pc < len(prog)
}

func solve(p *Problem) {
	prog := program(nil)
	for p.NextLine() {
		f := SplitFields(p.Line())
		inst := instruction{
			code: f[0],
			op1:  parseOperand(f[1]),
		}
		if len(f) > 2 {
			inst.op2 = parseOperand(f[2])
		}
		prog = append(prog, inst)
	}

	p.PartOne(runProgram(prog, 0))
	p.PartTwo(runProgram(prog, 1))
}

func runProgram(prog program, a int) int {
	c := cpu{}
	c.reg[0] = a
	for {
		//fmt.Printf("%d %v\n", c.pc, c.reg)
		if !c.step(prog) {
			return c.reg[1]
		}
	}
}
