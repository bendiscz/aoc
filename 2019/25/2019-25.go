package main

import (
	"bufio"
	"fmt"
	"github.com/bendiscz/aoc/2019/intcode"
	. "github.com/bendiscz/aoc/aoc"
	"os"
	"strings"
)

const (
	year    = 2019
	day     = 25
	example = ``
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	prog := intcode.Parse(p)
	comp := prog.Exec()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		echo(comp)
		if comp.State() == intcode.Exited || !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "crack ") {
			crack(comp, strings.TrimPrefix(line, "crack "))
		} else {
			comp.WriteLine(line)
		}
	}
}

func echo(comp *intcode.Computer) {
	for {
		if line, ok := comp.ReadLine(); ok {
			fmt.Printf("%s\n", line)
		} else {
			break
		}
	}
}

func crack(comp *intcode.Computer, dir string) {
	items := []string(nil)
	comp.WriteLine("inv")
	for {
		line, ok := comp.ReadLine()
		if !ok {
			break
		}
		if strings.HasPrefix(line, "- ") {
			items = append(items, strings.TrimPrefix(line, "- "))
		}
	}

	mask := 1<<len(items) - 1

	for {
		comp.WriteLine(dir)
		echo(comp)
		if comp.State() == intcode.Exited || mask == 0 {
			break
		}

		m := mask - 1
		for i, item := range items {
			b := 1 << i
			if mask&b != 0 && m&b == 0 {
				comp.WriteLine("drop " + item)
			}
			if mask&b == 0 && m&b != 0 {
				comp.WriteLine("take " + item)
			}
		}
		mask = m
	}
}
