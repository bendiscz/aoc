package main

import (
	"fmt"
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 22
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

type state struct {
	boss struct {
		hp  int
		dmg int
	}
	player struct {
		hp   int
		mana int
		ar   int

		spent int

		shield   int
		poison   int
		recharge int
	}
}

func (s *state) tick() {
	// shield
	if s.player.shield > 0 {
		s.player.shield--
	}
	if s.player.shield > 0 {
		s.player.ar = 7
	} else {
		s.player.ar = 0
	}

	// poison
	if s.player.poison > 0 {
		s.boss.hp -= 3
		s.player.poison--
	}

	// recharge
	if s.player.recharge > 0 {
		s.player.mana += 101
		s.player.recharge--
	}
}

type result struct {
	best int
}

func (r *result) checkFinished(s state) bool {
	switch {
	case s.boss.hp <= 0:
		fmt.Printf("!\n")
		r.best = min(r.best, s.player.spent)
		return true

	case s.player.hp <= 0:
		return true

	default:
		return false
	}
}

func solve(p *Problem) {
	s := state{}
	p.NextLine()
	p.Scanf("Hit Points: %d", &s.boss.hp)
	p.NextLine()
	p.Scanf("Damage: %d", &s.boss.dmg)

	s.player.hp = 50
	s.player.mana = 500

	r := result{best: math.MaxInt}
	playPlayerTurn(&r, s, false)
	p.PartOne(r.best)

	r.best = math.MaxInt
	playPlayerTurn(&r, s, true)
	p.PartTwo(r.best)
}

func cast(s state, cost int) (state, bool) {
	if s.player.mana < cost {
		return state{}, false
	}

	s.player.mana -= cost
	s.player.spent += cost
	return s, true
}

func playPlayerTurn(r *result, s state, hard bool) {
	if hard {
		s.player.hp--
		if r.checkFinished(s) {
			return
		}
	}

	s.tick()
	if r.checkFinished(s) {
		return
	}

	if s.player.mana < 53 {
		//fmt.Printf("loss (mana)\n")
		return
	}

	if ns, ok := cast(s, 53); ok {
		ns.boss.hp -= 4
		if r.checkFinished(s) {
			return
		}

		playBossTurn(r, ns, hard)
	}

	if ns, ok := cast(s, 73); ok {
		ns.player.hp += 2
		ns.boss.hp -= 2
		if r.checkFinished(s) {
			return
		}

		playBossTurn(r, ns, hard)
	}

	if ns, ok := cast(s, 113); ok && ns.player.shield == 0 {
		ns.player.shield = 6
		playBossTurn(r, ns, hard)
	}

	if ns, ok := cast(s, 173); ok && ns.player.poison == 0 {
		ns.player.poison = 6
		playBossTurn(r, ns, hard)
	}

	if ns, ok := cast(s, 229); ok && ns.player.recharge == 0 {
		ns.player.recharge = 5
		playBossTurn(r, ns, hard)
	}
}

func playBossTurn(r *result, s state, hard bool) {
	s.tick()
	if r.checkFinished(s) {
		return
	}

	s.player.hp -= max(1, s.boss.dmg-s.player.ar)
	if r.checkFinished(s) {
		return
	}

	playPlayerTurn(r, s, hard)
}
