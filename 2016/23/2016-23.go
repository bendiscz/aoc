package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 23
	example = `

cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a

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
	if s[0] >= 'a' && s[0] <= 'd' {
		return operand{reg: int(s[0] - 'a')}
	} else {
		return operand{direct: true, value: ParseInt(s)}
	}
}

type instruction struct {
	code     string
	op1, op2 operand
}

type program []*instruction

func (p program) clone() program {
	p2 := make(program, len(p))
	for i, inst := range p {
		inst2 := *inst
		p2[i] = &inst2
	}
	return p2
}

type cpu struct {
	reg [4]int
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
	case "cpy":
		if !inst.op2.direct {
			c.reg[inst.op2.reg] = c.lookup(inst.op1)
		}
		c.pc++

	case "inc":
		if !inst.op1.direct {
			c.reg[inst.op1.reg]++
		}
		c.pc++

	case "dec":
		if !inst.op1.direct {
			c.reg[inst.op1.reg]--
		}
		c.pc++

	case "jnz":
		if c.lookup(inst.op1) != 0 {
			c.pc += c.lookup(inst.op2)
		} else {
			c.pc++
		}

	case "tgl":
		t := c.pc + c.lookup(inst.op1)
		if t >= 0 && t < len(prog) {
			inst2 := prog[t]
			switch inst2.code {
			case "inc":
				inst2.code = "dec"

			case "dec":
				inst2.code = "inc"

			case "tgl":
				inst2.code = "inc"

			case "jnz":
				inst2.code = "cpy"

			case "cpy":
				inst2.code = "jnz"
			}
		}

		c.pc++

	}

	return c.pc >= 0 && c.pc < len(prog)
}

func solve(p *Problem) {
	prog := program(nil)
	for p.NextLine() {
		f := SplitFields(p.Line())
		inst := &instruction{
			code: f[0],
			op1:  parseOperand(f[1]),
		}
		if len(f) > 2 {
			inst.op2 = parseOperand(f[2])
		}
		prog = append(prog, inst)
	}

	if p.Example() {
		p.PartOne(part(prog, 0))
	} else {
		p.PartOne(part(prog, 7))
		p.PartTwo(part(prog, 12))
	}
}

func part(prog program, a int) int {
	c := cpu{}
	c.reg[0] = a

	p := prog.clone()
	for {
		//fmt.Printf("%d %v\n", c.pc, c.reg)
		if !c.step(p) {
			break
		}
	}
	return c.reg[0]
}
