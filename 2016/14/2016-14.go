package main

import (
	"crypto"
	"encoding/hex"
	"strconv"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 14
	example = `

abc

`
)

func main() {
	Run(year, day, example, solve)
}

type key struct {
	i, d  int
	valid bool
}

func solve(p *Problem) {
	seed := p.ReadLine()
	p.PartOne(part(seed, false))
	p.PartTwo(part(seed, true))
}

func part(seed string, two bool) int {
	i, keys, found := 0, []key(nil), 0
	for {
		for j := len(keys) - 1; j >= 0 && i-keys[j].i > 1000; j-- {
			if keys[j].valid {
				found++
				if found == 64 {
					return keys[j].i
				}
			}
			keys = keys[:j]
		}

		s := inspect(md5sum(seed, i, two))

		for _, q := range s.quintets {
			for k := len(keys) - 1; k >= 0; k-- {
				if keys[k].d == q {
					keys[k].valid = true
				}
			}
		}

		if s.candidate {
			keys = append(keys, key{})
			copy(keys[1:], keys)
			keys[0] = key{i: i, d: s.triplet}
		}

		i++
	}
}

var md5 = crypto.MD5.New()

func md5sum(seed string, i int, two bool) []byte {
	s := seed + strconv.Itoa(i)
	md5.Reset()
	_, _ = md5.Write([]byte(s))
	hash := md5.Sum(nil)

	if two {
		for j := 0; j < 2016; j++ {
			md5.Reset()
			_, _ = md5.Write([]byte(hex.EncodeToString(hash)))
			hash = md5.Sum(hash[:0])
		}
	}

	return hash
}

func digit(hash []byte, i int) int {
	b := hash[i/2]
	if i%2 == 0 {
		return int(b >> 4)
	} else {
		return int(b & 0xf)
	}
}

type seqs struct {
	candidate bool
	triplet   int
	quintets  []int
}

func inspect(hash []byte) seqs {
	s := seqs{}
	l, c, d1 := 2*len(hash), 1, digit(hash, 0)
	for i := 1; i <= l; i++ {
		d2 := -1
		if i < l {
			d2 = digit(hash, i)
		}
		if d1 == d2 {
			c++
		} else {
			if c >= 3 {
				if !s.candidate {
					s.candidate = true
					s.triplet = d1
				}
			}
			if c >= 5 {
				s.quintets = append(s.quintets, d1)
			}
			c = 1
		}
		d1 = d2
	}
	return s
}
