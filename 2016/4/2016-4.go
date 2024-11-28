package main

import (
	"fmt"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 4
	example = `

aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]

`
)

func main() {
	Run(year, day, example, solve)
}

type hist struct {
	count [128]int
}

func (h *hist) add(c byte) {
	if c < 'a' || c > 'z' {
		return
	}
	h.count[c]++
}

func (h *hist) takeMax() byte {
	max, hit := -1, '?'
	for c := 'a'; c <= 'z'; c++ {
		if h.count[c] > max {
			hit = c
			max = h.count[c]
		}
	}

	if hit != '?' {
		h.count[hit] = -1
	}
	return byte(hit)
}

func decode(s string, id int) string {
	shift := byte(id % 26)
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			c := s[i] + shift
			if c > 'z' {
				c -= 26
			}
			r[i] = c
		} else {
			r[i] = ' '
		}
	}
	return string(r)
}

func solve(p *Problem) {
	count := 0
	for p.NextLine() {
		m := p.Parse(`^([a-z-]+)-(\d+)\[([a-z]{5})]$`)
		if m == nil {
			continue
		}

		h := hist{}
		for _, c := range []byte(m[1]) {
			h.add(c)
		}

		var sum string
		for i := 0; i < 5; i++ {
			sum += string(h.takeMax())
		}

		if sum != m[3] {
			continue
		}

		id := ParseInt(m[2])
		count += id
		if decode(m[1], id) == "northpole object storage" {
			p.PartTwo(fmt.Sprintf("%s %s", m[2], m[0]))
		}
	}
	p.PartOne(count)
}
