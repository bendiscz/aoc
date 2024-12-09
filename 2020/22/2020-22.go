package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"slices"
)

const (
	year    = 2020
	day     = 22
	example = `

Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10

`
)

func main() {
	Run(year, day, example, solve)
}

type deck struct {
	cards []byte
}

func newDeck() *deck {
	return &deck{}
}

func (d *deck) len() int {
	return len(d.cards)
}

func (d *deck) append(c int) {
	d.cards = append(d.cards, byte(c))
}

func (d *deck) take() int {
	c := d.cards[0]
	d.cards = slices.Delete(d.cards, 0, 1)
	return int(c)
}

func (d *deck) cloneAll() *deck {
	return &deck{slices.Clone(d.cards)}
}

func (d *deck) cloneSome(n int) *deck {
	return &deck{slices.Clone(d.cards[:n])}
}

func (d *deck) score() int {
	s := 0
	for i, c := range d.cards {
		s += int(c) * (d.len() - i)
	}
	return s
}

func solve(p *Problem) {
	deck1, deck2 := newDeck(), newDeck()
	p.NextLine()
	for p.NextLine() {
		if p.Line() == "" {
			break
		}
		deck1.append(ParseInt(p.Line()))
	}
	p.NextLine()
	for p.NextLine() {
		deck2.append(ParseInt(p.Line()))
	}

	d1, d2 := deck1.cloneAll(), deck2.cloneAll()
	for d1.len() > 0 && d2.len() > 0 {
		c1, c2 := d1.take(), d2.take()
		if c1 > c2 {
			d1.append(c1)
			d1.append(c2)
		} else {
			d2.append(c2)
			d2.append(c1)
		}
	}

	p.PartOne(win(d1, d2).score())

	playGame(deck1, deck2)
	p.PartTwo(win(deck1, deck2).score())
}

func win(d1, d2 *deck) *deck {
	if d1.len() > 0 {
		return d1
	} else {
		return d2
	}
}

func snap(d1, d2 *deck) snapshot {
	return snapshot{string(d1.cards), string(d2.cards)}
}

type snapshot struct {
	d1, d2 string
}

func playGame(d1, d2 *deck) (result int) {
	states := Set[snapshot]{}
	for {
		state := snap(d1, d2)
		if states.Contains(state) {
			return 1
		}
		states[state] = SET

		c1, c2 := d1.take(), d2.take()
		var winner int
		if c1 <= d1.len() && c2 <= d2.len() {
			winner = playGame(d1.cloneSome(c1), d2.cloneSome(c2))
		} else {
			if c1 > c2 {
				winner = 1
			} else {
				winner = 2
			}
		}

		switch winner {
		case 1:
			d1.append(c1)
			d1.append(c2)
		case 2:
			d2.append(c2)
			d2.append(c1)
		}

		if d1.len() == 0 {
			return 2
		}
		if d2.len() == 0 {
			return 1
		}
	}
}
