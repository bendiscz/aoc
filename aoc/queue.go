package aoc

const queueChunk = 1024

type Queue[T any] struct {
	h, t   int
	chunks [][queueChunk]T
}

func (q *Queue[T]) Len() int {
	if len(q.chunks) == 0 {
		return 0
	}
	return queueChunk*(len(q.chunks)-1) - q.h + q.t
}

func (q *Queue[T]) Push(v T) {
	if q.chunks == nil {
		q.chunks = make([][queueChunk]T, 1)
	}

	if q.t == queueChunk {
		q.chunks = append(q.chunks, [queueChunk]T{})
		q.t = 0
	}

	q.chunks[len(q.chunks)-1][q.t] = v
	q.t++
}

func (q *Queue[T]) Pop() T {
	if q.chunks == nil || len(q.chunks) == 1 && q.h == q.t {
		panic("queue underflow")
	}

	v := q.chunks[0][q.h]
	q.h++
	if q.h == queueChunk {
		if len(q.chunks) == 1 {
			q.h, q.t = 0, 0
		} else {
			chunks := make([][queueChunk]T, len(q.chunks)-1)
			copy(chunks, q.chunks[1:])
			q.chunks = chunks
			q.h = 0
		}
	}

	return v
}
