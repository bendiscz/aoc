package main

import (
	"crypto"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 17
	example = `

ihgpwlah
kglvqrro
ulqzkmiv

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	for p.NextLine() {
		code := p.Line()

		type state struct {
			path string
			pos  XY
		}

		r1, r2 := "", 0

		q := Queue[state]{}
		q.Push(state{"", XY{}})
		for q.Len() > 0 {
			s := q.Pop()
			if !Square(4).HasInside(s.pos) {
				continue
			}
			if s.pos == (XY{3, 3}) {
				if r1 == "" {
					r1 = s.path
				}
				r2 = len(s.path)
				continue
			}

			for _, d := range checkDoors(code, s.path) {
				q.Push(state{s.path + string(d), s.pos.Add(dirs[d[0]])})
			}
		}

		p.PartOne(r1)
		p.PartTwo(r2)
	}
}

type dir string

const (
	up    dir = "U"
	down      = "D"
	left      = "L"
	right     = "R"
)

var dirs = [...]XY{
	'U': NegY,
	'D': PosY,
	'L': NegX,
	'R': PosX,
}

var md5 = crypto.MD5.New()

func checkDoors(code, path string) []dir {
	md5.Reset()
	md5.Write([]byte(code))
	md5.Write([]byte(path))
	hash := md5.Sum(nil)

	d := make([]dir, 0, 4)
	if hash[0]>>4 > 10 {
		d = append(d, up)
	}
	if hash[0]&15 > 10 {
		d = append(d, down)
	}
	if hash[1]>>4 > 10 {
		d = append(d, left)
	}
	if hash[1]&15 > 10 {
		d = append(d, right)
	}
	return d
}
