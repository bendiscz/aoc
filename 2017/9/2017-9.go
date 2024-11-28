package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 9
	example = `

<{!>}>
{}
{{{}}}
{{},{}}
{{{},{},{{}}}}
{{<a!>},{<a!>},{<a!>},{<ab>}}
<{o"i!a,<{i<a>

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		parser := &parser{s: p.Line()}
		p.PartOne(parser.parse(1))
		p.PartTwo(parser.gc)
	}
}

type parser struct {
	s      string
	ch     byte
	pushed bool
	gc     int
}

func (p *parser) pushBack(ch byte) {
	p.ch = ch
	p.pushed = true
}

func (p *parser) parse(depth int) int {
	ch := p.nextByte()
	if ch == '<' {
		// garbage
		for p.nextByte() != '>' {
			p.gc++
		}
		return 0
	} else {
		// group
		sum := depth
		if n := p.nextByte(); n == '}' {
			return sum
		} else {
			p.pushBack(n)
		}

		for {
			sum += p.parse(depth + 1)
			ch = p.nextByte()
			if ch != ',' {
				return sum
			}
		}
	}
}

func (p *parser) nextByte() byte {
	if p.pushed {
		p.pushed = false
		return p.ch
	}

	i := 0
	for p.s[i] == '!' {
		i += 2
	}

	ch := p.s[i]
	p.s = p.s[i+1:]
	return ch
}
