package main

import (
	"bytes"
	"slices"
	"strconv"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 12
	example = `

initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #

`
)

func main() {
	Run(year, day, example, solve)
}

type pots struct {
	b      []byte
	offset int
}

func (p pots) String() string {
	return strconv.Itoa(p.offset) + " " + string(p.b)
}

func (p pots) env(i int) string {
	b := [5]byte{'.', '.', '.', '.', '.'}
	for j := 0; j < 5; j++ {
		k := i - 2 + j
		if k >= 0 && k < len(p.b) {
			b[j] = p.b[k]
		}
	}
	return string(b[:])
}

func (p pots) live(rules map[string]byte) pots {
	const margin = 2
	p2 := pots{b: make([]byte, len(p.b)+2*margin), offset: p.offset - margin}
	for i := 0; i < len(p2.b); i++ {
		if b, ok := rules[p.env(i-margin)]; ok {
			p2.b[i] = b
		} else {
			p2.b[i] = '.'
		}
	}

	i1 := bytes.IndexByte(p2.b, '#')
	i2 := bytes.LastIndexByte(p2.b, '#')
	return pots{b: p2.b[i1 : i2+1], offset: p2.offset + i1}
}

func (p pots) count() int {
	s := 0
	for i, b := range p.b {
		if b == '#' {
			s += i + p.offset
		}
	}
	return s
}

func solve(p *Problem) {
	ps := pots{b: []byte(strings.TrimPrefix(p.ReadLine(), "initial state: "))}
	rules := map[string]byte{}

	p.SkipLines(1)
	for p.NextLine() {
		rules[p.Line()[:5]] = p.Line()[9]
	}

	p.Printf("%s", ps)
	var i, d int
	for i = 0; i < 20; i++ {
		ps = ps.live(rules)
	}
	p.PartOne(ps.count())

	for {
		ps2 := ps.live(rules)
		if slices.Equal(ps2.b, ps.b) {
			d = ps2.offset - ps.offset
			break
		}
		ps = ps2
		i++
	}
	ps.offset += (50000000000 - i) * d
	p.PartTwo(ps.count())
}
