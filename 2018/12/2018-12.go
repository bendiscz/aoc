package main

import (
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

func solve(p *Problem) {
	line := []byte("...." + strings.TrimPrefix(p.ReadLine(), "initial state: ") + "....")
	rules := map[string]byte{}

	p.SkipLines(1)
	for p.NextLine() {
		rules[p.Line()[:5]] = p.Line()[9]
	}

	p.PartOne(pots(line))
	for i, l := 0, make([]byte, len(line)-4); i < 20; i++ {
		for j := 0; j < len(l); j++ {
			if ch, ok := rules[string(line[j:j+5])]; ok {
				l[j] = ch
			} else {
				l[j] = line[j+2]
				l[j] = '.'
			}
		}
		copy(line[2:], l)
		p.PartOne(pots(line))
	}

	p.PartTwo("TODO")
}

func pots(line []byte) string {
	return string(line[2 : len(line)-2])
}
