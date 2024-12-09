package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2020
	day     = 24
	example = `

sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	tiles := Set[XY]{}
	for p.NextLine() {
		s := p.Line()
		xy := XY{}
		for i := 0; i < len(s); i++ {
			var dir string
			if s[i] == 'n' || s[i] == 's' {
				dir = s[i : i+2]
				i++
			} else {
				dir = s[i : i+1]
			}
			xy = move(xy, dir)
		}
		if tiles.Contains(xy) {
			delete(tiles, xy)
		} else {
			tiles[xy] = SET
		}
	}
	p.PartOne(len(tiles))

	for i := 0; i < 100; i++ {
		tiles = apply(tiles)
	}
	p.PartTwo(len(tiles))
}

func move(xy XY, dir string) XY {
	switch dir {
	case "e":
		return xy.Add(XY{2, 0})
	case "se":
		return xy.Add(XY{1, 1})
	case "ne":
		return xy.Add(XY{1, -1})
	case "w":
		return xy.Add(XY{-2, 0})
	case "sw":
		return xy.Add(XY{-1, 1})
	case "nw":
		return xy.Add(XY{-1, -1})
	default:
		panic("undefined direction: " + dir)
	}
}

func adjacent(xy XY) []XY {
	return []XY{
		xy.Add(XY{2, 0}),
		xy.Add(XY{1, 1}),
		xy.Add(XY{1, -1}),
		xy.Add(XY{-2, 0}),
		xy.Add(XY{-1, 1}),
		xy.Add(XY{-1, -1}),
	}
}

func apply(black Set[XY]) Set[XY] {
	white := Set[XY]{}
	for t := range black {
		for _, xy := range adjacent(t) {
			if !black.Contains(xy) {
				white[xy] = SET
			}
		}
	}

	next := Set[XY]{}
	for t := range black {
		c := count(t, black)
		if c == 1 || c == 2 {
			next[t] = SET
		}
	}
	for t := range white {
		c := count(t, black)
		if c == 2 {
			next[t] = SET
		}
	}

	return next
}

func count(tile XY, black Set[XY]) int {
	s := 0
	for _, xy := range adjacent(tile) {
		if black.Contains(xy) {
			s++
		}
	}
	return s
}
