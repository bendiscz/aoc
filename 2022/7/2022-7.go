package main

import (
	"math"
	"strings"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2022
	day     = 7
	example = `

$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`
)

func main() {
	Run(year, day, example, solve)
}

type dir struct {
	size  int
	total int
	dirs  map[string]*dir
}

func mkdir() *dir {
	return &dir{dirs: map[string]*dir{}}
}

func (d *dir) countTotal() {
	d.total = d.size
	for _, cd := range d.dirs {
		cd.countTotal()
		d.total += cd.total
	}
}

func (d *dir) traverse(fn func(d *dir)) {
	for _, cd := range d.dirs {
		cd.traverse(fn)
	}
	fn(d)
}

func solve(p *Problem) {
	root := mkdir()
	cwd := Stack[*dir]{}
	cwd.Push(root)

	for p.NextLine() {
		line := p.Line()
		if _, cmd, ok := strings.Cut(line, "$ "); ok {
			switch cmd {
			case "ls":
				cwd.Top().size = 0

			case "cd /":
				cwd.Clear()
				cwd.Push(root)

			case "cd ..":
				cwd.Pop()

			default:
				d := mkdir()
				cwd.Top().dirs[cmd[3:]] = d
				cwd.Push(d)
			}
		} else if n, _, ok := strings.Cut(line, " "); ok && n != "dir" {
			cwd.Top().size += ParseInt(n)
		}
	}

	root.countTotal()

	total100000 := 0
	root.traverse(func(d *dir) {
		if d.total <= 100000 {
			total100000 += d.total
		}
	})
	p.PartOne(total100000)

	min, deleted := root.total-40000000, math.MaxInt
	root.traverse(func(d *dir) {
		if d.total >= min && d.total < deleted {
			deleted = d.total
		}
	})
	p.PartTwo(deleted)
}
