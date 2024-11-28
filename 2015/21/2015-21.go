package main

import (
	"sort"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2015
	day     = 21
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

type entity struct {
	hp  int
	ar  int
	dmg int
}

type item struct {
	name string
	cost int
	dmg  int
	ar   int
}

type build struct {
	items []item
	cost  int
	dmg   int
	ar    int
}

func (b build) with(i item) build {
	items := make([]item, len(b.items)+1)
	copy(items, b.items)
	items[len(b.items)] = i
	return build{
		items: items,
		cost:  b.cost + i.cost,
		dmg:   b.dmg + i.dmg,
		ar:    b.ar + i.ar,
	}
}

func solve(p *Problem) {
	boss := entity{}
	p.NextLine()
	p.Scanf("Hit Points: %d", &boss.hp)
	p.NextLine()
	p.Scanf("Damage: %d", &boss.dmg)
	p.NextLine()
	p.Scanf("Armor: %d", &boss.ar)

	builds := generateBuilds()

	for _, b := range builds {
		if playerWins(b, boss) {
			p.PartOne(b.cost)
			break
		}
	}

	for i := len(builds) - 1; i >= 0; i-- {
		b := builds[i]
		p.Printf("cost: %d, dmg: %d, ar: %d", b.cost, b.dmg, b.ar)
		if !playerWins(b, boss) {
			p.PartTwo(b.cost)
			break
		}
	}
}

func generateBuilds() []build {
	builds := pickWeapon(build{})
	sort.Slice(builds, func(i, j int) bool {
		return builds[i].cost < builds[j].cost
	})
	return builds
}

func pickWeapon(b build) []build {
	weapons := [...]item{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsowrd", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	}

	builds := []build(nil)
	for _, w := range weapons {
		builds = append(builds, pickArmor(b.with(w))...)
	}
	return builds
}

func pickArmor(b build) []build {
	armors := [...]item{
		{"No Armor", 0, 0, 0},
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
	}

	builds := []build(nil)
	for _, a := range armors {
		builds = append(builds, pickRings(b.with(a))...)
	}
	return builds
}

func pickRings(b build) []build {
	rings := [...]item{
		{"No Ring", 0, 0, 0},
		{"No Ring", 0, 0, 0},
		{"Damage +1", 25, 1, 0},
		{"Damage +2", 50, 2, 0},
		{"Damage +3", 100, 3, 0},
		{"Armor +1", 20, 0, 1},
		{"Armor +2", 40, 0, 2},
		{"Armor +3", 80, 0, 3},
	}
	return pickRing(b, rings[:], 0)
}

func pickRing(b build, rings []item, count int) []build {
	builds := []build(nil)
	if count < 2 {
		for i, r := range rings {
			builds = append(builds, pickRing(b.with(r), rings[i+1:], count+1)...)
		}
	} else {
		builds = append(builds, b)
	}
	return builds
}

func playerWins(b build, boss entity) bool {
	return battle(entity{
		hp:  100,
		ar:  b.ar,
		dmg: b.dmg,
	}, boss)
}

func battle(player, boss entity) bool {
	pd := max(1, player.dmg-boss.ar)
	bd := max(1, boss.dmg-player.ar)

	return (boss.hp+pd-1)/pd <= (player.hp+bd-1)/bd
}
