package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 2
	example = `

A Y
B X
C Z

`
)

var outcome = [3][3]int{
	{3, 0, 6},
	{6, 3, 0},
	{0, 6, 3},
}

var play = [3][3]int{
	{2, 0, 1},
	{0, 1, 2},
	{1, 2, 0},
}

func main() {
	Run(year, day, example, solve)
}

func solve(t *Problem) {
	score1, score2 := 0, 0
	for t.NextLine() {
		var line = t.Line()
		elf, me := line[0]-'A', line[2]-'X'
		score1 += int(me) + 1 + outcome[me][elf]
		score2 += play[me][elf] + 1 + int(me)*3
	}

	t.PartOne(score1)
	t.PartTwo(score2)
}
