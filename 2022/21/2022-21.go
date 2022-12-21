package main

import (
	"fmt"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 21
	example = `

root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32

`
)

func main() {
	Run(year, day, example, solve)
}

type monkey struct {
	monkeys map[string]*monkey
	name    string
	yelled  bool
	op      string
	m1, m2  string
	v       int
	a, b    float64
}

func parseMonkey(monkeys map[string]*monkey, s string) {
	var m *monkey
	if s[6] >= '0' && s[6] <= '9' {
		name, value := "", 0
		_, _ = fmt.Sscanf(s, "%s %d", &name, &value)
		name = name[:4]

		a, b := 0, value
		if name == "humn" {
			a, b = 1, 0
		}

		m = &monkey{
			monkeys: monkeys,
			name:    name,
			yelled:  true,
			a:       float64(a),
			b:       float64(b),
			v:       value,
		}
	} else {
		name, op, m1, m2 := "", "", "", ""
		_, _ = fmt.Sscanf(s, "%s %s %s %s", &name, &m1, &op, &m2)
		name = name[:4]

		m = &monkey{
			monkeys: monkeys,
			name:    name,
			op:      op,
			m1:      m1,
			m2:      m2,
		}
	}
	monkeys[m.name] = m
}

func (m *monkey) eval() (int, float64, float64) {
	if m.yelled {
		return m.v, m.a, m.b
	}

	m.yelled = true
	v1, a1, b1 := m.monkeys[m.m1].eval()
	v2, a2, b2 := m.monkeys[m.m2].eval()
	switch m.op {
	case "+":
		m.v = v1 + v2
		m.a = a1 + a2
		m.b = b1 + b2
	case "-":
		m.v = v1 - v2
		m.a = a1 - a2
		m.b = b1 - b2
	case "*":
		m.v = v1 * v2
		m.a = a1*b2 + a2*b1
		m.b = b1 * b2
	case "/":
		m.v = v1 / v2
		m.a = a1/b2 + a2/b1
		m.b = b1 / b2
	}
	return m.v, m.a, m.b
}

func solve(p *Problem) {
	monkeys := map[string]*monkey{}
	for p.NextLine() {
		parseMonkey(monkeys, p.Line())
	}

	root := monkeys["root"]
	root.eval()
	p.PartOne(root.v)

	m1, m2 := monkeys[root.m1], monkeys[root.m2]
	p.PartTwo(int((m2.b - m1.b) / (m1.a - m2.a)))
}
