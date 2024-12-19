package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2024
	day     = 19
	example = `

r, wr, b, g, bwu, rb, gb, br

bwurrg
brwrr
bggr
gbbr
rrbgbr
ubwu
brgr
bbrgwb

`
)

func main() {
	Run(year, day, example, solve)
}

var index = []int{'b': 0, 'w': 1, 'u': 2, 'r': 3, 'g': 4}

type trie struct {
	term bool
	n    [5]int
}

func buildTrie(patterns []string) []trie {
	t := []trie{{}}
	for _, p := range patterns {
		ti := 0
		for i := 0; i < len(p); i++ {
			ch := index[p[i]]
			ni := t[ti].n[ch]
			if ni == 0 {
				ni = len(t)
				t[ti].n[ch] = ni
				t = append(t, trie{})
			}
			ti = ni
		}
		t[ti].term = true
	}
	return t
}

func matchTrie(t []trie, towel string, lengths []int) []int {
	for i, ti := 0, 0; ; i++ {
		if t[ti].term {
			lengths = append(lengths, i)
		}
		if i == len(towel) {
			break
		}
		ti = t[ti].n[index[towel[i]]]
		if ti == 0 {
			break
		}
	}
	return lengths
}

func solve(p *Problem) {
	t := buildTrie(SplitFields(p.ReadLine()))
	p.NextLine()

	s1, s2 := 0, 0
	for p.NextLine() {
		c := count(t, p.Line())
		if c > 0 {
			s1++
		}
		s2 += c
	}
	p.PartOne(s1)
	p.PartTwo(s2)
}

func count(t []trie, towel string) int {
	d, ls := make([]int, len(towel)+1), make([]int, 0, 64)
	d[0] = 1
	for i := 0; i < len(towel); i++ {
		ls = matchTrie(t, towel[i:], ls[:0])
		for _, l := range ls {
			d[i+l] += d[i]
		}
	}
	return d[len(d)-1]
}
