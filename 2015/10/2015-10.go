package main

import (
	"strconv"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2015
	day  = 10
)

func main() {
	Run(year, day, "", solve)
}

func apply(seq []byte, times int) []byte {
	for n := 0; n < times; n++ {
		var next []byte
		last, run := byte(0), 0
		for _, x := range seq {
			if x == last {
				run++
				continue
			}

			next = appendRun(next, last, run)
			last, run = x, 1
		}
		seq = appendRun(next, last, run)
	}
	return seq
}

func appendRun(s []byte, last byte, run int) []byte {
	if run == 0 {
		return s
	}
	rs := strconv.Itoa(run)
	s = append(s, rs...)
	s = append(s, last)
	return s
}

func solve(p *Problem) {
	p.NextLine()
	seq := p.LineBytes()
	seq = apply(seq, 40)
	p.PartOne(len(seq))
	seq = apply(seq, 10)
	p.PartTwo(len(seq))
}
