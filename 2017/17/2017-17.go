package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 17
	example = `

3

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	const n1, n2 = 2017, 50_000_000
	buf := make([]int, 1, n1+1)
	buf[0] = 0

	pos := 0
	step := ParseInt(p.ReadLine())

	for i := 1; i <= n1; i++ {
		pos = (pos + step) % len(buf)
		buf = append(buf, 0)
		pos++
		copy(buf[pos+1:], buf[pos:len(buf)-1])
		buf[pos] = i
	}
	p.PartOne(buf[(pos+1)%len(buf)])

	pos = 0
	buf1, length := 0, 1
	for length <= n2 {
		skip := (length - pos) / step
		length += skip
		pos = (pos + max(skip*step, 1)) % length
		pos++
		if pos == 1 {
			buf1 = length
		}
		length++

	}
	p.PartTwo(buf1)
}
