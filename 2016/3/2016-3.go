package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year = 2016
	day  = 3
)

func main() {
	Run(year, day, "", solve)
}

func parseLine(line string) (int, int, int) {
	fields := strings.Fields(line)
	return ParseInt(fields[0]), ParseInt(fields[1]), ParseInt(fields[2])
}

func isTriangle(x, y, z int) bool {
	return x+y > z && y+z > x && x+z > y
}

func solvePartOne(p *Problem) int {
	count := 0
	for p.NextLine() {
		x, y, z := parseLine(p.Line())
		if isTriangle(x, y, z) {
			count++
		}
	}
	return count
}

func solvePartTwo(p *Problem) int {
	count := 0
	for p.NextLine() {
		x0, x1, x2 := parseLine(p.Line())
		if !p.NextLine() {
			break
		}
		y0, y1, y2 := parseLine(p.Line())
		if !p.NextLine() {
			break
		}
		z0, z1, z2 := parseLine(p.Line())

		if isTriangle(x0, y0, z0) {
			count++
		}
		if isTriangle(x1, y1, z1) {
			count++
		}
		if isTriangle(x2, y2, z2) {
			count++
		}
	}
	return count
}

func solve(p *Problem) {
	p.PartOne(solvePartOne(p))
	p.Reset()
	p.PartTwo(solvePartTwo(p))
}
