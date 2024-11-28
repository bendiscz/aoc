package main

import (
	"sort"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 7
	example = `

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483

`
)

func main() {
	Run(year, day, example, solve)
}

const (
	fiveOfAKind = iota
	fourOfAKind
	fullHouse
	threeOfAKind
	twoPair
	onePair
	highCard
)

const (
	order1 = "AKQJT98765432"
	order2 = "AKQT98765432J"
)

type hand struct {
	cards []byte
	bet   int
	value int
	kind  int
}

func parseHand(s string) *hand {
	f := SplitFields(s)
	h := &hand{
		cards: []byte(f[0]),
		bet:   ParseInt(f[1]),
	}

	return h
}

func (h *hand) rank(order string, jokers bool) {
	h.value = 0
	for _, b := range h.cards {
		h.value <<= 4
		h.value |= strings.IndexByte(order, b)
	}

	const offset = '2'
	counts := ['T' - offset + 1]int{}
	for _, b := range h.cards {
		counts[b-offset]++
	}

	cj := 0
	if jokers {
		const j int = 'J' - offset
		cj = counts[j]
		counts[j] = 0

		mc, mi := 0, j
		for i, c := range counts {
			if c > mc {
				mc = c
				mi = i
			}
		}
		counts[mi] += cj
	}

	c5, c4, c3, c2 := 0, 0, 0, 0
	for _, c := range counts {
		switch c {
		case 5:
			c5++
		case 4:
			c4++
		case 3:
			c3++
		case 2:
			c2++
		}
	}

	switch {
	case c5 == 1:
		h.kind = fiveOfAKind
	case c4 == 1:
		h.kind = fourOfAKind
	case c3 == 1 && c2 == 1:
		h.kind = fullHouse
	case c3 == 1:
		h.kind = threeOfAKind
	case c2 == 2:
		h.kind = twoPair
	case c2 == 1:
		h.kind = onePair
	default:
		h.kind = highCard
	}
}

func solve(p *Problem) {
	hands := []*hand(nil)
	for p.NextLine() {
		h := parseHand(p.Line())
		h.rank(order1, false)
		hands = append(hands, h)
	}
	p.PartOne(rankHands(hands))

	for _, h := range hands {
		h.rank(order2, true)
	}
	p.PartTwo(rankHands(hands))
}

func rankHands(hands []*hand) int {
	sort.Slice(hands, func(i, j int) bool {
		h1, h2 := hands[i], hands[j]
		if h1.kind == h2.kind {
			return h1.value < h2.value
		} else {
			return h1.kind < h2.kind
		}
	})

	rank := 0
	for i, h := range hands {
		rank += h.bet * (len(hands) - i)
	}
	return rank
}
