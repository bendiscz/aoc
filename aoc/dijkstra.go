package aoc

import (
	"container/heap"
	"log"
)

type dijkstraNode[V Vertex] struct {
	v V
	p *dijkstraNode[V]
	d int
	w int
	i int
}

type dijkstraQueue[V Vertex] struct {
	nodes []*dijkstraNode[V]
}

func (q *dijkstraQueue[V]) Len() int {
	return len(q.nodes)
}

func (q *dijkstraQueue[V]) Less(i, j int) bool {
	return q.nodes[i].w < q.nodes[j].w
}

func (q *dijkstraQueue[V]) Swap(i, j int) {
	n1, n2 := q.nodes[i], q.nodes[j]
	n1.i, n2.i = j, i
	q.nodes[i], q.nodes[j] = n2, n1
}

func (q *dijkstraQueue[V]) Push(x any) {
	n := x.(*dijkstraNode[V])
	n.i = len(q.nodes)
	q.nodes = append(q.nodes, n)
}

func (q *dijkstraQueue[V]) Pop() any {
	l := len(q.nodes) - 1
	n := q.nodes[l]
	n.i = 0
	q.nodes = q.nodes[:l]
	return n
}

type Path[V Vertex] struct {
	Cost  int
	Steps []V
}

func ShortestPath[V Vertex](start, finish V) Path[V] {
	paths := ShortestPaths[V](start, func(v V) (bool, bool) {
		term := v.Key() == finish.Key()
		return term, term
	})
	if len(paths) == 0 {
		return Path[V]{}
	}
	return paths[0]
}

func ShortestPaths[V Vertex](start V, check func(V) (term, stop bool)) []Path[V] {
	paths := []Path[V](nil)
	startNode := &dijkstraNode[V]{start, nil, 0, 0, 0}

	nodes := map[any]*dijkstraNode[V]{}
	nodes[start.Key()] = startNode

	queue := &dijkstraQueue[V]{}
	heap.Push(queue, startNode)

	for queue.Len() > 0 {
		n := heap.Pop(queue).(*dijkstraNode[V])
		term, stop := check(n.v)
		if term {
			steps := make([]V, n.d+1)
			for p := n; p != nil; p = p.p {
				steps[p.d] = p.v
			}
			paths = append(paths, Path[V]{Cost: n.w, Steps: steps})
		}
		if stop {
			break
		}

		for _, e := range n.v.Edges() {
			if e.W < 0 {
				log.Panicf("edge from %v to %v has negative weight: %d", n.v.Key(), e.V.Key(), e.W)
			}

			n2 := nodes[e.V.Key()]
			if n2 == nil {
				n2 = &dijkstraNode[V]{e.V.(V), n, n.d + 1, n.w + e.W, 0}
				nodes[e.V.Key()] = n2
				heap.Push(queue, n2)
			} else {
				w2 := n.w + e.W
				if w2 < n2.w {
					n2.p = n
					n2.d = n.d + 1
					n2.w = w2
					heap.Fix(queue, n2.i)
				}
			}

		}
	}

	return paths
}
