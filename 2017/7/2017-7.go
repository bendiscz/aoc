package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 7
	example = `

pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)

`
)

func main() {
	Run(year, day, example, solve)
}

type node struct {
	id       string
	value    int
	parent   *node
	children []*node
}

type tree map[string]*node

func (t tree) node(id string) *node {
	if n, ok := t[id]; ok {
		return n
	}
	n := &node{id: id}
	t[id] = n
	return n
}

func solve(p *Problem) {
	t := tree{}
	for p.NextLine() {
		v := SplitFieldsDelim(p.Line(), " ()->,")
		n := t.node(v[0])
		n.value = ParseInt(v[1])
		for _, id := range v[2:] {
			child := t.node(id)
			child.parent = n
			n.children = append(n.children, child)
		}
	}

	var root *node
	for _, n := range t {
		if n.parent == nil {
			root = n
			break
		}
	}

	p.PartOne(root.id)
	if w, found := root.balance(); found {
		p.PartTwo(w)
	}
}

func (n *node) balance() (int, bool) {
	total, weights := n.value, make([]int, len(n.children))
	for i, child := range n.children {
		w, found := child.balance()
		if found {
			return w, true
		}
		weights[i] = w
		total += weights[i]
	}

	if len(weights) > 2 {
		for i := 0; i < len(weights); i++ {
			w0, w1, w2 := weights[i], weights[(i+1)%len(weights)], weights[(i+2)%len(weights)]
			if w0 != w1 && w0 != w2 {
				return n.children[i].value - w0 + w1, true
			}
		}
	}

	return total, false
}
