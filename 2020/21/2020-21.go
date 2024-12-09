package main

import (
	"cmp"
	. "github.com/bendiscz/aoc/aoc"
	"maps"
	"slices"
	"strings"
)

const (
	year    = 2020
	day     = 21
	example = `

mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)

`
)

func main() {
	Run(year, day, example, solve)
}

type set = Set[string]

func insert(s set, values []string) {
	for _, v := range values {
		s[v] = SET
	}
}

type food struct {
	ing set
	alr set
}

func solve(p *Problem) {
	foods := []*food(nil)
	for p.NextLine() {
		s := SplitFieldsDelim(p.Line(), " ,()")
		i := slices.Index(s, "contains")
		f := &food{ing: set{}, alr: set{}}
		insert(f.ing, s[:i])
		insert(f.alr, s[i+1:])
		foods = append(foods, f)
	}

	alr2f := map[string][]*food{}
	for _, f := range foods {
		for a := range f.alr {
			alr2f[a] = append(alr2f[a], f)
		}
	}

	alr2ing := map[string]set{}
	for a, fs := range alr2f {
		ing := set(nil)
		for _, f := range fs {
			if ing == nil {
				ing = maps.Clone(f.ing)
			} else {
				ing = ing.Intersect(f.ing)
			}
		}
		alr2ing[a] = ing
	}

	ing2alr := map[string]string{}
	for len(alr2ing) > 0 {
		for alr, is := range alr2ing {
			if len(is) == 1 {
				var ing string
				for ing = range is {
					break
				}
				ing2alr[ing] = alr
				delete(alr2ing, alr)
				for alr2, is2 := range alr2ing {
					delete(is2, ing)
					if len(is2) == 0 {
						delete(alr2ing, alr2)
					}
				}
				break
			}
		}
	}

	s1 := 0
	for _, f := range foods {
		for ing := range f.ing {
			if _, ok := ing2alr[ing]; !ok {
				s1++
			}
		}
	}
	p.PartOne(s1)

	is := slices.Collect(maps.Keys(ing2alr))
	slices.SortFunc(is, func(a, b string) int {
		return cmp.Compare(ing2alr[a], ing2alr[b])
	})

	p.PartTwo(strings.Join(is, ","))
}
