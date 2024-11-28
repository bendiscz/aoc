package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 6
	example = `

turn on 0,0 through 999,999
toggle 0,0 through 999,0
turn off 499,499 through 500,500

`
)

func main() {
	Run(year, day, example, solve)
}

func sort2(x, y int) (int, int) {
	if x < y {
		return x, y
	} else {
		return y, x
	}
}

func solve(p *Problem) {
	var digital [1000][1000]bool
	var analog [1000][1000]int

	for p.NextLine() {
		m := p.Parse(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
		x1, y1 := ParseInt(m[2]), ParseInt(m[3])
		x2, y2 := ParseInt(m[4]), ParseInt(m[5])

		x1, x2 = sort2(x1, x2)
		y1, y2 = sort2(y1, y2)

		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {
				switch m[1] {
				case "turn on":
					digital[x][y] = true
					analog[x][y]++
				case "turn off":
					digital[x][y] = false
					if analog[x][y] > 0 {
						analog[x][y]--
					}
				case "toggle":
					digital[x][y] = !digital[x][y]
					analog[x][y] += 2
				}
			}
		}
	}

	count1, count2 := 0, 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if digital[x][y] {
				count1++
			}
			count2 += analog[x][y]
		}
	}

	p.PartOne(count1)
	p.PartTwo(count2)
}
