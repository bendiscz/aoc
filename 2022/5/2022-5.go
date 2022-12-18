package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 5
	example = `

    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	h, w := 0, 0
	for p.NextLine() {
		if !strings.HasPrefix(strings.TrimLeft(p.Line(), " "), "[") {
			break
		}
		h++
	}

	nums := strings.TrimSpace(p.Line())
	w = ParseInt(nums[strings.LastIndexByte(nums, ' ')+1:])

	solvePart(p, w, h, 1)
	solvePart(p, w, h, 2)
}

func solvePart(p *Problem, w, h int, part int) {
	p.Reset()

	stacks := make([]string, w)
	for h > 0 {
		h--
		p.NextLine()
		line := p.Line()
		for i, p := 0, 1; p < len(line); i, p = i+1, p+4 {
			ch := line[p]
			if ch == ' ' {
				continue
			}

			stacks[i] = stacks[i] + string(ch)
		}
	}

	p.NextLine()
	p.NextLine()

	for p.NextLine() {
		m := p.Parse(`move (\d+) from (\d+) to (\d+)`)
		n, s, d := ParseInt(m[1]), ParseInt(m[2])-1, ParseInt(m[3])-1
		p1, p2 := stacks[s][:n], stacks[s][n:]
		stacks[s] = p2
		if part == 1 {
			for _, ch := range []byte(p1) {
				stacks[d] = string(ch) + stacks[d]
			}
		} else {
			stacks[d] = p1 + stacks[d]
		}
	}

	r1 := ""
	for _, s := range stacks {
		if len(s) > 0 {
			r1 = r1 + string(s[0])
		}
	}

	if part == 1 {
		p.PartOne(r1)
	} else {
		p.PartTwo(r1)
	}
}
