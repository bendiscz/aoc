package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 18
	example = `

snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d

`
)

//set a 1
//add a 2
//mul a a
//mod a 5
//snd a
//set a 0
//rcv a
//jgz a -1
//set a 1
//jgz a -2

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
	regs    [26]int
	pc      int
	snd     int
	rcv     []int
	blocked bool
}

func newCpu(id int) *cpu {
	c := &cpu{}
	c.regs['p'-'a'] = id
	return c
}

func (c *cpu) run(i inst) {
	switch i.code {
	case "set":
		c.regs[i.op0.reg] = c.lookup(i.op1)
		c.pc++

	case "add":
		c.regs[i.op0.reg] += c.lookup(i.op1)
		c.pc++

	case "mul":
		c.regs[i.op0.reg] *= c.lookup(i.op1)
		c.pc++

	case "mod":
		c.regs[i.op0.reg] %= c.lookup(i.op1)
		c.pc++

	case "jgz":
		if c.lookup(i.op0) > 0 {
			c.pc += c.lookup(i.op1)
		} else {
			c.pc++
		}
	}
}

func (c *cpu) run1(i inst) (int, bool) {
	switch i.code {
	case "snd":
		c.snd = c.lookup(i.op0)
		c.pc++

	case "rcv":
		c.pc++
		if c.regs[i.op0.reg] != 0 {
			return c.snd, true
		}

	default:
		c.run(i)
	}
	return 0, false
}

func (c *cpu) run2(i inst) (int, bool) {
	switch i.code {
	case "snd":
		c.snd++
		c.pc++
		return c.lookup(i.op0), true

	case "rcv":
		if c.blocked {
			break
		}
		if len(c.rcv) == 0 {
			c.blocked = true
			break
		}
		c.regs[i.op0.reg] = c.rcv[0]
		copy(c.rcv, c.rcv[1:])
		c.rcv = c.rcv[:len(c.rcv)-1]
		c.pc++

	default:
		c.run(i)
	}
	return 0, false
}

func (c *cpu) send(x int) {
	c.rcv = append(c.rcv, x)
	c.blocked = false
}

func (c *cpu) lookup(x op) int {
	if x.reg >= 0 {
		return c.regs[x.reg]
	} else {
		return x.val
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

	p.PartOne(part1(prog))
	p.PartTwo(part2(prog))
}

func part1(prog []inst) int {
	c := newCpu(0)
	for c.pc < len(prog) {
		if recovered, ok := c.run1(prog[c.pc]); ok {
			return recovered
		}
	}
	return 0
}

func part2(prog []inst) int {
	c0, c1 := newCpu(0), newCpu(1)
	c := c0
	for !c0.blocked || !c1.blocked {
		out, sent := c.run2(prog[c.pc])

		if c == c0 {
			c = c1
		} else {
			c = c0
		}
		if sent {
			c.send(out)
		}
	}
	return c1.snd
}
