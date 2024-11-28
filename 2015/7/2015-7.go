package main

import (
	"strconv"
	"unicode"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 7
	example = `

123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i

`
)

func main() {
	Run(year, day, example, solve)
}

type source struct {
	value uint16
	w     *wire
}

type wire struct {
	label string
	value *uint16
	gates []gate
}

func (w *wire) addGate(g gate) {
	w.gates = append(w.gates, g)
}

func (w *wire) send(value uint16) {
	if w.value != nil {
		return
	}
	w.value = &value
	for _, g := range w.gates {
		g.send(w)
	}
}

type gate interface {
	send(w *wire)
}

type idGate struct {
	w *wire
}

func (g idGate) send(w *wire) {
	g.w.send(*w.value)
}

type notGate struct {
	w *wire
}

func (g notGate) send(w *wire) {
	g.w.send(^*w.value)
}

type shiftGate struct {
	dir byte
	l   int
	w   *wire
}

func (g shiftGate) send(w *wire) {
	v := *w.value
	switch g.dir {
	case 'R':
		v >>= g.l
	case 'L':
		v <<= g.l
	}
	g.w.send(v)
}

type opGate struct {
	v  *uint16
	op string
	w  *wire
}

func (g *opGate) send(w *wire) {
	if g.v == nil {
		g.v = w.value
		return
	}

	var v uint16
	switch g.op {
	case "AND":
		v = *g.v & *w.value
	case "OR":
		v = *g.v | *w.value
	}
	g.w.send(v)
}

type circuit struct {
	sources []source
	wires   map[string]*wire
}

func (c *circuit) load(p *Problem) {
	for p.NextLine() {
		if m := p.Parse(`^(\d+) -> ([a-z]+)$`); m != nil {
			c.addSource(m)
		} else if m = p.Parse(`^([a-z]+) -> ([a-z]+)$`); m != nil {
			c.addId(m)
		} else if m = p.Parse(`^NOT ([a-z]+) -> ([a-z]+)$`); m != nil {
			c.addNot(m)
		} else if m = p.Parse(`^([a-z]+) ([LR])SHIFT (\d+) -> ([a-z]+)$`); m != nil {
			c.addShift(m)
		} else if m = p.Parse(`^([a-z]+|\d+) (AND|OR) ([a-z]+) -> ([a-z]+)$`); m != nil {
			c.addOp(m)
		}
	}
}

func (c *circuit) run() {
	for _, src := range c.sources {
		src.w.send(src.value)
	}
}

func (c *circuit) getValue(label string) string {
	w := c.wires[label]
	if w == nil || w.value == nil {
		return "none"
	}
	return strconv.Itoa(int(*w.value))
}

func (c *circuit) addSource(m []string) {
	c.sources = append(c.sources, source{
		value: uint16(ParseInt(m[1])),
		w:     c.getWire(m[2]),
	})
}

func (c *circuit) getWire(label string) *wire {
	if w, ok := c.wires[label]; ok {
		return w
	}

	if c.wires == nil {
		c.wires = map[string]*wire{}
	}

	w := &wire{
		label: label,
	}
	c.wires[label] = w

	return w
}

func (c *circuit) addId(m []string) {
	c.getWire(m[1]).addGate(idGate{
		w: c.getWire(m[2]),
	})
}

func (c *circuit) addNot(m []string) {
	c.getWire(m[1]).addGate(notGate{
		w: c.getWire(m[2]),
	})
}

func (c *circuit) addShift(m []string) {
	c.getWire(m[1]).addGate(shiftGate{
		dir: m[2][0],
		l:   ParseInt(m[3]),
		w:   c.getWire(m[4]),
	})
}

func (c *circuit) addOp(m []string) {
	g := &opGate{
		op: m[2],
		w:  c.getWire(m[4]),
	}

	if unicode.IsDigit(rune(m[1][0])) {
		v := uint16(ParseInt(m[1]))
		g.v = &v
	} else {
		c.getWire(m[1]).addGate(g)
	}
	c.getWire(m[3]).addGate(g)
}

func solve(p *Problem) {
	c := circuit{}
	c.load(p)
	c.run()
	a := c.getValue("a")
	p.PartOne(a)

	p.Reset()
	c = circuit{}
	c.load(p)
	for i := 0; i < len(c.sources); i++ {
		if c.sources[i].w.label == "b" {
			c.sources[i].value = uint16(ParseInt(a))
		}
	}
	c.run()
	p.PartTwo(c.getValue("a"))
}
