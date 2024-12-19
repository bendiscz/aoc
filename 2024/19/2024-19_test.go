package main

import (
	"github.com/bendiscz/aoc/aoc"
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	p := aoc.Day(year, day)
	p.Silence()
	for i := 0; i < b.N; i++ {
		p.Reset()
		solve(p)
	}
}
