package main

import (
	"math"
	"strconv"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 18
	example = `

[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]

`
)

func main() {
	Run(year, day, example, solve)
}

type node struct {
	term  bool
	value int
	l, r  *node
}

func parse(s string, p int) (*node, int) {
	n := &node{}
	ch := s[p]
	if ch == '[' {
		p++
		n.l, p = parse(s, p)
		p++
		n.r, p = parse(s, p)
		p++
	} else {
		n.term = true
		n.value = int(ch - '0')
		p++
	}
	return n, p
}

func (n *node) String() string {
	var sb strings.Builder
	n.write(&sb)
	return sb.String()
}

func (n *node) write(sb *strings.Builder) {
	if n.term {
		sb.WriteString(strconv.Itoa(n.value))
	} else {
		sb.WriteRune('[')
		n.l.write(sb)
		sb.WriteRune(',')
		n.r.write(sb)
		sb.WriteRune(']')
	}
}

func (n *node) magnitude() int {
	if n.term {
		return n.value
	}
	return 3*n.l.magnitude() + 2*n.r.magnitude()
}

func (n *node) add(m *node) *node {
	return &node{
		l: n,
		r: m,
	}
}

func (n *node) reduce() {
	for {
		if _, _, ok := n.explode(0); ok {
			continue
		}
		if n.split() {
			continue
		}
		break
	}
}

func (n *node) explode(depth int) (*int, *int, bool) {
	if n.term {
		return nil, nil, false
	}

	if depth == 4 {
		l, r := n.l.value, n.r.value
		n.term = true
		n.value, n.l, n.r = 0, nil, nil
		return &l, &r, true
	}

	if l, r, ok := n.l.explode(depth + 1); ok {
		if r != nil {
			n.r.addRight(*r)
			r = nil
		}
		return l, r, true
	}

	if l, r, ok := n.r.explode(depth + 1); ok {
		if l != nil {
			n.l.addLeft(*l)
			l = nil
		}
		return l, r, true
	}

	return nil, nil, false
}

func (n *node) addLeft(x int) {
	if n.term {
		n.value += x
	} else {
		n.r.addLeft(x)
	}
}

func (n *node) addRight(x int) {
	if n.term {
		n.value += x
	} else {
		n.l.addRight(x)
	}
}

func (n *node) split() bool {
	if !n.term {
		return n.l.split() || n.r.split()
	}

	if n.value < 10 {
		return false
	}

	n.l = &node{
		term:  true,
		value: n.value / 2,
	}
	n.r = &node{
		term:  true,
		value: (n.value + 1) / 2,
	}
	n.term, n.value = false, 0
	return true
}

func (n *node) copy() *node {
	c := *n
	if !c.term {
		c.l = c.l.copy()
		c.r = c.r.copy()
	}
	return &c
}

func solve(p *Problem) {
	var sum *node
	var nums []*node
	for p.NextLine() {
		n, _ := parse(p.Line(), 0)
		nums = append(nums, n.copy())
		if sum == nil {
			sum = n
		} else {
			sum = sum.add(n)
		}
		sum.reduce()
	}

	p.PartOne(sum.magnitude())

	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			n := nums[i].copy().add(nums[j].copy())
			n.reduce()
			m := n.magnitude()
			if m > max {
				max = m
			}
		}
	}

	p.PartTwo(max)
}
