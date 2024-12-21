package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 22
	example = `


`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

func (c cell) String() string { return string([]byte{c.ch}) }

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	//g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}

	//s := p.ReadAll()

	for p.NextLine() {

		//f := ParseInts(p.Line())
		//f := ParseUints(p.Line())
		//f := SplitFields(p.Line())
		//f := SplitFieldsDelim(p.Line(), ",;")
		//f := p.Parse(`^(\w+)$`)
		//f := p.Scanf("Begin in state %c.", &start)

		//CopyVector(g.AppendRow(), p.Line())
		//ParseVectorFunc(g.AppendRow(), p.Line(), func(b byte) cell { return cell{ch: b} })

	}

	p.PartOne(s1)

	p.PartTwo(s2)
}
