package main

import (
	"encoding/json"
	"log"
	"sort"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 13
	example = `

[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]

`
)

func main() {
	Run(year, day, example, solve)
}

type order int

const (
	Unknown order = iota
	Right
	Wrong
)

func compareLists(x, y []any) order {
	n := Min(len(x), len(y))
	for i := 0; i < n; i++ {
		o := compareValues(x[i], y[i])
		if o != Unknown {
			return o
		}
	}

	switch {
	case len(y) > n:
		return Right
	case len(x) > n:
		return Wrong
	default:
		return Unknown
	}
}

func compareValues(x, y any) order {
	xn, xok := x.(float64)
	yn, yok := y.(float64)
	if !xok || !yok {
		return compareLists(valueAsList(x), valueAsList(y))
	}
	switch {
	case xn < yn:
		return Right
	case xn > yn:
		return Wrong
	default:
		return Unknown
	}
}

func valueAsList(x any) []any {
	if n, ok := x.(float64); ok {
		return []any{n}
	} else {
		return x.([]any)
	}
}

func parse(s string) []any {
	var x []any
	if err := json.Unmarshal([]byte(s), &x); err != nil {
		log.Panicf("failed to parse: %s", s)
	}
	return x
}

type packet struct {
	x []any
	d bool
}

func solve(p *Problem) {
	packets := []packet(nil)
	sum := 0
	for i := 1; p.NextLine(); i++ {
		x := parse(p.Line())
		p.NextLine()
		y := parse(p.Line())
		p.NextLine()

		if compareLists(x, y) == Right {
			sum += i
		}

		packets = append(packets, packet{x, false}, packet{y, false})
	}

	p.PartOne(sum)

	packets = append(packets, packet{parse("[[2]]"), true}, packet{parse("[[6]]"), true})
	sort.Slice(packets, func(i, j int) bool { return compareLists(packets[i].x, packets[j].x) == Right })

	key := 1
	for i, pkt := range packets {
		if pkt.d {
			key *= i + 1
		}
	}
	p.PartTwo(key)
}
