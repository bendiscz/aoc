package main

import (
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 5
	example = `

ugknbfddgicrmopn
aaa
jchzalrnumimnmhp
haegwjzuvuyypxyu
dvszwmarrgswjxmb

`
)

func main() {
	Run(year, day, example, solve)
}

func checkNicePartOne(word string) bool {
	const vowels = "aeiou"

	wc := 0
	dd := false
	for i := 0; i < len(word); i++ {
		c := word[i]
		if strings.ContainsRune(vowels, rune(c)) {
			wc++
		}
		if i > 0 {
			if word[i-1] == c {
				dd = true
			}
			s := word[i-1 : i+1]
			if s == "ab" || s == "cd" || s == "pq" || s == "xy" {
				return false
			}
		}
	}

	return wc >= 3 && dd
}

func checkNicePartTwo(word string) bool {
	r1 := false
	for i := 0; i < len(word)-3; i++ {
		if strings.Contains(word[i+2:], word[i:i+2]) {
			r1 = true
			break
		}
	}
	if !r1 {
		return false
	}

	r2 := false
	for i := 0; i < len(word)-2; i++ {
		if word[i] == word[i+2] {
			r2 = true
			break
		}
	}
	if !r2 {
		return false
	}

	return true
}

func solve(p *Problem) {
	count1, count2 := 0, 0
	for p.NextLine() {
		if checkNicePartOne(p.Line()) {
			count1++
		}
		if checkNicePartTwo(p.Line()) {
			count2++
		}
	}

	p.PartOne(count1)
	p.PartTwo(count2)
}
