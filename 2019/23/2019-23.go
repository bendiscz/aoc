package main

import (
	"github.com/bendiscz/aoc/2019/intcode"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 23
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

type packet struct {
	src, dst int
	x, y     int
}

type node struct {
	comp *intcode.Computer
	id   int
	in   Queue[packet]
}

func newNode(prog *intcode.Program, id int) *node {
	n := &node{
		comp: prog.Exec(),
		id:   id,
	}
	n.comp.WriteInt(n.id)
	return n
}

func (n *node) run() ([]packet, bool) {
	out := []packet(nil)
	for {
		data := n.comp.ReadInts(3)
		if len(data) != 3 {
			break
		}
		out = append(out, packet{src: n.id, dst: data[0], x: data[1], y: data[2]})
	}

	if n.in.Len() == 0 && n.comp.State() == intcode.Blocked {
		n.comp.WriteInt(-1)
		n.comp.Run()
		return out, n.comp.State() == intcode.Blocked
	}

	for n.in.Len() > 0 {
		p := n.in.Pop()
		n.comp.WriteInts([]int{p.x, p.y})
	}

	return out, false
}

type network struct {
	nodes []*node
	nat   bool
	np    *packet
	lastY int
}

func newNetwork(prog *intcode.Program, nat bool) *network {
	net := &network{
		nodes: make([]*node, 50),
		nat:   nat,
	}
	for i := range net.nodes {
		net.nodes[i] = newNode(prog, i)
	}
	return net
}

func (net *network) run() int {
	for {
		active := false
		for _, n := range net.nodes {
			out, idle := n.run()
			if !idle {
				active = true
			}
			for _, p := range out {
				if p.dst == 255 {
					if !net.nat {
						return p.y
					}

					net.np = &p
					continue
				}
				net.nodes[p.dst].in.Push(p)
				active = true
			}
		}

		if !active {
			if net.np != nil {
				p := *net.np
				net.np = nil

				if p.y == net.lastY {
					return p.y
				}

				net.lastY = p.y
				net.nodes[0].in.Push(p)
			}
		}
	}
}

func solve(p *Problem) {
	prog := intcode.Parse(p)
	net := newNetwork(prog, false)
	p.PartOne(net.run())
	net = newNetwork(prog, true)
	p.PartTwo(net.run())
}
