package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 3
	example = `

>
^>v<
^v^v^v^v^v

`
)

func main() {
	Run(year, day, example, solve)
}

// TODO(mbenda) p.LineBytes() - pouzit

var dirs = map[byte]XY{
	'^': PosY,
	'v': NegY,
	'>': PosX,
	'<': NegX,
}

func solvePartOne(bytes []byte) int {
	count := 0
	visited := make(map[XY]struct{})
	var pos XY

	visited[pos] = struct{}{}
	count++

	for _, d := range bytes {
		pos = pos.Add(dirs[d])

		if _, ok := visited[pos]; !ok {
			count++
			visited[pos] = struct{}{}
		}
	}
	return count
}

func solvePartTwo(bytes []byte) int {
	count := 0
	visited := make(map[XY]struct{})
	var pos1, pos2 XY

	visited[pos1] = struct{}{}
	count++

	for i := 0; i < len(bytes); i += 2 {
		pos1 = pos1.Add(dirs[bytes[i]])
		if i < len(bytes)-1 {
			pos2 = pos2.Add(dirs[bytes[i+1]])
		}

		if _, ok := visited[pos1]; !ok {
			count++
			visited[pos1] = struct{}{}
		}
		if _, ok := visited[pos2]; !ok {
			count++
			visited[pos2] = struct{}{}
		}
	}
	return count
}

func solveLine(p *Problem, line []byte) {
	p.PartOne(solvePartOne(line))
	p.PartTwo(solvePartTwo(line))
}

func solve(p *Problem) {
	for n := 1; p.NextLine(); n++ {
		p.Printf("line #%d", n)
		solveLine(p, p.LineBytes())
	}
}
