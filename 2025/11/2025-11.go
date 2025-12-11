package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 11
	example = `

svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
you: svr

`
)

/*

svr: aaa bbb
aaa: out
bbb: out

svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out

*/

func main() {
	Run(year, day, example, solve)
}

type node struct {
	name string
	id   int
	rank int
	next []string
}

type graph struct {
	nodes map[string]*node
	conn  [][]int
}

func newGraph() *graph {
	return &graph{
		nodes: map[string]*node{},
	}
}

func (g *graph) insert(name string, next []string) {
	n := g.nodes[name]
	if n == nil {
		n = &node{name: name, id: len(g.nodes)}
		g.nodes[name] = n
	}

	for _, n2 := range next {
		n.next = append(n.next, n2)
		if _, ok := g.nodes[n2]; !ok {
			g.nodes[n2] = &node{name: n2, id: len(g.nodes)}
		}
	}
}

func (g *graph) rankNodes() {
	inc := map[string]int{}
	for _, n := range g.nodes {
		for _, n2 := range n.next {
			inc[n2]++
		}
	}

	q := Queue[string]{}
	for n := range g.nodes {
		if inc[n] == 0 {
			q.Push(n)
		}
	}

	rank := 0
	for q.Len() > 0 {
		n := g.nodes[q.Pop()]
		n.rank = rank
		rank++

		for _, n2 := range n.next {
			inc[n2]--
			if inc[n2] == 0 {
				q.Push(n2)
			}
		}
	}
}

func (g *graph) computeConn() {
	n := len(g.nodes)
	conn := make([][]int, n)
	for i := 0; i < n; i++ {
		conn[i] = make([]int, n)
	}

	for _, n1 := range g.nodes {
		for _, id2 := range n1.next {
			conn[n1.id][g.nodes[id2].id] = 1
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				conn[i][j] += conn[i][k] * conn[k][j]
			}
		}
	}

	g.conn = conn
}

func solve(p *Problem) {
	g := newGraph()
	for p.NextLine() {
		f := SplitFieldsDelim(p.Line(), " :")
		g.insert(f[0], f[1:])
	}

	g.rankNodes()

	p.PartOne(g.countPaths("you", "out"))
	//p.PartOne(g.countPathsByMatrix("you", "out"))

	svrFft := g.countPaths("svr", "fft")
	svrDac := g.countPaths("svr", "dac")
	fftDac := g.countPaths("fft", "dac")
	dacFft := g.countPaths("dac", "fft")
	fftOut := g.countPaths("fft", "out")
	dacOut := g.countPaths("dac", "out")
	p.PartTwo(svrFft*fftDac*dacOut + svrDac*dacFft*fftOut)
}

func (g *graph) countPaths(start, stop string) int {
	heap := NewHeap[*node](func(n1 *node, n2 *node) bool { return n1.rank < n2.rank })
	heap.Push(g.nodes[start])

	counts := map[string]int{}
	counts[start] = 1

	for heap.Len() > 0 {
		n := heap.Pop()

		if n.name == stop {
			continue
		}

		c := counts[n.name]
		for _, n2 := range n.next {
			if counts[n2] == 0 {
				counts[n2] = c
				heap.Push(g.nodes[n2])
			} else {
				counts[n2] += c
			}
		}
	}

	return counts[stop]
}

func (g *graph) countPathsByMatrix(start, stop string) int {
	if g.conn == nil {
		g.computeConn()
	}
	return g.conn[g.nodes[start].id][g.nodes[stop].id]
}

//type graph struct {
//	nodes map[string][]string
//	ids   map[string]int
//	conn  [][]int
//}
//
//func (g *graph) allocateId(id string) {
//	if _, ok := g.ids[id]; !ok {
//		g.ids[id] = len(g.ids)
//	}
//}
//
//func countPaths(g *graph, start, stop string) int {
//	queue := Queue[string]{}
//	queue.Push(start)
//
//	count := 0
//	for queue.Len() > 0 {
//		at := queue.Pop()
//		if at == stop {
//			count++
//			continue
//		}
//
//		for _, id := range g.nodes[at] {
//			queue.Push(id)
//		}
//	}
//	return count
//}
//
//type path struct {
//	num  int
//	from map[string]int
//}
//
//func (p *path) add(id string) {
//	if p.from == nil {
//		p.from = make(map[string]int)
//	}
//	p.from[id]++
//}
//
//func (p *path) addAll(p2 *path) {
//	for id := range p2.from {
//		p.add(id)
//	}
//}
//
//func countPaths2(g *graph, start, stop string) int {
//	queue := Queue[string]{}
//	queue.Push(start)
//
//	paths := map[string]*path{}
//	paths[start] = &path{num: 1}
//
//	for queue.Len() > 0 {
//		at := queue.Pop()
//		p := paths[at]
//
//		if at == stop {
//			continue
//		}
//
//		for _, id := range g.nodes[at] {
//			next := paths[id]
//			if next == nil {
//				next = &path{num: p.num}
//				next.addAll(p)
//				next.add(at)
//				paths[id] = next
//				queue.Push(id)
//			} else {
//				next.num += p.num
//				next.addAll(p)
//				next.add(at)
//
//				for _, p2 := range paths {
//					if x := p2.from[id]; x != 0 {
//						p2.num += x * p.num
//						p2.addAll(p)
//						p2.add(at)
//					}
//				}
//			}
//		}
//	}
//
//	p := paths[stop]
//	if p == nil {
//		return 0
//	}
//
//	return p.num
//}
