package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 15
	example = `

rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		sum1, steps := 0, []string(nil)
		for _, step := range SplitFields(p.Line()) {
			steps = append(steps, step)
			sum1 += hash(step)
		}
		p.PartOne(sum1)

		boxes := make([]box, 256)
		for _, s := range steps {
			if strings.HasSuffix(s, "-") {
				label := s[:len(s)-1]
				b := &boxes[hash(label)]
				for i := 0; i < len(b.lenses); i++ {
					if b.lenses[i].label == label {
						copy(b.lenses[i:], b.lenses[i+1:])
						b.lenses = b.lenses[:len(b.lenses)-1]
						break
					}
				}
			} else {
				label, focal := s[:len(s)-2], int(s[len(s)-1]-'0')
				b := &boxes[hash(label)]
				i := 0
				for i < len(b.lenses) && b.lenses[i].label != label {
					i++
				}
				if i < len(b.lenses) {
					b.lenses[i].focal = focal
				} else {
					b.lenses = append(b.lenses, lens{label, focal})
				}
			}
		}

		sum2 := 0
		for i, b := range boxes {
			for j, l := range b.lenses {
				sum2 += (i + 1) * (j + 1) * l.focal
			}
		}
		p.PartTwo(sum2)
	}
}

type lens struct {
	label string
	focal int
}

type box struct {
	lenses []lens
}

func hash(s string) int {
	h := 0
	for _, ch := range []byte(s) {
		h += int(ch)
		h *= 17
		h %= 256
	}
	return h
}
