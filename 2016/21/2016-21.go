package main

import (
	"bytes"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 21
	example = `

swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	input := "abcdefgh"
	if p.Example() {
		input = "abcde"
	}

	steps := []string(nil)
	for p.NextLine() {
		steps = append(steps, p.Line())
	}

	p.PartOne(scramble(input, steps))
	if p.Example() {
		return
	}

	const scrambled = "fbgdceah"
	for pw := range Permutations([]byte(input)) {
		if scramble(string(pw), steps) == scrambled {
			p.PartTwo(string(pw))
			break
		}
	}
}

func scramble(input string, steps []string) string {
	pw := []byte(input)
	for _, step := range steps {
		f := SplitFields(step)
		switch {
		case strings.HasPrefix(step, "swap p"):
			swapPosition(pw, ParseInt(f[2]), ParseInt(f[5]))
		case strings.HasPrefix(step, "swap l"):
			swapLetter(pw, f[2][0], f[5][0])
		case strings.HasPrefix(step, "rotate l"):
			rotate(pw, -ParseInt(f[2]))
		case strings.HasPrefix(step, "rotate r"):
			rotate(pw, ParseInt(f[2]))
		case strings.HasPrefix(step, "rotate b"):
			rotateLetter(pw, f[6][0])
		case strings.HasPrefix(step, "reverse"):
			reverse(pw, ParseInt(f[2]), ParseInt(f[4])+1)
		case strings.HasPrefix(step, "move"):
			move(pw, ParseInt(f[2]), ParseInt(f[5]))
		}
	}
	return string(pw)
}

func swapPosition(pw []byte, p1, p2 int) {
	pw[p1], pw[p2] = pw[p2], pw[p1]
}

func swapLetter(pw []byte, ch1, ch2 byte) {
	p1, p2 := -1, -1
	for i := 0; p1 < 0 || p2 < 0; i++ {
		if pw[i] == ch1 {
			p1 = i
		}
		if pw[i] == ch2 {
			p2 = i
		}
	}
	swapPosition(pw, p1, p2)
}

func rotate(pw []byte, d int) {
	d = Mod(d, len(pw))
	reverse(pw, 0, len(pw))
	reverse(pw, 0, d)
	reverse(pw, d, len(pw))
}

func rotateLetter(pw []byte, ch byte) {
	i := bytes.IndexByte(pw, ch) + 1
	if i > 4 {
		i++
	}
	rotate(pw, i)
}

func reverse(pw []byte, p1, p2 int) {
	n := (p2 - p1) / 2
	for i := 0; i < n; i++ {
		pw[p1+i], pw[p2-i-1] = pw[p2-i-1], pw[p1+i]
	}
}

func move(pw []byte, p1, p2 int) {
	if p1 == p2 {
		return
	}

	d := Sign(p2 - p1)
	ch := pw[p1]
	for i := p1; i != p2; i += d {
		pw[i] = pw[i+d]
	}
	pw[p2] = ch
}
