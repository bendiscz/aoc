package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 3
	example = `

987654321111111
811111111111119
234234234234278
818181911112111

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	//s1, s2 := 0, big.NewInt(0)
	s1, s2 := 0, 0

	for p.NextLine() {
		l := p.Line()
		b := []byte(l)

		x1 := 0
		for i := 0; i < len(b); i++ {
			for j := i + 1; j < len(b); j++ {
				jolt := int(b[i]-'0')*10 + int(b[j]-'0')
				x1 = max(x1, jolt)
			}
		}
		s1 += x1

		for i := 0; i < len(b); i++ {
			b[i] = b[i] - '0'
		}

		//x2 := big.NewInt(0)
		//for i := 0; i < len(b); i++ {
		//	for j := i + 1; j < len(b); j++ {
		//		for k := j + 1; k < len(b); k++ {
		//			jolt := big.NewInt(0)
		//			for n := 0; n < len(b); n++ {
		//				if n == i || n == j || n == k {
		//					continue
		//				}
		//				jolt = jolt.Add(jolt.Mul(jolt, big.NewInt(10)), big.NewInt(int64(b[n]-'0')))
		//			}
		//			if jolt.Cmp(x2) > 0 {
		//				x2 = jolt
		//			}
		//			//x2 = max(x2, jolt)
		//		}
		//	}
		//}
		//s2 = s2.Add(s2, x2)

		j2, i2 := 0, 0
		for k := 11; k >= 0; k-- {
			d := 0
			for i := i2; i < len(b)-k; i++ {
				if int(b[i]) > d {
					d = int(b[i])
					i2 = i + 1
				}
				d = max(d, int(b[i]))
			}
			j2 = j2*10 + d
		}
		//p.Printf("%d", j2)
		s2 += j2
	}

	//for p.NextLine() {
	//	f := ParseInts(p.Line())
	//
	//}

	//g := &grid{Matrix: NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	//for p.NextLine() {
	//	ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell {
	//		return cell{ch: b}
	//	})
	//}
	//PrintGrid(g)

	p.PartOne(s1)
	p.PartTwo(s2)

	//os.Exit(1)
}
