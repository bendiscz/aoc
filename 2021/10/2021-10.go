package main

import (
	"sort"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 10
	example = `

[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]

`
)

func main() {
	Run(year, day, example, solve)
}

var pairs = [255]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var errorScore = [255]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completeScore = [255]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func score(line string) (int, int) {
	stack := Stack[byte]{}
	for _, ch := range []byte(line) {
		if pairs[ch] != 0 {
			stack.Push(pairs[ch])
		} else {
			if stack.Len() == 0 {
				return errorScore[ch], 0
			}
			if ch != stack.Pop() {
				return errorScore[ch], 0
			}
		}
	}

	complete := 0
	for stack.Len() > 0 {
		complete = complete*5 + completeScore[stack.Pop()]
	}

	return 0, complete
}

func solve(p *Problem) {
	total := 0
	var scores []int
	for p.NextLine() {
		score1, score2 := score(p.Line())
		total += score1
		if score2 > 0 {
			scores = append(scores, score2)
		}
	}

	sort.Ints(scores)

	p.PartOne(total)
	p.PartTwo(scores[len(scores)/2])
}
