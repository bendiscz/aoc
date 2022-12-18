package main

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 11
	example = `

Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1

`
)

func main() {
	Run(year, day, example, solve)
}

type game struct {
	monkeys []*monkey
}

type monkey struct {
	g      *game
	id     int
	items  []int
	op     func(int) int
	div    int
	mt, mf int
	count  int
}

func parse(p *Problem) *game {
	g := &game{}
	for _, s := range strings.Split(p.ReadAll(), "\n\n") {
		m, items := &monkey{g: g}, ""
		op1, op2 := "", ""
		_, _ = fmt.Sscanf(strings.Replace(s, ", ", ",", -1), `Monkey %d:
  Starting items: %s
  Operation: new = old %s %s
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`, &m.id, &items, &op1, &op2, &m.div, &m.mt, &m.mf)
		m.items = ParseInts(items)
		switch {
		case op1 == "*" && op2 == "old":
			m.op = func(wl int) int { return wl * wl }
		case op1 == "+":
			x := ParseInt(op2)
			m.op = func(wl int) int { return wl + x }
		case op1 == "*":
			x := ParseInt(op2)
			m.op = func(wl int) int { return wl * x }
		}
		g.monkeys = append(g.monkeys, m)
	}
	return g
}

func (m *monkey) inspectAndThrow(relax func(int) int) {
	for _, wl := range m.items {
		m.count++
		wl = m.op(wl)
		wl = relax(wl)

		var t *monkey
		if wl%m.div == 0 {
			t = m.g.monkeys[m.mt]
		} else {
			t = m.g.monkeys[m.mf]
		}
		t.items = append(t.items, wl)
	}
	m.items = nil
}

func solveGame(p *Problem, n int, part int) int {
	p.Reset()
	g := parse(p)

	var relax func(int) int
	if part == 1 {
		relax = func(wl int) int { return wl / 3 }
	} else {
		mod := 1
		for _, m := range g.monkeys {
			mod *= m.div
		}
		relax = func(wl int) int { return wl % mod }
	}

	for i := 0; i < n; i++ {
		for _, m := range g.monkeys {
			m.inspectAndThrow(relax)
		}
	}

	sort.Slice(g.monkeys, func(i, j int) bool { return g.monkeys[i].count > g.monkeys[j].count })
	return g.monkeys[0].count * g.monkeys[1].count
}

func solve(p *Problem) {
	p.PartOne(solveGame(p, 20, 1))
	p.PartTwo(solveGame(p, 10000, 2))
}
