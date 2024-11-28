package main

import (
	"regexp"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 7
	example = `

abba[mnop]qrst
abcd[bddb]xyyx
aaaa[qwer]tyui
ioxxoj[asdfgh]zxcvbn

`
)

func main() {
	Run(year, day, example, solve)
}

var brackets = regexp.MustCompile(`\[[a-z]*]`)

func checkTLS(line []byte) bool {
	ok := false

	part := line[:]
	for len(part) > 0 {
		m := brackets.FindIndex(part)
		if m == nil {
			ok = ok || abba(part)
			part = nil
		} else {
			ok = ok || abba(part[:m[0]])
			if abba(part[m[0]+1 : m[1]-1]) {
				return false
			}
			part = part[m[1]:]
		}
	}

	return ok
}

func checkSSL(line []byte) bool {
	a := make(map[aba]struct{})
	b := make(map[aba]struct{})

	part := line[:]
	for len(part) > 0 {
		m := brackets.FindIndex(part)
		if m == nil {
			abas(a, part)
			part = nil
		} else {
			abas(a, part[:m[0]])
			abas(b, part[m[0]+1:m[1]-1])
			part = part[m[1]:]
		}
	}

	for o := range a {
		if _, ok := b[o.inv()]; ok {
			return true
		}
	}

	return false
}

func abba(s []byte) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] && s[i+1] == s[i+2] && s[i] == s[i+3] {
			return true
		}
	}
	return false
}

type aba struct {
	a, b byte
}

func (o aba) inv() aba {
	return aba{o.b, o.a}
}

func abas(m map[aba]struct{}, s []byte) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i] == s[i+2] {
			m[aba{s[i], s[i+1]}] = struct{}{}
		}
	}
}

func solve(p *Problem) {
	countTLS, countSSL := 0, 0
	for p.NextLine() {
		if checkTLS([]byte(p.Line())) {
			countTLS++
		}
		if checkSSL([]byte(p.Line())) {
			countSSL++
		}
	}
	p.PartOne(countTLS)
	p.PartTwo(countSSL)
}
