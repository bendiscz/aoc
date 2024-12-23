package main

import (
	"github.com/bendiscz/aoc/2018/asm"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 19
	example = `

#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	ip := ParseInts(p.ReadLine())[0]
	prog := []asm.Inst(nil)
	for p.NextLine() {
		f := SplitFields(p.Line())
		prog = append(prog, asm.Inst{Name: f[0], A: ParseInt(f[1]), B: ParseInt(f[2]), C: ParseInt(f[3])})
	}

	p.PartOne(runWithShortcut(prog, 0, ip))
	if p.Example() {
		return
	}
	p.PartTwo(runWithShortcut(prog, 1, ip))
}

func runWithShortcut(prog []asm.Inst, r0, ip int) int {
	r := asm.Regs{r0}
	for r[ip] != 1 {
		inst := prog[r[ip]]
		r = inst.Run(r)
		r[ip]++
	}

	n := prog[4].A
	if n == prog[4].C {
		n = prog[4].B
	}

	return sumDivisors(r[n])
}

func sumDivisors(n int) int {
	s := n + 1
	for x := 2; x < n; x++ {
		if n%x == 0 {
			s += x
		}
	}
	return s
}

/*

r1 = 892
if r0 != 0 {
	r1 = 10551292
}

r0 = 0
for r5 := 1; r5 <= r1; r5++ {
	for r3 := 1; r3 <= r1; r3++ {
		if r5 * r3 == r1 {
			r0 += r5
		}
	}
}

-----

r1 = 892
if r0 != 0 {
	r4 = 10550400
	r1 = 10551292
	r0 = 0
}

r5 = 1

0:
r3 = 1

1:
r4 = r5 * r3
if r4 == r1 {
	r0 += r5
}

r3 += 1
if r3 <= r1 {
	goto 1
}

r5 += 1
if r5 <= r1 {
	goto 0
}

END

*/
