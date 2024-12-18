package main

import (
	"github.com/bendiscz/aoc/2019/intcode"
	. "github.com/bendiscz/aoc/aoc"
	"sync"
)

const (
	year    = 2019
	day     = 7
	example = `

3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	prog := intcode.Parse(p)

	for setting := range Permutations([]int{0, 1, 2, 3, 4}) {
		signal := amplify1(prog, setting)
		s1 = max(s1, signal)
	}
	p.PartOne(s1)

	for setting := range Permutations([]int{5, 6, 7, 8, 9}) {
		signal := amplify2(prog, setting)
		s2 = max(s2, signal)
	}
	p.PartTwo(s2)
}

func amplify1(prog *intcode.Program, setting []int) int {
	signal := 0
	for _, phase := range setting {
		signal = prog.Exec().ReadWrite(phase, signal)[0]
	}
	return signal
}

func amplify2(prog *intcode.Program, setting []int) int {
	var channels []chan int
	wg := &sync.WaitGroup{}
	for _, phase := range setting {
		ch := make(chan int, 1)
		ch <- phase
		channels = append(channels, ch)
	}

	for i := range setting {
		c := prog.Exec()
		in, out := channels[i], channels[(i+1)%len(channels)]
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(out)
			for c.State() != intcode.Exited {
				value, ok := <-in
				if !ok {
					return
				}
				c.WriteInt(value)
				if value, ok = c.ReadInt(); ok {
					out <- value
				}
			}
		}()
	}

	channels[0] <- 0
	wg.Wait()
	return <-channels[0]
}
