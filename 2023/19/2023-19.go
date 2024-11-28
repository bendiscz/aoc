package main

import (
	"math"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 19
	example = `

px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}

`
)

func main() {
	Run(year, day, example, solve)
}

const (
	cool = iota
	musical
	aero
	shiny
)

const (
	lt = iota
	gt
)

type rule struct {
	id         string
	conditions []condition
	defTarget  string
}

type condition struct {
	param  int
	op     int
	value  int
	target string
}

func (c condition) eval(it item) bool {
	p := it.params[c.param]
	switch c.op {
	case lt:
		return p < c.value
	case gt:
		return p > c.value
	}
	panic("invalid op")
}

type item struct {
	params [4]int
}

func (it item) sum() int {
	return it.params[0] + it.params[1] + it.params[2] + it.params[3]
}

func solve(p *Problem) {
	rules := map[string]*rule{}
	for p.NextLine() && p.Line() != "" {
		f := SplitFieldsDelim(p.Line(), "{},")
		r := rule{id: f[0], defTarget: f[len(f)-1]}
		for i := 1; i < len(f)-1; i++ {
			c := condition{}

			switch f[i][0] {
			case 'x':
				c.param = cool
			case 'm':
				c.param = musical
			case 'a':
				c.param = aero
			case 's':
				c.param = shiny
			}

			switch f[i][1] {
			case '<':
				c.op = lt
			case '>':
				c.op = gt
			}

			j := strings.IndexByte(f[i], ':')
			c.value = ParseInt(f[i][2:j])
			c.target = f[i][j+1:]

			r.conditions = append(r.conditions, c)
		}
		rules[r.id] = &r
	}

	sum1 := 0
	for p.NextLine() {
		f := ParseInts(p.Line())
		it := item{params: [4]int{
			cool:    f[0],
			musical: f[1],
			aero:    f[2],
			shiny:   f[3],
		}}

		if evaluate(it, rules) {
			sum1 += it.sum()
		}
	}
	p.PartOne(sum1)

	it := itemset{params: [4]IntervalSet[int]{
		cool:    NewIntervalSet(1, 4001),
		musical: NewIntervalSet(1, 4001),
		aero:    NewIntervalSet(1, 4001),
		shiny:   NewIntervalSet(1, 4001),
	}}
	r := part2(it, "in", rules)

	sum2 := 0
	for _, i := range r {
		sum2 += i.sum()
	}
	p.PartTwo(sum2)
}

func evaluate(it item, rules map[string]*rule) bool {
	id := "in"
loop:
	for id != "A" && id != "R" {
		r := rules[id]
		for _, c := range r.conditions {
			if c.eval(it) {
				id = c.target
				continue loop
			}
		}
		id = r.defTarget
	}
	return id == "A"
}

type itemset struct {
	params [4]IntervalSet[int]
}

func (it itemset) sum() int {
	sum := 1
	for _, is := range it.params {
		count := 0
		for _, in := range is {
			count += in.X2 - in.X1
		}
		sum *= count
	}
	return sum
}

func (it itemset) apply(c condition) (itemset, itemset) {
	var i1, i2 Interval[int]
	if c.op == lt {
		i1 = Interval[int]{math.MinInt, c.value}
		i2 = Interval[int]{c.value, math.MaxInt}
	} else {
		i1 = Interval[int]{c.value + 1, math.MaxInt}
		i2 = Interval[int]{math.MinInt, c.value + 1}
	}

	it1, it2 := it, it
	it1.params[c.param] = intersect(it1.params[c.param], i1)
	it2.params[c.param] = intersect(it2.params[c.param], i2)
	return it1, it2
}

func intersect(is IntervalSet[int], it Interval[int]) IntervalSet[int] {
	r := IntervalSet[int](nil)
	for _, i := range is {
		x1 := max(i.X1, it.X1)
		x2 := min(i.X2, it.X2)
		if x1 < x2 {
			r = append(r, Interval[int]{x1, x2})
		}
	}
	return r
}

func part2(it itemset, id string, rules map[string]*rule) []itemset {
	if id == "A" {
		return []itemset{it}
	}
	if id == "R" {
		return nil
	}

	r, accepted := rules[id], []itemset(nil)
	for _, c := range r.conditions {
		it1, it2 := it.apply(c)
		accepted = append(accepted, part2(it1, c.target, rules)...)
		it = it2
	}
	accepted = append(accepted, part2(it, r.defTarget, rules)...)

	return accepted
}
