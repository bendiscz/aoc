package main

import (
	"fmt"
	. "github.com/bendiscz/aoc/aoc"
	"math"
	"strings"
)

const (
	year    = 2024
	day     = 21
	example = `

029A
980A
179A
456A
379A

`
)

/*
    +---+---+
    | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+

<v<A >>^A vA ^A <vA <A A >>^A A vA <^A >A A vA ^A      |   <vA >^A A <A >A <v<A >A  >^A A A vA <^A >A
v<<A >>^A vA ^A v<<A >>^A A v<A <A >>^A A vA A ^<A >A  |   v<A >^A A <A >A  v<A <A >>^A A A vA ^<A >A ---- moje

* <v<A>>^AvA^A     <vA<AA>>^AAvA<^A>AAvA^A|    <vA>^AA<A>A<v<A>A>^AAAvA<^A>A
* v<<A>>^AvA^Av<<A>>^AAv<A<A>>^AAvAA^<A>Av<A>^AA<A>Av<A<A>>^AAAvA^<A>A ----moje

<A        >A    v<<A          A >^A       A >A         |  vA A ^A <vA A A >^A
<A        >A    <A        A v<A         A >>^A         |  vA A ^A v<A A A >^A ---- moje

* <A>A   v<<AA >^AA   >A|  vAA^A<vAAA>^A
* <A>A<AAv<AA>>^AvAA^Av<AAA>^A ----moje

^A              <<^^A >>A vvvA
^A              ^^<<A >>A vvvA ---- moje

379A

*/

func main() {
	Run(year, day, example, solve)
}

func codeXY(ch byte) XY {
	switch ch {
	case 'A':
		return XY{2, 3}
	case '0':
		return XY{1, 3}
	case '1':
		return XY{0, 2}
	case '2':
		return XY{1, 2}
	case '3':
		return XY{2, 2}
	case '4':
		return XY{0, 1}
	case '5':
		return XY{1, 1}
	case '6':
		return XY{2, 1}
	case '7':
		return XY{0, 0}
	case '8':
		return XY{1, 0}
	case '9':
		return XY{2, 0}
	}
	panic("invalid code")
}

func dirXY(ch byte) XY {
	switch ch {
	case 'A':
		return XY{2, 0}
	case '^':
		return XY{1, 0}
	case '<':
		return XY{0, 1}
	case 'v':
		return XY{1, 1}
	case '>':
		return XY{2, 1}
	}
	panic("invalid dir")
}

var (
	up    = "^^^^^^^^^^^"
	down  = "vvvvvvvvvvv"
	left  = "<<<<<<<<<<<"
	right = ">>>>>>>>>>>"
)

func numSeq(ch1, ch2 byte) string {
	c1, c2 := codeXY(ch1), codeXY(ch2)
	dx, dy := c2.X-c1.X, c2.Y-c1.Y
	sx, sy := "", ""
	if dx < 0 {
		sx = left[:Abs(dx)]
	} else {
		sx = right[:dx]
	}
	if dy < 0 {
		sy = up[:Abs(dy)]
	} else {
		sy = down[:dy]
	}
	if dx < 0 {
		return sy + sx + "A"
	} else {
		return sx + sy + "A"
	}
}

func dirSeq(ch1, ch2 byte) string {
	c1, c2 := dirXY(ch1), dirXY(ch2)
	dx, dy := c2.X-c1.X, c2.Y-c1.Y
	sx, sy := "", ""
	if dx < 0 {
		sx = left[:Abs(dx)]
	} else {
		sx = right[:dx]
	}
	if dy < 0 {
		sy = up[:Abs(dy)]
	} else {
		sy = down[:dy]
	}
	if dx < 0 {
		return sy + sx + "A"
	} else {
		return sx + sy + "A"
	}
}

func filter(dirs []string) []string {
	l := math.MaxInt
	for _, dir := range dirs {
		l = min(l, len(dir))
	}

	r := []string(nil)
	for _, dir := range dirs {
		if l == len(dir) {
			r = append(r, dir)
		}
	}

	l2 := math.MaxInt
	for _, dir := range r {
		cnt := 0
		for i := 0; i < len(dir); i++ {
			if dir[i] == '<' {
				cnt++
			}
		}
		l2 = min(l2, cnt)
	}

	r2 := []string(nil)
	for _, dir := range r {
		cnt := 0
		for i := 0; i < len(dir); i++ {
			if dir[i] == '<' {
				cnt++
			}
		}
		if l2 == cnt {
			r2 = append(r2, dir)
		}
	}

	return r2
}
func filter2(dirs []string) []string {
	r := []string(nil)
	for _, dir := range dirs {
		if check(dir) {
			r = append(r, dir)
		}
	}
	return r
}

func check(seq string) bool {
loop:
	for len(seq) > 0 {
		for _, a := range allowed {
			if strings.HasPrefix(seq, a) {
				seq = seq[len(a):]
				continue loop
			}
		}
		return false
	}
	return true
}

func countAll(seq string, n int) int {
	s := 0
loop:
	for len(seq) > 0 {
		for i, a := range steps {
			if strings.HasPrefix(seq, a) {
				seq = seq[len(a):]
				s += count(i, n)
				continue loop
			}
		}
		panic("not allowed " + seq)
	}
	return s
}

var rules = [...][]int{
	{0},
	{10, 9},
	{10, 5, 0, 8},
	{10, 5, 8},
	{5, 7},
	{2, 8},
	{2, 9, 7},
	{1, 4},
	{1, 0, 6, 7},
	{1, 6, 7},
	{2, 7, 9},
	{2, 7, 5, 8},
}

const N = 12

func count(rule, n int) int {
	c := make([]int, len(steps))
	c[rule] = 1

	for i := 0; i < n; i++ {
		c2 := make([]int, len(steps))
		for j := 0; j < len(c); j++ {
			for _, r := range nextSteps[j] {
				c2[r] += c[j]
			}
		}
		copy(c, c2)
	}

	s := 0
	for i, x := range c {
		s += x * len(steps[i])
	}
	return s
}

var allowed = []string{
	"A",    //  0 0
	"vA",   //  1 39   | a9
	"v<<A", //  2 3508 | a508
	"v<A",  //  3 358  | a58
	"^A",   //  4 57   | 57
	"<A",   //  5 28   | 28
	"<^A",  //  6 297  | 297
	">A",   //  7 14   | 14
	">>^A", //  8 1067 | 1067
	">^A",  //  9 167  | 167

	"<vA",  //  a (10) | 279
	"<v<A", //  b (11) | 2758
}

var steps = []string{
	"A", //  0 0
	"<A",
	"<<A",
	"<vA",
	"<v<A",
	"<^A",
	">A",
	">>A",
	">^A",
	">^>A",
	">>^A",
	">vA",
	"vA",
	"^A",
	"^>A",
	"^<A",
	"v<A",
	"v<<A",
	"v>A",
}

var nextSteps [][]int

func findBestNextStep(step string) (string, string) {
	i0, l0, s0 := 0, math.MaxInt, ""
	seq1 := dirSeqAll(step)
	for i := 0; i < len(seq1); i++ {
		seq2 := dirSeqAll(seq1[i])
		for j := 0; j < len(seq2); j++ {
			seq3 := dirSeqAll(seq2[j])
			for k := 0; k < len(seq3); k++ {
				l := len(seq3[k])
				if l < l0 {
					l0 = l
					i0 = i
					s0 = seq3[k]
				}
			}
		}
	}
	return seq1[i0], s0
}

func toSteps(seq0 string) []int {
	s := []int(nil)
	seq := seq0
loop:
	for len(seq) > 0 {
		for i, a := range steps {
			if strings.HasPrefix(seq, a) {
				seq = seq[len(a):]
				s = append(s, i)
				continue loop
			}
		}
		panic("not allowed " + seq0 + " " + seq)
	}
	return s

}

func try2() {
	for i := 0; i < len(steps); i++ {
		s0, _ := findBestNextStep(steps[i])
		next := toSteps(s0)
		fmt.Printf("%s %s %v\n", steps[i], s0, next)
		nextSteps = append(nextSteps, next)
	}
}

// vA 1
// 39
// 358167 v<A<A>>^AvA<^A>A 3+2+4+2+3+2
//        v<A<A>>^AvA<^A>A

//    +---+---+
//    | ^ | A |
//+---+---+---+
//| < | v | > |
//+---+---+---+

//<v<A >>^A A <vA <A >>^A A vA A <^A >A <vA >^A <A >A <vA >^A <A >A <v<A >A >^A A vA <^A >A
//<A A v<A A >>^A vA ^A vA ^A <vA A >^A

var ddd = []string{
	"<A",
	">A",
	"^A",
	"vA",
	"A",
}

func try() {
	// A 	1 -> 1
	// vA	6 -> 16

	//<vA>^A 6
	//<v<A>A>^AvA<^A>A 16
	//v<<A>A>^AvA<^A>A 16

	// 1
	// 44
	// 64
	// 63
	// 28
	// 46
	// 55
	// 28
	// 48
	// 47

	for i, a := range steps {
		fmt.Println(a)
		_ = i
		//fmt.Println(count(i, 3))
		seq := dirSeqAll(a)
		//seq = filter2(seq)
		for _, s := range seq {
			fmt.Printf("%s %d\n", s, len(s))

			seq2 := dirSeqAll(s)
			//seq2 = filter2(seq2)
			for _, s2 := range seq2 {
				fmt.Printf("  %s %d\n", s2, len(s2))

				seq3 := dirSeqAll(s2)
				seq3 = filter2(seq3)
				for _, s3 := range seq3 {
					fmt.Printf("    %s %d\n", s3, len(s3))

					//seq4 := filter2(dirSeqAll(s3))
					//for _, s4 := range seq4 {
					//	fmt.Printf("    %s %d\n", s4, len(s4))
					//}
				}
			}
		}
		fmt.Println("----------")
	}

}

/*
<v<A>A<A>>^AvAA<^A>A
v<<A>A<A>>^AvAA<^A>A
*/

func solve(p *Problem) {
	//s1, s2 := 0, 0
	//
	//s1a, s3, s4 := 0, 0, 0
	//
	////g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	//
	////s := p.ReadAll()
	//
	//try2()
	////os.Exit(1)
	//
	//for p.NextLine() {
	//	code := p.Line()
	//
	//	num := 0
	//	for i := 0; i < 4; i++ {
	//		ch := code[i]
	//		if ch != 'A' {
	//			num = num*10 + int(ch-'0')
	//		}
	//	}
	//	//p.Printf("%s %d", code, num)
	//
	//	//p.Printf("%v", numSeqAll(code))
	//	dirs1 := numSeqAll(code)
	//	//p.Printf("%v", dirs1)
	//
	//	dirs2 := []string(nil)
	//	for _, dir := range dirs1 {
	//		dirs2 = append(dirs2, dirSeqAll(dir)...)
	//	}
	//	//dirs2 = filter2(dirs2)
	//
	//	//2024/12/21 10:29:04 ✔ part two: 187223591936878
	//	//2024/12/21 10:29:04 ✔ part two: 475288764580252   <---- NOT RIGHT ANSWER !!!!
	//	//                                335879439503508
	//	//2024/12/21 10:29:04 ✔ part two: 1206575546418030
	//
	//	//2024/12/21 10:57:17 ✔ part two: 238078
	//	//2024/12/21 10:57:17 ✔ part two: 134180557577040 <<---- NENE
	//	//2024/12/21 10:57:17 ✔ part two: 335879439503508
	//
	//	// 331289147197568 <------------ NEEEE
	//
	//	//
	//	// 133590329564666 <--------- NENENENE
	//
	//	dirs3 := []string(nil)
	//	for _, dir := range dirs2 {
	//		dirs3 = append(dirs3, dirSeqAll(dir)...)
	//	}
	//	dirs3 = filter2(dirs3)
	//
	//	//dirs4 := []string(nil)
	//	//for _, dir := range dirs3 {
	//	//	dirs4 = append(dirs4, dirSeqAll(dir)...)
	//	//}
	//	//dirs4 = filter2(dirs4)
	//
	//	l1, l23, l24, l25 := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	//	for _, d := range dirs2 {
	//		l1 = min(l1, countAll(d, 1))
	//		l23 = min(l23, countAll(d, 23))
	//		l24 = min(l24, countAll(d, 24))
	//		l25 = min(l25, countAll(d, 25))
	//	}
	//
	//	//dirs4 := []string(nil)
	//	//for _, dir := range dirs3 {
	//	//	dirs4 = append(dirs4, dirSeqAll(dir)...)
	//	//}
	//	//dirs4 = filter2(dirs4)
	//
	//	l, best := math.MaxInt, ""
	//	for _, dir := range dirs3 {
	//		if len(dir) < l {
	//			l = len(dir)
	//			best = dir
	//		}
	//	}
	//
	//	_ = best
	//	p.Printf("%d %s", l, best)
	//
	//	s1 += num * l
	//	s1a += num * l1
	//	s2 += num * l23
	//	s3 += num * l24
	//	s4 += num * l25
	//
	//	//                               246810588779586
	//	//2024/12/21 12:56:21  part two: 132346787973836
	//	//2024/12/21 12:56:21  part two: 331289147197568
	//	//2024/12/21 12:56:21  part two: 829279671886492
	//}
	//
	//p.PartOne(s1)
	//p.PartOne(s1a)
	//
	//p.PartTwo(s2)
	//p.PartTwo(s3)
	//p.PartTwo(s4)
	//
	////os.Exit(1)

	s1, s2 := 0, 0
	for p.NextLine() {
		s1 += solve2(p.Line(), 2)
		s2 += solve2(p.Line(), 25)
	}
	p.PartOne(s1)
	p.PartOne(s2)
}

type cell struct {
	ch byte
}

func (c cell) String() string { return string([]byte{c.ch}) }

type grid struct {
	*Matrix[cell]
	A XY
}

//func (g grid) findSeqs(code string, at XY) []string {
//
//}

var (
	numpad, dpad grid
)

func appendRow(g grid, s string) {
	ParseVectorFunc(g.AppendRow(), s, func(b byte) cell { return cell{ch: b} })
}

func init() {
	numpad = grid{NewMatrix[cell](Rectangle(3, 4)), XY{2, 3}}
	appendRow(numpad, "789")
	appendRow(numpad, "456")
	appendRow(numpad, "123")
	appendRow(numpad, " 0A")

	dpad = grid{NewMatrix[cell](Rectangle(3, 2)), XY{2, 0}}
	appendRow(dpad, " ^A")
	appendRow(dpad, "<v>")
}

func solve2(code string, depth int) int {
	num := ParseInt(code[:len(code)-1])
	cache := map[key]int{}

	// vygeneruju vvsechny cesty po ciselne klavesnici
	// rekurze countPresses

	seqs := numSeqAll(code)
	best := math.MaxInt
	for _, seq := range seqs {
		best = min(best, countPresses(cache, seq, depth))
	}
	//fmt.Printf("%s %d\n", code, best)

	return num * best
}

type key struct {
	code  string
	depth int
}

func countPresses(cache map[key]int, code string, depth int) (result int) {
	if depth == 0 {
		return len(code)
	}

	k := key{code, depth}
	if s, ok := cache[k]; ok {
		return s
	}

	s, at := 0, XY{2, 0}
	for i := 0; i < len(code); i++ {
		seqs := dirSeqXY(code[i:i+1], at)
		best := math.MaxInt
		for _, seq := range seqs {
			best = min(best, countPresses(cache, seq, depth-1))
		}

		at = dirXY(code[i])
		s += best
	}

	cache[k] = s
	return s
}

func numSeqAll(code string) []string {
	return numSeqGen(code, XY{2, 3}, "")
}

func numSeqGen(code string, xy XY, seq string) []string {
	if len(code) == 0 {
		return []string{seq}
	}
	p := codeXY(code[0])
	if p == xy {
		return numSeqGen(code[1:], p, seq+"A")
	}

	s := []string(nil)
	d := p.Sub(xy)
	if d.X < 0 && xy != (XY{1, 3}) {
		s = append(s, numSeqGen(code, xy.Add(NegX), seq+"<")...)
	}
	if d.X > 0 {
		s = append(s, numSeqGen(code, xy.Add(PosX), seq+">")...)
	}
	if d.Y < 0 {
		s = append(s, numSeqGen(code, xy.Add(NegY), seq+"^")...)
	}
	if d.Y > 0 && xy != (XY{0, 2}) {
		s = append(s, numSeqGen(code, xy.Add(PosY), seq+"v")...)
	}
	return s
}

func dirSeqXY(code string, at XY) []string {
	return dirSeqGen(code, at, "")
}

func dirSeqAll(code string) []string {
	return dirSeqGen(code, XY{2, 0}, "")
}

func dirSeqGen(code string, xy XY, seq string) []string {
	if len(code) == 0 {
		return []string{seq}
	}
	p := dirXY(code[0])
	if p == xy {
		return dirSeqGen(code[1:], p, seq+"A")
	}

	s := []string(nil)
	d := p.Sub(xy)
	//h := false
	if d.X < 0 && xy != (XY{1, 0}) {
		s = append(s, dirSeqGen(code, xy.Add(NegX), seq+"<")...)
	}
	if d.X > 0 {
		s = append(s, dirSeqGen(code, xy.Add(PosX), seq+">")...)
	}
	if d.Y < 0 && xy != (XY{0, 1}) {
		s = append(s, dirSeqGen(code, xy.Add(NegY), seq+"^")...)
	}
	if d.Y > 0 {
		s = append(s, dirSeqGen(code, xy.Add(PosY), seq+"v")...)
	}
	return s
}
