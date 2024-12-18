package intcode

import "github.com/bendiscz/aoc/aoc"

const bufferSize = 1024

type buffer struct {
	data []int
	h, t int
}

func newBuffer() *buffer {
	return &buffer{
		data: make([]int, bufferSize),
	}
}

func (b *buffer) read() (int, bool) {
	if b.t == b.h {
		return 0, false
	}

	v := b.data[b.t]
	b.t = (b.t + 1) % len(b.data)
	return v, true
}

func (b *buffer) write(v int) {
	nh := (b.h + 1) % len(b.data)
	if nh == b.t {
		b.grow()
		nh = (b.h + 1) % len(b.data)
	}

	b.data[b.h] = v
	b.h = nh
}

func (b *buffer) len() int {
	return aoc.Mod(b.h-b.t, len(b.data))
}

func (b *buffer) grow() {
	l := len(b.data)
	data := make([]int, 2*l)
	if b.h >= b.t {
		copy(data[b.t:], b.data[b.t:b.h])
	} else {
		copy(data[b.t:], b.data[b.t:])
		copy(data[l:], b.data[:b.h])
		b.h += l
	}
	b.data = data
}
