package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"math"
)

const (
	year    = 2019
	day     = 8
	example = `

123456789012

`
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

func (c cell) String() string {
	switch c.ch {
	case '0':
		return "."
	case '1':
		return "#"
	default:
		return " "
	}
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	w, h := 25, 6
	if p.Example() {
		w, h = 3, 2
	}

	p.NextLine()
	img := []byte(p.Line())
	layers := []grid(nil)

	n := w * h
	h0 := []int{math.MaxInt}
	for i := 0; i < len(img); i += n {
		hist := histogram(img[i : i+n])
		if hist[0] < h0[0] {
			h0 = hist
		}

		layer := grid{NewMatrix[cell](Rectangle(w, h))}
		for y := 0; y < h; y++ {
			off := i + y*w
			ParseVectorFunc(layer.Row(y), string(img[off:off+w]), func(b byte) cell { return cell{b} })
		}
		layers = append(layers, layer)
	}
	p.PartOne(h0[1] * h0[2])

	result := grid{NewMatrix[cell](Rectangle(w, h))}
	for xy, c := range result.All() {
		var ch byte
		for _, l := range layers {
			ch = l.At(xy).ch
			if ch != '2' {
				break
			}
		}
		c.ch = ch
	}
	PrintGrid(result)
}

func histogram(layer []byte) []int {
	h := [10]int{}
	for _, b := range layer {
		h[b-'0']++
	}
	return h[:]
}
