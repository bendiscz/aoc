package main

import (
	"crypto"
	"math"
	"strconv"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2016
	day     = 5
	example = `

abc

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	input := strings.TrimSpace(p.ReadAll())

	md5 := crypto.MD5.New()
	var code []byte

	for i := 0; i < math.MaxInt; i++ {
		md5.Reset()
		md5.Write([]byte(input + strconv.Itoa(i)))
		sum := md5.Sum(nil)
		if sum[0] == 0 && sum[1] == 0 && sum[2] < 16 {
			code = append(code, Hex[sum[2]])
			if len(code) == 8 {
				break
			}
		}
	}

	p.PartOne(string(code))

	for i := 0; i < 8; i++ {
		code[i] = 0
	}

	for i, count := 0, 0; i < math.MaxInt; i++ {
		md5.Reset()
		md5.Write([]byte(input + strconv.Itoa(i)))
		sum := md5.Sum(nil)
		if sum[0] == 0 && sum[1] == 0 && sum[2] < 16 {
			p := sum[2]
			if p < 0 || p > 7 {
				continue
			}
			if code[p] != 0 {
				continue
			}

			code[p] = Hex[sum[3]/16]
			count++
			if count == 8 {
				break
			}
		}
	}

	p.PartTwo(string(code))
}
