package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"maps"
)

const (
	year    = 2020
	day     = 19
	example = `

0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb

`
)

func main() {
	Run(year, day, example, solve)
}

type rule struct {
	ch string
	v  []string
}

type rules map[uint8]rule

func solve(p *Problem) {
	rs := rules{}
	for p.NextLine() {
		if p.Line() == "" {
			break
		}

		parts := SplitFieldsDelim(p.Line(), ":|")
		id := uint8(ParseInt(parts[0]))
		r := rule{}
		if len(parts) == 2 && parts[1][len(parts[1])-1] == '"' {
			r.ch = parts[1][len(parts[1])-2 : len(parts[1])-1]
		} else {
			for _, v := range parts[1:] {
				vs := ParseInts(v)
				b := make([]uint8, len(vs))
				for i, x := range vs {
					b[i] = uint8(x)
				}
				s := string(b)
				r.v = append(r.v, s)
			}
		}
		rs[id] = r
	}

	rs2 := maps.Clone(rs)
	r8 := rs2[8]
	r8.v = append(r8.v, string([]uint8{42, 8}))
	rs2[8] = r8
	r11 := rs2[11]
	r11.v = append(r11.v, string([]uint8{42, 11, 31}))
	rs2[11] = r11

	s1, s2 := 0, 0
	for p.NextLine() {
		s := p.Line()
		b0 := string([]uint8{0})
		if check(rs, s, b0) {
			s1++
		}
		if check(rs2, s, b0) {
			s2++
		}
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}

func check(rs rules, s, b string) bool {
	if len(s) == 0 || len(b) == 0 {
		return len(s) == len(b)
	}

	r := rs[b[0]]
	if r.ch != "" {
		return r.ch[0] == s[0] && check(rs, s[1:], b[1:])
	} else {
		for _, v := range r.v {
			if check(rs, s, v+b[1:]) {
				return true
			}
		}
		return false
	}
}
