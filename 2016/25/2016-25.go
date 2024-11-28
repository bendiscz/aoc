package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 25
	example = ``
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

type cpu struct {
	reg  [4]int
	pc   int
	last int
}

func (c *cpu) lookup(op operand) int {
	if op.direct {
		return op.value
	} else {
		return c.reg[op.reg]
	}
}

func (c *cpu) step(prog program) (bool, bool) {
	if c.pc == len(prog)-1 {
		return false, true
	}

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

	case "out":
		out := c.lookup(inst.op1)
		if out == c.last {
			return false, false
		}
		c.last = out
		c.pc++
	}

	return c.pc >= 0 && c.pc < len(prog), false
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

	for a := 0; ; a++ {
		if part(prog, a) {
			p.PartOne(a)
			break
		}
	}
}

func part(prog program, a int) bool {
	c := cpu{}
	c.reg[0] = a
	c.last = math.MaxInt

	for {
		if ok, found := c.step(prog); !ok {
			return found
		}
	}
}
