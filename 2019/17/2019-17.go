package main

import (
	"github.com/bendiscz/aoc/2019/intcode"
	. "github.com/bendiscz/aoc/aoc"
	"slices"
	"strconv"
	"strings"
)

const (
	year    = 2019
	day     = 17
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

func (c cell) String() string { return string([]byte{c.ch}) }

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	prog := intcode.Parse(p)

	c := prog.Exec()
	line, ok := c.ReadLine()
	g := grid{NewMatrix[cell](Rectangle(len(line), 0))}
	start := XY0
	for {
		if x := strings.IndexAny(line, "^v<>"); x >= 0 {
			start = XY{X: x, Y: g.Dim.Y}
		}
		ParseVectorFunc(g.AppendRow(), line, func(b byte) cell { return cell{ch: b} })
		line, ok = c.ReadLine()
		if !ok || len(line) == 0 {
			break
		}
	}

	//PrintGrid(g)
	//fmt.Printf("start: %v\n", start)

	p.PartOne(countCrosses(g))

	path := findPath(g, start)
	cmd := makeCommands(path)

	c = prog.Exec()
	c.Poke(0, 2)

	c.WriteLine(strings.Join(cmd.main, ","))
	c.WriteLine(strings.Join(cmd.fn[0], ","))
	c.WriteLine(strings.Join(cmd.fn[1], ","))
	c.WriteLine(strings.Join(cmd.fn[2], ","))
	c.WriteLine("n")

	for {
		line, ok = c.ReadLine()
		if !ok {
			break
		}
		p.Printf("%s", line)
	}

	result, _ := c.ReadInt()
	p.PartTwo(result)
}

func countCrosses(g grid) int {
	s := 0
	for xy, c := range g.All() {
		x := 0
		if c.ch != '.' {
			x++
		}
		for _, d := range HVDirs {
			n := xy.Add(d)
			if g.Dim.HasInside(n) && g.At(n).ch != '.' {
				x++
			}
		}
		if x == 5 {
			s += xy.X * xy.Y
			c.ch = 'O'
		}
	}
	return s
}

func findPath(g grid, start XY) []string {
	var dir XY
	switch g.At(start).ch {
	case '^':
		dir = NegY
	case 'v':
		dir = PosY
	case '<':
		dir = NegX
	case '>':
		dir = PosX
	}

	prog := []string(nil)
	c, l := start, 0
	for {
		c2 := c.Add(dir)
		if g.Dim.HasInside(c2) && g.At(c2).ch != '.' {
			c = c2
			g.At(c).ch = '*'
			l++
		} else {
			if l > 0 {
				prog = append(prog, strconv.Itoa(l))
				l = 0
			}

			right := XY{-dir.Y, dir.X}
			left := right.Neg()
			if c2 = c.Add(right); g.Dim.HasInside(c2) && g.At(c2).ch != '.' {
				prog = append(prog, "R")
				dir = right
			} else if c2 = c.Add(left); g.Dim.HasInside(c2) && g.At(c2).ch != '.' {
				prog = append(prog, "L")
				dir = left
			} else {
				break
			}
		}
	}
	return prog
}

const N = 20

type commands struct {
	main []string
	fn   [3][]string
}

func makeCommands(main []string) commands {
	if c, ok := findCommands(main, nil); ok && len(c.main)*2-1 <= N {
		return c
	}
	panic("unable to create valid commands")
}

func findCommands(main []string, fn [][]string) (commands, bool) {
	bi := -1
	for i, cmd := range main {
		if !isFn(cmd) {
			bi = i
			break
		}
	}

	if bi == -1 {
		c := commands{main: main}
		copy(c.fn[:], fn)
		return c, true
	}

	if len(fn) == 3 {
		return commands{}, false
	}

	bj, bl := bi, -1
	for bj < len(main) && !isFn(main[bj]) {
		bl = bl + len(main[bj]) + 1
		if bl > N {
			break
		}
		bj++
	}

	fn2 := make([][]string, len(fn)+1)
	copy(fn2, fn)
	for i := bj; i > bi+1; i-- {
		fn2[len(fn)] = slices.Clone(main[bi:bj])
		main2 := substituteFn(main, fn2[len(fn)], string([]byte{'A' + byte(len(fn))}))
		if c, ok := findCommands(main2, fn2); ok {
			return c, true
		}
	}

	return commands{}, false
}

func isFn(cmd string) bool {
	return len(cmd) == 1 && cmd[0] >= 'A' && cmd[0] <= 'C'
}

func substituteFn(main, fn []string, name string) []string {
	main = slices.Clone(main)
	i := 0
	for i < len(main) {
		if slices.Equal(main[i:i+len(fn)], fn) {
			main = slices.Replace(main, i, i+len(fn), name)
		}
		i++
	}
	return main
}
