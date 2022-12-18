package main

import (
	"log"
	"sort"
	"strings"
	"time"
	"unicode"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 8
	example = `

be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce

`
)

func main() {
	Run(year, day, example, solve)
}

var digits = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

var perms [][7]int

func init() {
	t0 := time.Now()
	var perm [7]int
	makePerms(perm, 0)
	log.Printf("⏱ permutations initialized in %v", time.Since(t0))
}

func makePerms(perm [7]int, index int) {
	if index == 7 {
		var hit [7]int
		copy(hit[:], perm[:])
		perms = append(perms, hit)
		return
	}

loop:
	for d := 0; d < 7; d++ {
		for i := 0; i < index; i++ {
			if perm[i] == d {
				continue loop
			}
		}

		perm[index] = d
		makePerms(perm, index+1)
	}
}

func solve(p *Problem) {
	count1, count2 := 0, 0
	for p.NextLine() {
		fields := strings.FieldsFunc(p.Line(), func(r rune) bool { return unicode.IsSpace(r) || r == '|' })

		for _, f := range fields[10:] {
			if len(f) == 2 || len(f) == 3 || len(f) == 4 || len(f) == 7 {
				count1++
			}
		}

		if x, ok := findMatch(fields); ok {
			count2 += x
		}
	}

	p.PartOne(count1)
	p.PartTwo(count2)
}

func mapDigit(digit string, perm [7]int) (int, bool) {
	var mapped []byte
	for i := 0; i < len(digit); i++ {
		d := perm[digit[i]-'a'] + 'a'
		mapped = append(mapped, byte(d))
	}

	sort.Slice(mapped, func(i, j int) bool {
		return mapped[i] < mapped[j]
	})

	x, ok := digits[string(mapped)]
	return x, ok
}

func findMatch(fields []string) (int, bool) {
loop:
	for _, perm := range perms {
		count := 0
		for i, f := range fields {
			x, ok := mapDigit(f, perm)
			if !ok {
				continue loop
			}
			if i >= 10 {
				count = count*10 + x
			}
		}
		return count, true
	}
	return 0, false
}
