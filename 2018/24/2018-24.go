package main

import (
	"cmp"
	. "github.com/bendiscz/aoc/aoc"
	"regexp"
	"slices"
	"sort"
	"strings"
)

const (
	year    = 2018
	day     = 24
	example = `

Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4

`
)

func main() {
	Run(year, day, example, solve)
}

var (
	groupRE = regexp.MustCompile(`^(\d+) units each with (\d+) hit points(?: \(([\w ,;]*)\))? with an attack that does (\d+) (\w+) damage at initiative (\d+)$`)
)

type group struct {
	army        *army
	id          int
	units       int
	health      int
	attackPower int
	initiative  int
	attackType  string
	immunities  Set[string]
	weaknesses  Set[string]
}

func (g *group) done() bool {
	return g.units == 0
}

func (g *group) power() int {
	return g.units * g.attackPower
}

func (g *group) damageTo(g2 *group) int {
	switch {
	case g2.immunities.Contains(g.attackType):
		return 0
	case g2.weaknesses.Contains(g.attackType):
		return 2 * g.power()
	default:
		return g.power()
	}
}

func (g *group) hit(damage int) {
	g.units = max(0, g.units-damage/g.health)
}

type army struct {
	name   string
	groups []*group
}

func (a *army) done() bool {
	for _, g := range a.groups {
		if !g.done() {
			return false
		}
	}
	return true
}

func (a *army) units() int {
	s := 0
	for _, g := range a.groups {
		s += g.units
	}
	return s
}

func parseGroup(army *army, id int, s string) *group {
	g := &group{army: army, id: id}
	m := groupRE.FindStringSubmatch(s)
	g.units = ParseInt(m[1])
	g.health = ParseInt(m[2])
	g.attackPower = ParseInt(m[4])
	g.attackType = m[5]
	g.initiative = ParseInt(m[6])

	if m[3] != "" {
		for _, p := range SplitFieldsDelim(m[3], ";") {
			f := SplitFields(p)
			switch f[0] {
			case "weak":
				g.weaknesses = Set[string]{}
				for _, e := range f[2:] {
					g.weaknesses[e] = SET
				}
			case "immune":
				g.immunities = Set[string]{}
				for _, e := range f[2:] {
					g.immunities[e] = SET
				}
			}
		}
	}
	return g
}

func parseArmy(p *Problem) *army {
	if !p.NextLine() {
		return nil
	}

	a := &army{name: strings.TrimSuffix(p.Line(), ":")}
	for p.NextLine() && p.Line() != "" {
		a.groups = append(a.groups, parseGroup(a, len(a.groups)+1, p.Line()))
	}
	return a
}

func solve(p *Problem) {
	armies := []*army(nil)
	for a := parseArmy(p); a != nil; a = parseArmy(p) {
		armies = append(armies, a)
	}

	p.PartOne(battle(armies).units())

	units := map[int]int{}
	for maxBoost := 1; ; maxBoost <<= 1 {
		if boost := sort.Search(maxBoost, func(boost int) bool {
			p.Reset()
			armies = nil
			for a := parseArmy(p); a != nil; a = parseArmy(p) {
				if a.name == "Immune System" {
					for _, g := range a.groups {
						g.attackPower += boost
					}
				}
				armies = append(armies, a)
			}
			winner := battle(armies)
			if winner != nil && winner.name == "Immune System" {
				units[boost] = winner.units()
				return true
			}
			return false
		}); boost < maxBoost {
			p.PartTwo(units[boost])
			break
		}
	}
}

func battle(armies []*army) *army {
	for {
		if w, done := fightArmies(armies); done {
			return w
		}
	}
}

type fight struct {
	attacker *group
	target   *group
	damage   int
}

func fightArmies(armies []*army) (winner *army, done bool) {
	fights := selectTargets(armies)
	if fightGroups(fights) {
		return nil, true
	}

	survivors := []*army(nil)
	for _, a := range armies {
		if !a.done() {
			survivors = append(survivors, a)
		}
	}

	switch {
	case len(survivors) == 0:
		return nil, true
	case len(survivors) == 1:
		return survivors[0], true
	default:
		return nil, false
	}
}

func selectTargets(armies []*army) []fight {
	groups := []*group(nil)
	for _, a := range armies {
		groups = append(groups, a.groups...)
	}
	slices.SortFunc(groups, func(a, b *group) int {
		return -cmp.Or(
			cmp.Compare(a.power(), b.power()),
			cmp.Compare(a.initiative, b.initiative))
	})

	fights := []fight(nil)
	targeted := Set[*group]{}

	for _, g := range groups {
		if g.done() {
			continue
		}

		f := fight{attacker: g}
		for _, target := range groups {
			if target.army == g.army || target.done() || targeted.Contains(target) {
				continue
			}

			dmg := g.damageTo(target)
			if dmg == 0 {
				continue
			}

			if f.target == nil {
				f.target = target
				f.damage = dmg
				continue
			}

			if cmp.Or(
				cmp.Compare(dmg, f.damage),
				cmp.Compare(target.power(), f.target.power()),
				cmp.Compare(target.initiative, f.target.initiative)) > 0 {
				f.target = target
				f.damage = dmg
			}
		}
		if f.target != nil {
			targeted[f.target] = SET
			fights = append(fights, f)
		}
	}

	return fights
}

func fightGroups(fights []fight) bool {
	slices.SortFunc(fights, func(a, b fight) int {
		return -cmp.Compare(a.attacker.initiative, b.attacker.initiative)
	})

	tie := true
	for _, f := range fights {
		if f.attacker.done() {
			continue
		}

		units := f.target.units
		f.target.hit(f.attacker.damageTo(f.target))
		if f.target.units < units {
			tie = false
		}
	}
	return tie
}
