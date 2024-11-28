package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2018
	day     = 2
	example = `

abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	ids := []string(nil)
	for p.NextLine() {
		ids = append(ids, p.Line())
	}

	count2, count3 := 0, 0
	for _, id := range ids {
		h := [26]int{}
		for _, ch := range []byte(id) {
			h[ch-'a']++
		}

		has2, has3 := false, false
		for _, x := range h {
			has2 = has2 || x == 2
			has3 = has3 || x == 3
		}

		if has2 {
			count2++
		}
		if has3 {
			count3++
		}
	}
	p.PartOne(count2 * count3)

	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			d, pos := 0, 0
			for k := 0; d < 2 && k < len(ids[i]); k++ {
				if ids[i][k] != ids[j][k] {
					d++
					pos = k
				}
			}

			if d == 1 {
				p.PartTwo(ids[i][:pos] + ids[i][pos+1:])
			}
		}
	}
}
