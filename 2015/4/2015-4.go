package main

import (
	"crypto"
	"math"
	"strconv"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 4
	example = `

abcdef
pqrstuv

`
)

func main() {
	Run(year, day, example, solve)
}

func solveLine(p *Problem) {
	md5 := crypto.MD5.New()
	p1, p2 := false, false
	for i := 1; i < math.MaxInt; i++ {
		md5.Reset()
		_, _ = md5.Write([]byte(p.Line() + strconv.Itoa(i)))
		sum := md5.Sum(nil)
		if !p1 && sum[0] == 0 && sum[1] == 0 && sum[2] < 16 {
			p.PartOne(i)
			p1 = true
		}
		if !p2 && sum[0] == 0 && sum[1] == 0 && sum[2] == 0 {
			p.PartTwo(i)
			p2 = true
		}

		if p1 && p2 {
			break
		}
	}
}

func solve(p *Problem) {
	for n := 1; p.NextLine(); n++ {
		p.Printf("line #%d", n)
		solveLine(p)
	}
}
