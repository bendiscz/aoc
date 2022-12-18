package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 16
	example = `

D2FE28
8A004A801A8002F478
620080001611562C8802118E34
C0015000016115A2E0802F182340
A0016C880162017C3686B18A3D4780
C200B40A82
04005AC33890
880086C3E88112
CE00C43D881120
D8005AC2A8F0
F600BC2D8F
9C005AC2F8F0
9C0141080250320F1802104A08

`
)

func main() {
	Run(year, day, example, solve)
}

type bits struct {
	b []int
	p int
}

func (b *bits) read(n int) int {
	x := 0
	for n > 0 && b.p < len(b.b) {
		x = x*2 + b.b[b.p]
		b.p++
		n--
	}
	return x
}

func (b *bits) pos() int {
	return b.p
}

func readPacket(b *bits) (sum int, result int) {
	v := b.read(3)
	t := b.read(3)

	if t == 4 {
		return v, readLiteral(b)
	}

	sum, results := readContainer(b)

	switch t {
	case 0:
		for _, x := range results {
			result += x
		}

	case 1:
		result = 1
		for _, x := range results {
			result *= x
		}

	case 2:
		result = math.MaxInt
		for _, x := range results {
			if x < result {
				result = x
			}
		}

	case 3:
		result = math.MinInt
		for _, x := range results {
			if x > result {
				result = x
			}
		}

	case 5:
		if len(results) == 2 && results[0] > results[1] {
			result = 1
		}

	case 6:
		if len(results) == 2 && results[0] < results[1] {
			result = 1
		}

	case 7:
		if len(results) == 2 && results[0] == results[1] {
			result = 1
		}
	}

	return sum + v, result
}

func readLiteral(b *bits) (result int) {
	for {
		g := b.read(5)
		result = result*16 + (g & 15)
		if g&16 == 0 {
			break
		}
	}
	return
}

func readContainer(b *bits) (sum int, results []int) {
	l := b.read(1)
	if l == 0 {
		return readContainer0(b)
	} else {
		return readContainer1(b)
	}
}

func readContainer0(b *bits) (sum int, results []int) {
	bl := b.read(15)
	pos := b.pos()
	for b.pos()-pos < bl {
		v, r := readPacket(b)
		sum += v
		results = append(results, r)
	}
	return
}

func readContainer1(b *bits) (sum int, results []int) {
	pl := b.read(11)
	for i := 0; i < pl; i++ {
		v, r := readPacket(b)
		sum += v
		results = append(results, r)
	}
	return
}

func solve(p *Problem) {
	for p.NextLine() {
		s := p.Line()
		if len(s) > 8 {
			s = s[:8] + "..."
		}
		p.Printf("packet: %s", s)
		solvePacket(p, p.Line())
	}
}

func solvePacket(p *Problem, s string) {
	b := bits{}
	for _, ch := range s {
		var x int
		if ch >= 'A' {
			x = int(ch - 'A' + 10)
		} else {
			x = int(ch - '0')
		}

		b.b = append(b.b, (x&8)>>3, (x&4)>>2, (x&2)>>1, x&1)
	}

	v, r := readPacket(&b)
	p.PartOne(v)
	p.PartTwo(r)
}
