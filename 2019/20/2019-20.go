package main

import (
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 20
	example = `

             Z L X W       C                 
             Z P Q B       K                 
  ###########.#.#.#.#######.###############  
  #...#.......#.#.......#.#.......#.#.#...#  
  ###.#.#.#.#.#.#.#.###.#.#.#######.#.#.###  
  #.#...#.#.#...#.#.#...#...#...#.#.......#  
  #.###.#######.###.###.#.###.###.#.#######  
  #...#.......#.#...#...#.............#...#  
  #.#########.#######.#.#######.#######.###  
  #...#.#    F       R I       Z    #.#.#.#  
  #.###.#    D       E C       H    #.#.#.#  
  #.#...#                           #...#.#  
  #.###.#                           #.###.#  
  #.#....OA                       WB..#.#..ZH
  #.###.#                           #.#.#.#  
CJ......#                           #.....#  
  #######                           #######  
  #.#....CK                         #......IC
  #.###.#                           #.###.#  
  #.....#                           #...#.#  
  ###.###                           #.#.#.#  
XF....#.#                         RF..#.#.#  
  #####.#                           #######  
  #......CJ                       NM..#...#  
  ###.#.#                           #.###.#  
RE....#.#                           #......RF
  ###.###        X   X       L      #.#.#.#  
  #.....#        F   Q       P      #.#.#.#  
  ###.###########.###.#######.#########.###  
  #.....#...#.....#.......#...#.....#.#...#  
  #####.#.###.#######.#######.###.###.#.#.#  
  #.......#.......#.#.#.#.#...#...#...#.#.#  
  #####.###.#####.#.#.#.#.###.###.#.###.###  
  #.......#.....#.#...#...............#...#  
  #############.#.#.###.###################  
               A O F   N                     
               A A D   M                     

`
)

func main() {
	Run(year, day, example, solve)
}

type portal struct {
	target XY
	outer  bool
}

type cell struct {
	ch   byte
	port *portal
	v1   bool
	v2   []bool
}

func (c *cell) visit(level int) bool {
	for level >= len(c.v2) {
		c.v2 = append(c.v2, false)
	}
	if !c.v2[level] {
		c.v2[level] = true
		return true
	}
	return false
}

type maze struct {
	*Matrix[cell]
	start, finish XY
}

func findCode(m *maze, c0 XY) (string, bool, bool) {
	if m.At(c0).ch != '.' {
		return "", false, false
	}

	for _, d := range HVDirs {
		c1 := c0.Add(d)
		c2 := c1.Add(d)
		if !m.Dim.HasInside(c2) {
			continue
		}
		ch1, ch2 := m.At(c1).ch, m.At(c2).ch
		if ch1 >= 'A' && ch1 <= 'Z' {
			outer := c2.X == 0 || c2.X == m.Dim.X-1 || c2.Y == 0 || c2.Y == m.Dim.Y-1
			if d.X == 1 || d.Y == 1 {
				return string([]byte{ch1, ch2}), outer, true
			} else {
				return string([]byte{ch2, ch1}), outer, true
			}
		}
	}
	return "", false, false
}

func parseMaze(p *Problem) *maze {
	m := &maze{Matrix: NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}
	for p.NextLine() {
		ParseVectorFunc(m.AppendRow(), p.Line(), func(b byte) cell { return cell{ch: b} })
	}
	portals := map[string]XY{}
	for a, c := range m.All() {
		if code, outer, ok := findCode(m, a); ok {
			switch code {
			case "AA":
				m.start = a
			case "ZZ":
				m.finish = a
			default:
				if b, ok := portals[code]; ok {
					c.port = &portal{target: b, outer: outer}
					m.At(b).port = &portal{target: a, outer: !outer}
				} else {
					portals[code] = a
				}
			}
		}
	}
	return m
}

func solve(p *Problem) {
	m := parseMaze(p)
	p.PartOne(solvePart1(m))
	p.PartTwo(solvePart2(m))
}

func solvePart1(m *maze) int {
	type qs struct {
		at  XY
		len int
	}

	q := Queue[qs]{}
	q.Push(qs{at: m.start})

	for q.Len() > 0 {
		s := q.Pop()
		if s.at == m.finish {
			return s.len
		}

		c := m.At(s.at)
		if p := c.port; p != nil {
			if !m.At(p.target).v1 {
				m.At(p.target).v1 = true
				q.Push(qs{at: p.target, len: s.len + 1})
			}
		}
		for _, d := range HVDirs {
			at := s.at.Add(d)
			if !m.Dim.HasInside(at) {
				continue
			}
			c2 := m.At(at)
			if !c2.v1 && c2.ch == '.' {
				c2.v1 = true
				q.Push(qs{at: at, len: s.len + 1})
			}
		}
	}
	panic("no path found")
}

func solvePart2(m *maze) int {
	type qs struct {
		at    XY
		len   int
		level int
	}

	q := Queue[qs]{}
	q.Push(qs{at: m.start})

	for q.Len() > 0 {
		s := q.Pop()
		if s.level == 0 && s.at == m.finish {
			return s.len
		}

		c := m.At(s.at)
		if p := c.port; p != nil {
			level := s.level + 1
			if p.outer {
				level = s.level - 1
			}
			if level >= 0 && m.At(p.target).visit(level) {
				q.Push(qs{at: p.target, len: s.len + 1, level: level})
			}
		}
		for _, d := range HVDirs {
			at := s.at.Add(d)
			if !m.Dim.HasInside(at) {
				continue
			}
			c2 := m.At(at)
			if c2.ch == '.' && c2.visit(s.level) {
				q.Push(qs{at: at, len: s.len + 1, level: s.level})
			}
		}
	}
	panic("no path found")
}
