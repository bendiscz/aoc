package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"math"
)

const (
	year    = 2024
	day     = 21
	example = `

029A
980A
179A
456A
379A

`
)

func main() {
	Run(year, day, example, solve)
}

type pad struct {
	keys   []string
	coords map[byte]XY
}

func newPad(s string) pad {
	f := SplitFieldsDelim(s, "|")
	dx, dy := len(f[0]), len(f)
	p := pad{keys: f, coords: map[byte]XY{}}
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			p.coords[p.keys[y][x]] = XY{X: x, Y: y}
		}
	}
	return p
}

func (p pad) move(at XY, key byte) ([]string, XY) {
	to := p.coords[key]
	return p.searchDirs(at, to, ""), to
}

func (p pad) searchDirs(at, to XY, seq string) []string {
	if at == to {
		return []string{seq + "A"}
	}

	diff, s := to.Sub(at), []string(nil)
	for k, d := range dirs {
		if diff.X*d.X > 0 || diff.Y*d.Y > 0 {
			n := at.Add(d)
			if p.keys[n.Y][n.X] != ' ' {
				s = append(s, p.searchDirs(n, to, seq+k)...)
			}
		}
	}
	return s
}

type state struct {
	seq   string
	depth int
	limit int
}

var (
	dirs = map[string]XY{"^": NegY, "v": PosY, "<": NegX, ">": PosX}

	cache = map[state]int{}

	numPad = newPad("789|456|123| 0A")
	dirPad = newPad(" ^A|<v>")
)

func solve(p *Problem) {
	s1, s2 := 0, 0
	for p.NextLine() {
		s1 += complexity(p.Line(), 2)
		s2 += complexity(p.Line(), 25)
	}
	p.PartOne(s1)
	p.PartTwo(s2)
}

func complexity(code string, depth int) int {
	return ParseInt(code[:len(code)-1]) * countPresses(numPad, code, depth, depth)
}

func countPresses(p pad, code string, depth, limit int) int {
	if depth < 0 {
		return len(code)
	}
	if r, ok := cache[state{code, depth, limit}]; ok {
		return r
	}

	r, at := 0, p.coords['A']
	for i := 0; i < len(code); i++ {
		seqs, best := []string(nil), math.MaxInt
		seqs, at = p.move(at, code[i])
		for _, seq := range seqs {
			best = min(best, countPresses(dirPad, seq, depth-1, limit))
		}
		r += best
	}

	cache[state{code, depth, limit}] = r
	return r
}
