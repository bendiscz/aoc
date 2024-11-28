package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 15
	example = `

Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3

`
)

func main() {
	Run(year, day, example, solve)
}

type props struct {
	cap, dur, fla, tex, cal int
}

func generate(ch chan<- []int, limit, index int, a []int) {
	if index == len(a) {
		q := make([]int, len(a))
		copy(q, a)
		ch <- q
		return
	}

	for i := 0; i <= limit; i++ {
		a[index] = i
		generate(ch, limit-i, index+1, a)
	}

	if index == 0 {
		close(ch)
	}
}

func solve(p *Problem) {
	ingredients := []props(nil)
	for p.NextLine() {
		m := p.Parse(`^\w+: capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)$`)
		ingredients = append(ingredients, props{ParseInt(m[1]), ParseInt(m[2]), ParseInt(m[3]), ParseInt(m[4]), ParseInt(m[5])})
	}

	best1, best2 := 0, 0
	ch := make(chan []int, 100)
	go generate(ch, 100, 0, make([]int, len(ingredients)))
	for amounts := range ch {
		cookie := props{}
		for i, n := range amounts {
			cookie.cap += n * ingredients[i].cap
			cookie.dur += n * ingredients[i].dur
			cookie.fla += n * ingredients[i].fla
			cookie.tex += n * ingredients[i].tex
			cookie.cal += n * ingredients[i].cal
		}
		if cookie.cap <= 0 || cookie.dur <= 0 || cookie.fla <= 0 || cookie.tex <= 0 {
			continue
		}
		score := cookie.cap * cookie.dur * cookie.fla * cookie.tex
		best1 = max(score, best1)
		if cookie.cal == 500 {
			best2 = max(score, best2)
		}
	}

	p.PartOne(best1)
	p.PartTwo(best2)
}
