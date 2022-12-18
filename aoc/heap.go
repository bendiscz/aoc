package aoc

import "container/heap"

type Heap[T any] struct {
	data []T
	less func(x1, x2 T) bool
}

func (h *Heap[T]) impl() heapImpl[T] { return heapImpl[T]{Heap: h} }

type heapImpl[T any] struct {
	*Heap[T]
}

func (h heapImpl[T]) Len() int {
	return len(h.Heap.data)
}

func (h heapImpl[T]) Less(i, j int) bool {
	return h.Heap.less(h.Heap.data[i], h.Heap.data[j])
}

func (h heapImpl[T]) Swap(i, j int) {
	h.Heap.data[i], h.Heap.data[j] = h.Heap.data[j], h.Heap.data[i]
}

func (h heapImpl[T]) Push(x any) {
	h.Heap.data = append(h.Heap.data, x.(T))
}

func (h heapImpl[T]) Pop() any {
	l := len(h.Heap.data) - 1
	x := h.Heap.data[l]
	h.Heap.data = h.Heap.data[:l]
	return x
}

func NewHeap[T any](less func(T, T) bool) *Heap[T] {
	return &Heap[T]{
		less: less,
	}
}

func (h *Heap[T]) Clear() {
	h.data = nil
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Push(x T) {
	heap.Push(h.impl(), x)
}

func (h *Heap[T]) Pop() T {
	return heap.Pop(h.impl()).(T)
}

func (h *Heap[T]) Top() T {
	return h.data[0]
}

func (h *Heap[T]) Fix() {
	heap.Init(h.impl())
}
