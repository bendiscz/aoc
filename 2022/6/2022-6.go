package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 6
	example = `

mjqjpqmgbljsphdztnvjfqwrcgsmlb
bvwbjplbgvbhsrlpgdmjqwftvncz
nppdvjthqldpwncqszvftbrmjlhg
nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg
zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	count := 0
	for p.NextLine() {
		count++
		p.Printf("line #%d", count)
		p.PartOne(detectSequence(p.Line(), 4))
		p.PartTwo(detectSequence(p.Line(), 14))
	}

}

func detectSequence(line string, size int) int {
	counts := map[byte]int{}
	for i := 0; i < len(line); i++ {
		if i >= size {
			ch := line[i-size]
			if counts[ch] > 1 {
				counts[ch] = counts[ch] - 1
			} else {
				delete(counts, ch)
			}
		}

		ch := line[i]
		counts[ch] = counts[ch] + 1

		if len(counts) == size {
			return i + 1
		}
	}
	return 0
}
