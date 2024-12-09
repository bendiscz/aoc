package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 7
	example = `

light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.

`
)

func main() {
	Run(year, day, example, solve)
}

type bag struct {
	name         string
	inspected    bool
	containsGold bool
	contents     map[string]int
}

func newBag(name string) *bag {
	return &bag{
		name:     name,
		contents: make(map[string]int),
	}
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	bags := map[string]*bag{}
	for p.NextLine() {
		parseBag(bags, p.Line())
	}

	for n := range bags {
		if checkBag(bags, n) {
			s1++
		}
	}
	p.PartOne(s1)

	s2 = countBags(bags, "shiny gold") - 1
	p.PartTwo(s2)
}

func parseBag(bags map[string]*bag, s string) {
	f := SplitFields(s)
	name := f[0] + " " + f[1]
	b := bags[name]
	if b == nil {
		b = newBag(name)
		bags[name] = b
	}

	if len(f) == 7 {
		return
	}

	for i := 4; i < len(f); i += 4 {
		b.contents[f[i+1]+" "+f[i+2]] = ParseInt(f[i])
	}
}

func checkBag(bags map[string]*bag, name string) bool {
	b := bags[name]
	if b == nil {
		return false
	}
	if b.inspected {
		return b.containsGold
	}
	for n, _ := range b.contents {
		if n == "shiny gold" || checkBag(bags, n) {
			b.containsGold = true
		}
	}
	b.inspected = true
	return b.containsGold
}

func countBags(bags map[string]*bag, name string) int {
	b := bags[name]
	if b == nil {
		return 0
	}
	s := 1
	for n, c := range b.contents {
		s += c * countBags(bags, n)
	}
	return s
}
