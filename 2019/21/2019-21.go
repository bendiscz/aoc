package main

import (
	"github.com/bendiscz/aoc/2019/intcode"
	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2019
	day     = 21
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

type cell struct {
	ch byte
}

type grid struct {
	*Matrix[cell]
}

func solve(p *Problem) {
	prog := intcode.Parse(p)
	comp := prog.Exec()
	echo(p, comp)

	// ABCD
	// #### WALK false
	// ###. WALK false
	// ##.# JUMP true
	// ##.. WALK false
	// #.## JUMP true
	// #.#. WALK false
	// #..# JUMP true
	// #... WALK false
	// .### JUMP true
	// .##. FALL false
	// .#.# JUMP true
	// .#.. FALL false
	// ..## JUMP true
	// ..#. FALL false
	// ...# JUMP true
	// .... FALL false
	//
	// D and !(A and B and C and D)

	comp.WriteLine("OR A T")
	comp.WriteLine("AND B T")
	comp.WriteLine("AND C T")
	comp.WriteLine("AND D T")
	comp.WriteLine("NOT T J")
	comp.WriteLine("AND D J")
	comp.WriteLine("WALK")

	echo(p, comp)

	if damage, ok := comp.ReadInt(); ok {
		p.PartOne(damage)
	}

	comp = prog.Exec()
	echo(p, comp)

	// do not JUMP if
	// ABCDEFGHI
	// #__#.__._  and !(A and D and !E and !H) ... and (!(A and D) or E or H)

	comp.WriteLine("OR A T")
	comp.WriteLine("AND B T")
	comp.WriteLine("AND C T")
	comp.WriteLine("AND D T")
	comp.WriteLine("NOT T J")
	comp.WriteLine("AND D J")

	comp.WriteLine("NOT A T") // clear T
	comp.WriteLine("AND A T")
	comp.WriteLine("OR A T")
	comp.WriteLine("AND D T")
	comp.WriteLine("NOT T T")
	comp.WriteLine("OR E T")
	comp.WriteLine("OR H T")
	comp.WriteLine("AND T J")

	comp.WriteLine("RUN")

	echo(p, comp)

	if damage, ok := comp.ReadInt(); ok {
		p.PartTwo(damage)
	}
}

func echo(p *Problem, c *intcode.Computer) {
	for {
		line, ok := c.ReadLine()
		if !ok {
			break
		}
		p.Printf("> %s", line)
	}
}
