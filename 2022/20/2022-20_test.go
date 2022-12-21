package main

import (
	"testing"

	"github.com/bendiscz/aoc/aoc"
)

func Benchmark(b *testing.B) {
	p := aoc.Day(year, day)
	p.Silence()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Reset()
		solve(p)
	}
}
