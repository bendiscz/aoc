package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"strconv"
	"strings"
)

const (
	year    = 2024
	day     = 17
	example = `

Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	var a, b, c int
	p.NextLine()
	p.Scanf("Register A: %d", &a)
	p.NextLine()
	p.Scanf("Register B: %d", &b)
	p.NextLine()
	p.Scanf("Register C: %d", &c)
	p.NextLine()
	p.NextLine()
	prog := ParseInts(p.Line())

	s1 := strings.Builder{}
	for i, x := range eval(prog, a, b, c) {
		if i > 0 {
			s1.WriteByte(',')
		}
		s1.WriteString(strconv.Itoa(x))
	}
	p.PartOne(s1.String())

	if p.Example() {
		return
	}

	p.PartTwo(solveQuine(prog))
}

func solveQuine(prog []int) int {
	// assumptions:
	// - the program ends with jnz 0
	// - the program computes one output number from lower 11 bits of A
	//   ((A & 7) XOR ((A >> 0..7) & 7)
	// - before looping, the program shifts A, three bits to right

	// seed 1st 11 bits
	for a0 := 0; a0 < 2048; a0++ {
		// the program ends with 0
		if eval(body(prog), a0, 0, 0)[0] != 0 {
			continue
		}
		if a, ok := searchQuine(prog, len(prog)-2, a0); ok {
			return a
		}
	}
	return -1
}

func searchQuine(prog []int, pc, a0 int) (int, bool) {
	if pc < 0 {
		return a0, true
	}

	a0 <<= 3
	// try all next 3-bit variants
	for x := 0; x < 8; x++ {
		a := a0 | x
		out := eval(body(prog), a, 0, 0)
		if out[0] != prog[pc] {
			continue
		}
		if r, ok := searchQuine(prog, pc-1, a); ok {
			return r, true
		}
	}
	return 0, false
}

func body(prog []int) []int {
	// cut the loop
	return prog[:len(prog)-2]
}

func eval(prog []int, a, b, c int) []int {
	out := []int(nil)
	pc := 0
	for pc < len(prog) {
		in, op := prog[pc], prog[pc+1]

		switch in {
		case 0:
			a >>= combo(op, a, b, c)
			pc += 2

		case 1:
			b ^= op
			pc += 2

		case 2:
			b = combo(op, a, b, c) & 7
			pc += 2

		case 3:
			if a == 0 {
				pc += 2
			} else {
				pc = op
			}
		case 4:
			b = b ^ c
			pc += 2

		case 5:
			out = append(out, combo(op, a, b, c)&7)
			pc += 2

		case 6:
			b = a >> combo(op, a, b, c)
			pc += 2

		case 7:
			c = a >> combo(op, a, b, c)
			pc += 2
		}
	}
	return out
}

func combo(x, a, b, c int) int {
	switch x {
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		return x
	}
}

// bst a
// bxl 7
// cdv b
// adv 3
// bxl 7
// bxc
// out b
// jnz 0

// b = a % 8
// b = b ^ 7
// c = a >> b
// a = a >> 3
// b = b ^ 7
// b = b ^ c
// out b
// jnz a 0

// b = a & 7
// b = b ^ 7
// c = a >> b (& 7)
// b = b ^ 7
// b = b ^ c
// b = b & 7
// out b
// a = a >> 3
// jnz a 0

//func solve3(p *Problem, prog []int) {
//	a := 5
//	b := 0
//	c := 0
//
//	for a = 0; a < 2048; a++ {
//		b = a & 7
//		c = (a >> (b ^ 7)) & 7
//		b = b ^ c
//		//p.Printf("%d", b)
//		if b == 0 {
//			if a0, ok := solve4(prog, a, len(prog)-2); ok {
//				p.PartTwo(a0)
//				return
//			}
//		}
//	}
//}
//
//func solve4(prog []int, a0, pc int) (int, bool) {
//	if pc < 0 {
//		return a0, true
//	}
//
//	a0 <<= 3
//	for x := 0; x < 8; x++ {
//		a := a0 | x
//		b := 0
//		c := 0
//
//		b = a & 7
//		c = (a >> (b ^ 7)) & 7
//		b = b ^ c
//
//		if b != prog[pc] {
//			continue
//		}
//
//		if y, ok := solve4(prog, a, pc-1); ok {
//			return y, true
//		}
//	}
//
//	return 0, false
//}
