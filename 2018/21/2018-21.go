package main

import (
	"github.com/bendiscz/aoc/2018/asm"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 21
	example = ``
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

	r := asm.Regs{}
	for r[ip] != 28 {
		r = prog[r[ip]].Run(r)
		r[ip]++
	}
	p.PartOne(r[4])

	//p.PartTwo(searchLoop())

	seen, s2 := Set[int]{}, 0
	r4 := prog[5].C
	r5 := prog[26].C

	r = asm.Regs{}
	for {
		if r[ip] == 17 {
			// accelerate /= 256
			r[r5] /= 256
			r[ip] = 8
			continue
		}

		if r[ip] == 28 {
			if seen.Contains(r[r4]) {
				break
			}
			seen[r[r4]] = SET
			s2 = r[r4]
		}
		r = prog[r[ip]].Run(r)
		r[ip]++
	}
	p.PartTwo(s2)
}

//func searchLoop() int {
//	var r4, r5, r1 int
//	last := 0
//	seen := Set[int]{}
//
//	r4 = 0
//	for {
//		r5 = r4 | 0x10000
//		r4 = 1765573
//
//		for {
//			r1 = r5 & 0xff
//			r4 = ((r4 + r1) * 65899) & 0xffffff
//
//			if r5 < 256 {
//				break
//			}
//			r5 /= 256
//		}
//
//		if seen.Contains(r4) {
//			fmt.Printf("%d\n", len(seen))
//			break
//		}
//		seen[r4] = SET
//		last = r4
//	}
//	return last
//}

/*

r4 = 0
for {
	r5 = r4 | 0x10000
	r4 = 1765573

	for {
		r1 = r5 & 0xff

		r4 = ((r4 + r1) * 65899) & 0xffffff

		if r5 < 256 {
			break
		}

		r5 /= 256
	}

	if r4 == r0 {
		break
	}
}

*/
