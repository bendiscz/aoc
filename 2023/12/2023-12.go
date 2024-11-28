package main

import (
	"fmt"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 12
	example = `

???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	sum1, sum2 := 0, 0
	for p.NextLine() {
		i := strings.IndexByte(p.Line(), ' ')
		mask := p.Line()[:i]
		parts := ParseInts(p.Line()[i+1:])
		sum1 += compute(mask, parts)
		mask, parts = unfold(mask, parts)
		sum2 += compute(mask, parts)
	}

	p.PartOne(sum1)
	p.PartTwo(sum2)
}

func unfold(mask string, parts []int) (string, []int) {
	m := fmt.Sprintf("%s?%s?%s?%s?%s", mask, mask, mask, mask, mask)
	p := make([]int, 5*len(parts))
	for i := 0; i < 5; i++ {
		copy(p[i*len(parts):], parts)
	}
	return m, p
}

func compute(mask string, parts []int) int {
	mask = mask + "."

	d := make([][]int, len(parts)+1)
	for i := 0; i < len(d); i++ {
		d[i] = make([]int, len(mask)+1)
	}
	d[0][0] = 1
	for i := 0; i < len(mask) && (mask[i] == '?' || mask[i] == '.'); i++ {
		d[0][i+1] = 1
	}

	s := make([]int, len(parts))
	for i := 1; i < len(s); i++ {
		s[i] = s[i-1] + parts[i-1] + 1
	}

	for i := 0; i < len(parts); i++ {
		f := parts[i]
		for j := s[i]; j < len(mask); j++ {
			n := d[i][j]
			if n == 0 {
				continue
			}

			matches := true
			for k := 0; k < f; k++ {
				if mask[j+k] != '?' && mask[j+k] != '#' {
					matches = false
					break
				}
			}

			if !matches {
				continue
			}

			for k := j + f; k < len(mask) && (mask[k] == '?' || mask[k] == '.'); k++ {
				d[i+1][k+1] += n
			}
		}
	}

	return d[len(parts)][len(mask)]
}

//type state struct {
//	rm, rp int
//}
//
//func search(mask string, parts []int) int {
//	return search0(mask, parts, true, map[state]int{})
//}
//
//func search0(mask string, parts []int, first bool, cache map[state]int) (result int) {
//	s := state{len(mask), len(parts)}
//	if v, ok := cache[s]; ok {
//		return v
//	}
//
//	defer func() {
//		cache[s] = result
//	}()
//
//	if len(parts) == 0 {
//		if checkMask(mask, len(mask), 0) {
//			return 1
//		} else {
//			return 0
//		}
//	}
//
//	f, r := parts[0], 0
//	for i := 1; i < len(parts); i++ {
//		r += parts[i] + 1
//	}
//
//	sum, e := 0, 1
//	if first {
//		e = 0
//	}
//	for e <= len(mask)-r-f {
//		if checkMask(mask, e, f) {
//			sum += search0(mask[e+f:], parts[1:], false, cache)
//		}
//		e++
//	}
//	return sum
//}
//
//func checkMask(mask string, e, f int) bool {
//	if len(mask) < e+f {
//		return false
//	}
//
//	for i := 0; i < e+f; i++ {
//		ch := mask[i]
//		if !(ch == '?' || i < e && ch == '.' || i >= e && ch == '#') {
//			return false
//		}
//	}
//	return true
//}
