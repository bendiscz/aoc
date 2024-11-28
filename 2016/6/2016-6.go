package main

import (
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 6
	example = `

eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar

`
)

func main() {
	Run(year, day, example, solve)
}

type hist struct {
	count [26]int
}

func (h *hist) add(c byte) {
	h.count[c-'a']++
}

func (h *hist) max() byte {
	max, hit := -1, byte(0)
	for c := byte('a'); c <= 'z'; c++ {
		i := c - 'a'
		if h.count[i] > max {
			max = h.count[i]
			hit = c
		}
	}
	return hit
}

func (h *hist) min() byte {
	min, hit := math.MaxInt, byte(0)
	for c := byte('a'); c <= 'z'; c++ {
		i := c - 'a'
		if h.count[i] > 0 && h.count[i] < min {
			min = h.count[i]
			hit = c
		}
	}
	return hit
}

func solve(p *Problem) {
	var hists []*hist
	for p.NextLine() {
		for i, c := range []byte(p.Line()) {
			if i == len(hists) {
				hists = append(hists, &hist{})
			}
			hists[i].add(c)
		}
	}

	var code []byte
	for _, h := range hists {
		code = append(code, h.max())
	}
	p.PartOne(string(code))

	code = nil
	for _, h := range hists {
		code = append(code, h.min())
	}
	p.PartTwo(string(code))
}
