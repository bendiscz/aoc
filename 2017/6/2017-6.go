package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 6
	example = `

0  2  7  0

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		m := memory(ParseInts(p.Line()))
		c1, c2 := cycle(m)
		p.PartOne(c1)
		p.PartTwo(c2)
	}
}

type memory []int

func (m memory) key() string {
	s := strings.Builder{}
	for _, b := range m {
		s.WriteString(strconv.Itoa(b))
		s.WriteByte(';')
	}
	return s.String()
}

func (m memory) String() string {
	return fmt.Sprintf("%+v", []int(m))
}

func cycle(m memory) (int, int) {
	visited := map[string]int{}
	count := 0

	for {
		k := m.key()
		if c, ok := visited[k]; ok {
			return count, count - c
		}
		visited[k] = count

		i, x := 0, math.MinInt
		for j := 0; j < len(m); j++ {
			if m[j] > x {
				x = m[j]
				i = j
			}
		}

		d := m[i]
		m[i] = 0

		p, q := d/len(m), d%len(m)
		if p > 0 {
			for j := range m {
				m[j] += p
			}
		}
		for j := 1; j <= q; j++ {
			m[(i+j)%len(m)]++
		}

		count++
	}
}
