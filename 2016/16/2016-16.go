package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 16
	example = `

10000

`
)

//00 1
//01 0
//10 0
//11 1
//
//0000 11 1
//0001 10 0
//0010 10 0
//0011 11 1
//0100 01 0
//0101 00 1
//0110 00 1
//0111 01 0
//1000 01 0
//1001 00 1
//1010 00 1
//1011 01 0
//1100 11 1
//1101 10 0
//1110 10 0
//1111 11 1

/*

x
x0x
x0x0x1x
x0x0x1x0x0x1x1x
x0x0x1x0x0x1x1x0x0x0x1x1x0x1x1x
x0x0x1x0x0x1x1x0x0x0x1x1x0x1x1x0x0x0x1x0x0x1x1x1x0x0x1x1x0x1x1x
x0x0x1x0x0x1x1x0x0x0x1x1x0x1x1x0x0x0x1x0x0x1x1x1x0x0x1x1x0x1x1x0x0x0x1x0x0x1x1x0x0x0x1x1x0x1x1x1x0x0x1x0x0x1x1x1x0x0x1x1x0x1x1x
*/

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	input, n1, n2 := p.ReadLine(), 272, 35651584
	if p.Example() {
		n1, n2 = 20, 20
	}

	s := input
	for len(s) < n1 {
		s = dragon(s)
	}
	p.PartOne(checksum(s[:n1]))

	for len(s) < n2 {
		s = dragon(s)
	}
	p.PartTwo(checksum(s[:n2]))
}

func dragon(s string) string {
	d := make([]byte, 2*len(s)+1)
	for i := 0; i < len(s); i++ {
		b := s[i]
		d[i] = b
		if b == '0' {
			b = '1'
		} else {
			b = '0'
		}
		d[2*len(s)-i] = b
	}
	d[len(s)] = '0'
	return string(d)
}

func checksum(s string) string {
	d := []byte(s)
	for len(d)%2 == 0 {
		l := len(d) / 2
		for i := 0; i < l; i++ {
			if d[2*i] == d[2*i+1] {
				d[i] = '1'
			} else {
				d[i] = '0'
			}
		}
		d = d[:l]
	}
	return string(d)
}
