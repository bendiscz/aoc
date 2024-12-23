package main

import (
	"github.com/bendiscz/aoc/2018/asm"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 16
	example = `

Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]



---

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	ops := [16]Set[string]{}
	for i := 0; i < len(ops); i++ {
		ops[i] = Set[string]{}
		for inst := range asm.FnByName {
			ops[i][inst] = SET
		}
	}

	s1 := 0
	for p.NextLine() && p.Line() != "" {
		f := ParseInts(p.Line())
		r1 := asm.Regs{f[0], f[1], f[2], f[3]}
		p.NextLine()
		code := ParseInts(p.Line())
		p.NextLine()
		f = ParseInts(p.Line())
		r2 := asm.Regs{f[0], f[1], f[2], f[3]}
		p.NextLine()

		s := Set[string]{}
		for inst, fn := range asm.FnByName {
			if r2 == fn(code[1], code[2], code[3], r1) {
				s[inst] = SET
			}
		}
		if len(s) >= 3 {
			s1++
		}
		ops[code[0]] = ops[code[0]].Intersect(s)
	}
	p.PartOne(s1)

	if p.Example() {
		return
	}

	opcodes := map[int]asm.Fn{}
	for len(opcodes) < 16 {
		opcode := -1
		for i, set := range ops {
			if len(set) == 1 {
				opcode = i
				break
			}
		}
		if opcode == -1 {
			panic("ambiguous opcodes")
		}

		var inst string
		for i := range ops[opcode] {
			inst = i
		}

		opcodes[opcode] = asm.FnByName[inst]
		for _, set := range ops {
			delete(set, inst)
		}
	}

	r := asm.Regs{}
	for p.NextLine() {
		s := p.Line()
		if s == "" {
			continue
		}
		code := ParseInts(s)
		r = opcodes[code[0]](code[1], code[2], code[3], r)
	}
	p.PartTwo(r[0])
}
