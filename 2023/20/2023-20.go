package main

import (
	"fmt"
	"time"

	. "github.com/bendiscz/aoc/aoc"
	"golang.org/x/exp/slices"
)

const (
	year    = 2023
	day     = 20
	example = `

broadcaster -> hd
%hd -> dn, mm
%mm -> nv
%nv -> pn, dn
%pn -> dn, jk
%jk -> nr
%nr -> hf
%hf -> qf
%qf -> kf
%kf -> dz, dn
%dz -> zb, dn
%zb -> dn, zf
%zf -> dn
&dn -> jk, qf, gc, hf, hd, nr, mm
&gc -> zr
&zr -> rx

`
)

func main() {
	Run(year, day, example, solve)
}

type signal bool

const (
	low  signal = false
	high        = true
)

func (s signal) String() string {
	if s == high {
		return "high"
	} else {
		return "low"
	}
}

type pulse struct {
	src, dst string
	value    signal
}

func (p pulse) String() string {
	return fmt.Sprintf("%s -%v-> %s", p.src, p.value, p.dst)
}

type module interface {
	id() string
	targets() []string
	connect(source string)
	process(input pulse) []pulse
}

type baseModule struct {
	name string
	tgts []string
}

func newBaseModule(id string, targets []string) baseModule {
	return baseModule{
		name: id,
		tgts: slices.Clone(targets),
	}
}

func (m *baseModule) id() string {
	return m.name
}

func (m *baseModule) targets() []string {
	return m.tgts
}

func (m *baseModule) emit(value signal) []pulse {
	output := make([]pulse, len(m.tgts))
	for i, dst := range m.tgts {
		output[i] = pulse{m.id(), dst, value}
	}
	return output
}

type broadcastModule struct {
	baseModule
	value signal
}

func newBroadcast(targets []string) *broadcastModule {
	return &broadcastModule{
		baseModule: newBaseModule("broadcaster", targets),
	}
}

func (m *broadcastModule) connect(_ string) {
}

func (m *broadcastModule) process(input pulse) []pulse {
	return m.emit(input.value)
}

type flipFlopModule struct {
	baseModule
	value signal
}

func newFlipFlop(id string, targets []string) *flipFlopModule {
	return &flipFlopModule{
		baseModule: newBaseModule(id, targets),
	}
}

func (m *flipFlopModule) connect(_ string) {
}

func (m *flipFlopModule) process(input pulse) []pulse {
	if input.value == high {
		return nil
	}

	m.value = !m.value
	return m.emit(m.value)
}

type conjunctionModule struct {
	baseModule
	states map[string]signal
}

func newConjunction(id string, targets []string) *conjunctionModule {
	return &conjunctionModule{
		baseModule: newBaseModule(id, targets),
		states:     map[string]signal{},
	}
}

func (m *conjunctionModule) connect(source string) {
	m.states[source] = low
}

func (m *conjunctionModule) process(input pulse) []pulse {
	if _, ok := m.states[input.src]; ok {
		m.states[input.src] = input.value
	}

	value := low
	for _, v := range m.states {
		if v == low {
			value = high
			break
		}
	}
	return m.emit(value)
}

const bcast = "broadcaster"

type circuit struct {
	names     []string
	modules   map[string]module
	lowCount  int
	highCount int
	rx        bool
}

func newCircuit() *circuit {
	return &circuit{
		modules: map[string]module{},
	}
}

func (c *circuit) add(m module) {
	c.names = append(c.names, m.id())
	c.modules[m.id()] = m
}

func (c *circuit) connect() {
	for _, src := range c.modules {
		for _, tgt := range src.targets() {
			if dst, ok := c.modules[tgt]; ok {
				dst.connect(src.id())
			}
		}
	}
}

func (c *circuit) pushButton() {
	q := Queue[pulse]{}
	q.Push(pulse{
		src:   "button",
		dst:   bcast,
		value: false,
	})

	for q.Len() > 0 {
		p := q.Pop()
		//fmt.Printf("%v\n", p)
		if p.value == high {
			c.highCount++
		} else {
			c.lowCount++
		}

		if p.dst == "rx" {
			if p.value == low {
				c.rx = true
			}
		} else if m, ok := c.modules[p.dst]; ok {
			out := m.process(p)
			for _, o := range out {
				q.Push(o)
			}
		}
	}
}

func solve(p *Problem) {
	c := newCircuit()
	for p.NextLine() {
		if p.Line() == "" || p.Line()[0] == '#' {
			continue
		}
		f := SplitFields(p.Line())
		switch {
		case f[0] == bcast:
			c.add(newBroadcast(f[2:]))
		case f[0][0] == '%':
			c.add(newFlipFlop(f[0][1:], f[2:]))
		case f[0][0] == '&':
			c.add(newConjunction(f[0][1:], f[2:]))
		}
	}
	c.connect()

	for i := 0; i < 1000; i++ {
		c.pushButton()
	}
	p.PartOne(c.lowCount * c.highCount)

	t1 := time.Now()
	for i := 0; i < 1_000_000; i++ {
		c.pushButton()
	}
	p.Printf("%v", time.Since(t1))

	count := 1
	for _, id := range c.modules[bcast].targets() {
		count = LCM(count, inspect(c, id))
	}
	p.PartTwo(count)
}

func inspect(c *circuit, id string) int {
	value, bit := 0, 1
	m := c.modules[id]
	for {
		if len(m.targets()) == 2 {
			value |= bit
		} else if _, ok := c.modules[m.targets()[0]].(*conjunctionModule); ok {
			return value | bit
		}

		for _, tgt := range m.targets() {
			m = c.modules[tgt]
			if _, ok := m.(*flipFlopModule); ok {
				break
			}
		}

		bit <<= 1
	}
}
