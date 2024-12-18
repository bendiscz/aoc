package intcode

import (
	"fmt"
	"github.com/bendiscz/aoc/aoc"
	"maps"
	"math"
	"slices"
	"strings"
)

type Program struct {
	Code []int
}

func Parse(p *aoc.Problem) *Program {
	p.NextLine()
	return &Program{Code: aoc.ParseInts(p.Line())}
}

func (p *Program) Exec() *Computer {
	return newComputer(p)
}

type State string

const (
	Ready   State = "ready"
	Blocked State = "blocked"
	Exited  State = "exited"
)

type Computer struct {
	state State
	mem   map[int]int
	pc    int
	bp    int
	in    *buffer
}

func newComputer(prog *Program) *Computer {
	return &Computer{
		state: Ready,
		mem:   maps.Collect(slices.All(prog.Code)),
		in:    newBuffer(),
	}
}

func (c *Computer) State() State {
	return c.state
}

func (c *Computer) Run() {
	c.run(func(int) bool { return false })
}

func (c *Computer) ReadWrite(input ...int) []int {
	for _, v := range input {
		c.in.write(v)
	}
	return c.ReadAll()
}

func (c *Computer) ReadInt() (int, bool) {
	out, ok := 0, false
	c.run(func(value int) bool {
		if ok {
			return false
		}
		out, ok = value, true
		return true
	})
	return out, ok
}

func (c *Computer) ReadInts(n int) []int {
	out := []int(nil)
	c.run(func(value int) bool {
		if n <= 0 {
			return false
		}
		out = append(out, value)
		n--
		return true
	})
	return out
}

func (c *Computer) ReadAll() []int {
	return c.ReadInts(math.MaxInt)
}

func (c *Computer) ReadLine() (string, bool) {
	line, ok := &strings.Builder{}, false
	c.run(func(value int) bool {
		switch {
		case ok || value < 0 || value > 127:
			return false
		case value == '\n':
			ok = true
			return true
		default:
			line.WriteByte(byte(value))
			return true
		}
	})
	return line.String(), ok
}

func (c *Computer) WriteInt(value int) {
	c.in.write(value)
}

func (c *Computer) WriteInts(values []int) {
	for _, v := range values {
		c.in.write(v)
	}
}

func (c *Computer) WriteLine(line string) {
	for i := 0; i < len(line); i++ {
		c.in.write(int(line[i]))
	}
	c.in.write('\n')
}

func (c *Computer) Peek(addr int) int {
	return c.mem[addr]
}

func (c *Computer) Poke(addr, value int) {
	c.mem[addr] = value
}

func (c *Computer) preFetch(pos int) (mode, v int) {
	ins := c.mem[c.pc] / 100
	for i := 1; i < pos; i++ {
		ins /= 10
	}
	return ins % 10, c.mem[c.pc+pos]
}

func (c *Computer) fetch(pos int) int {
	mode, v := c.preFetch(pos)
	switch mode {
	case 0:
		return c.mem[v]
	case 1:
		return v
	case 2:
		return c.mem[v+c.bp]
	default:
		panic("invalid fetch mode")
	}
}

func (c *Computer) store(pos, value int) {
	mode, v := c.preFetch(pos)
	switch mode {
	case 0:
		c.mem[v] = value
	case 2:
		c.mem[v+c.bp] = value
	default:
		panic("invalid store mode")
	}
}

func (c *Computer) run(output func(int) bool) {
	for {
		switch c.mem[c.pc] % 100 {
		case 1:
			c.store(3, c.fetch(1)+c.fetch(2))
			c.pc += 4
		case 2:
			c.store(3, c.fetch(1)*c.fetch(2))
			c.pc += 4
		case 3:
			x, ok := c.in.read()
			if !ok {
				c.state = Blocked
				return
			}
			c.store(1, x)
			c.pc += 2
		case 4:
			if !output(c.fetch(1)) {
				c.state = Ready
				return
			}
			c.pc += 2
		case 5:
			if c.fetch(1) != 0 {
				c.pc = c.fetch(2)
			} else {
				c.pc += 3
			}
		case 6:
			if c.fetch(1) == 0 {
				c.pc = c.fetch(2)
			} else {
				c.pc += 3
			}
		case 7:
			if c.fetch(1) < c.fetch(2) {
				c.store(3, 1)
			} else {
				c.store(3, 0)
			}
			c.pc += 4
		case 8:
			if c.fetch(1) == c.fetch(2) {
				c.store(3, 1)
			} else {
				c.store(3, 0)
			}
			c.pc += 4
		case 9:
			c.bp += c.fetch(1)
			c.pc += 2
		case 99:
			c.state = Exited
			return
		default:
			panic(fmt.Sprintf("invalid opcode: %d", c.mem[c.pc]))
		}
	}
}
