package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 8
	example = `

""
"abc"
"aaa\"aaa"
"\x27"

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	count := 0
	for p.NextLine() {
		line := p.Line()
		lineCount := 0
		for i := 1; i < len(line)-1; i++ {
			if line[i] == '\\' {
				i++
				if line[i] == 'x' {
					i += 2
				}
			}
			lineCount++
		}
		count += len(line) - lineCount
	}

	p.PartOne(count)

	count = 0
	p.Reset()
	for p.NextLine() {
		line := p.Line()
		lineCount := 2
		for i := 0; i < len(line); i++ {
			if line[i] == '\\' || line[i] == '"' {
				lineCount++
			}
			lineCount++
		}
		count += lineCount - len(line)
	}

	p.PartTwo(count)
}
