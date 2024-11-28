package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 23
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

type op struct {
	reg int
	val int
}

func parseOp(s string) op {
	r := int(s[0] - 'a')
	if r >= 0 && r <= 26 {
		return op{reg: r}
	} else {
		return op{reg: -1, val: ParseInt(s)}
	}
}

type inst struct {
	code string
	op0  op
	op1  op
}

type cpu struct {
	regs [26]int
	pc   int
	cnt  int
}

func (c *cpu) reg(x byte) int {
	return c.regs[x-'a']
}

func (c *cpu) lookup(x op) int {
	if x.reg >= 0 {
		return c.regs[x.reg]
	} else {
		return x.val
	}
}

func (c *cpu) run(i inst) {
	switch i.code {
	case "set":
		c.regs[i.op0.reg] = c.lookup(i.op1)
		c.pc++

	case "sub":
		c.regs[i.op0.reg] -= c.lookup(i.op1)
		c.pc++

	case "mul":
		c.regs[i.op0.reg] *= c.lookup(i.op1)
		c.cnt++
		c.pc++

	case "jnz":
		if c.lookup(i.op0) != 0 {
			c.pc += c.lookup(i.op1)
		} else {
			c.pc++
		}
	}
}

func solve(p *Problem) {
	prog := []inst(nil)
	for p.NextLine() {
		f := SplitFields(p.Line())
		i := inst{
			code: f[0],
			op0:  parseOp(f[1]),
		}
		if len(f) == 3 {
			i.op1 = parseOp(f[2])
		}
		prog = append(prog, i)
	}

	c := &cpu{}
	for c.pc < len(prog) {
		c.run(prog[c.pc])
	}
	p.PartOne(c.cnt)

	c = &cpu{}
	c.regs[0] = 1
	for c.reg('f') == 0 {
		c.run(prog[c.pc])
	}
	p.PartTwo(program(c.reg('b'), c.reg('c')))
}

func program(b, c int) int {
	h := 0
	for x := b; x <= c; x += 17 {
		q := int(math.Sqrt(float64(x)))
		for i := 2; i <= q; i++ {
			if x%i == 0 {
				h++
				break
			}
		}
	}
	return h
}
