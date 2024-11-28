package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 12
	example = `

cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a

`
)

func main() {
	Run(year, day, example, solve)
}

type state struct {
	p int
	r [4]int
}

func reg(s string) int {
	if len(s) != 1 {
		return -1
	}
	r := s[0] - 'a'
	if r < 0 || r > 3 {
		return -1
	}
	return int(r)
}

func solve(p *Problem) {
	prog := []string(nil)
	for p.NextLine() {
		prog = append(prog, p.Line())
	}

	p.PartOne(exec(prog, 0))
	p.PartTwo(exec(prog, 1))
}

func exec(prog []string, c int) int {
	s := state{}
	s.r[2] = c
	for s.p < len(prog) {
		f := strings.Fields(prog[s.p])
		switch f[0] {
		case "cpy":
			x, y := reg(f[1]), reg(f[2])
			if x < 0 {
				s.r[y] = ParseInt(f[1])
			} else {
				s.r[y] = s.r[x]
			}
			s.p++
		case "inc":
			s.r[reg(f[1])]++
			s.p++
		case "dec":
			s.r[reg(f[1])]--
			s.p++
		case "jnz":
			x := reg(f[1])
			if x < 0 {
				x = ParseInt(f[1])
			} else {
				x = s.r[x]
			}
			if x != 0 {
				s.p += ParseInt(f[2])
			} else {
				s.p++
			}
		}
	}
	return s.r[0]
}
