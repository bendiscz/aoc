package main

import (
	"cmp"
	. "github.com/bendiscz/aoc/aoc"
	"sort"
)

const (
	year    = 2019
	day     = 14
	example = `

2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF
`
)

func main() {
	Run(year, day, example, solve)
}

type reaction struct {
	amount   int
	requires map[string]int
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	rs := map[string]*reaction{}
	for p.NextLine() {
		f := SplitFieldsDelim(p.Line(), " ,=>")
		l := len(f)

		r := &reaction{amount: ParseInt(f[l-2]), requires: map[string]int{}}
		rs[f[l-1]] = r
		for i := 0; i < l-2; i += 2 {
			r.requires[f[i+1]] = ParseInt(f[i])
		}
	}

	s1 = count(rs, map[string]int{}, "FUEL", 1)
	p.PartOne(s1)

	const ORE = 1_000_000_000_000
	fuel := 2
	for count(rs, map[string]int{}, "FUEL", fuel) < ORE {
		fuel *= 2
	}
	s2, _ = sort.Find(fuel, func(i int) int {
		return cmp.Compare(ORE, count(rs, map[string]int{}, "FUEL", i))
	})
	p.PartTwo(s2 - 1)
}

func count(rs map[string]*reaction, cache map[string]int, ingredient string, amount int) int {
	if ingredient == "ORE" {
		return amount
	}

	r := rs[ingredient]
	if c := cache[ingredient]; c > 0 {
		if c > amount {
			cache[ingredient] = c - amount
			return 0
		}
		delete(cache, ingredient)
		amount -= c
	}

	n, s := amount/r.amount+1, r.amount-amount%r.amount
	if s == r.amount {
		n, s = n-1, 0
	}

	total := 0
	for i, a := range r.requires {
		total += count(rs, cache, i, n*a)
	}

	if s > 0 {
		cache[ingredient] += s
	}

	return total
}
