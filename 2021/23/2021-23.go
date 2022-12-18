package main

import (
	"fmt"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2021
	day     = 23
	example = `

#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########

`
)

func main() {
	Run(year, day, example, solve)
}

const (
	None pod = iota
	A
	B
	C
	D
	Block
)

var podAttrs = [...]struct {
	string     string
	cost, home int
	key        key
}{
	None:  {".", 0, 0, 0},
	A:     {"A", 1, 0, 1},
	B:     {"B", 10, 1, 2},
	C:     {"C", 100, 2, 3},
	D:     {"D", 1000, 3, 4},
	Block: {"#", 0, 0, 0},
}

type pod byte

func (p pod) String() string { return podAttrs[p].string }
func (p pod) cost() int      { return podAttrs[p].cost }
func (p pod) home() int      { return podAttrs[p].home }
func (p pod) key() key       { return podAttrs[p].key }

type room struct {
	id   int
	cell [4]pod
}

func (r *room) hallEntry() int {
	return 2 * (r.id + 1)
}

func (r *room) top() pod {
	for i := len(r.cell) - 1; i >= 0; i-- {
		p := r.cell[i]
		if p == Block {
			break
		}
		if p != None {
			return p
		}
	}
	return None
}

func (r *room) clean() bool {
	for _, c := range r.cell {
		if c != None && c != Block && c.home() != r.id {
			return false
		}
	}
	return true
}

func (r *room) take() (pod, int, bool) {
	for i := 1; i <= len(r.cell); i++ {
		j := len(r.cell) - i
		p := r.cell[j]
		if p == Block {
			break
		}
		if p != None {
			r.cell[j] = None
			return p, i, true
		}
	}
	return None, 0, false
}

func (r *room) push(p pod) (int, bool) {
	if !r.clean() || p.home() != r.id {
		return 0, false
	}
	for i := 0; i < len(r.cell); i++ {
		if r.cell[i] == None {
			r.cell[i] = p
			return len(r.cell) - i, true
		}
	}
	return 0, false
}

type hall struct {
	cell [11]pod
}

func (h *hall) empty() bool {
	for _, c := range h.cell {
		if c != None {
			return false
		}
	}
	return true
}

func (h *hall) isEntry(i int) bool {
	return i == 2 || i == 4 || i == 6 || i == 8
}

func (h *hall) free(from, to int) bool {
	var dir int
	if to > from {
		dir = 1
	} else if to < from {
		dir = -1
	} else {
		return true
	}

	for i := from + dir; i != to; i += dir {
		if h.cell[i] != None {
			return false
		}
	}
	return true
}

type burrow struct {
	hall hall
	room [4]room
	cost int
	step int
	prev *burrow
}

type key uint64

func (b *burrow) key() (k key) {
	for _, c := range b.hall.cell {
		k = k*5 + c.key()
	}
	for ri := 0; ri < len(b.room); ri++ {
		for _, c := range b.room[ri].cell {
			k = k*5 + c.key()
		}
	}
	return k
}

func (b *burrow) next() *burrow {
	n := *b
	n.step++
	n.prev = b
	return &n
}

func (b *burrow) done() bool {
	if !b.hall.empty() {
		return false
	}
	for i := 0; i < len(b.room); i++ {
		if !b.room[i].clean() {
			return false
		}
	}
	return true
}

func (b *burrow) steps() (s []*burrow) {
	for hi := 0; hi < len(b.hall.cell); hi++ {
		if n := b.hall2room(hi); n != nil {
			s = append(s, n)
		}
	}
	if len(s) > 0 {
		return
	}

	for ri := 0; ri < len(b.room); ri++ {
		if n := b.room2room(ri); n != nil {
			s = append(s, n)
		}
	}
	if len(s) > 0 {
		return
	}

	for ri := 0; ri < len(b.room); ri++ {
		for hi := b.room[ri].hallEntry() - 1; hi >= 0 && b.hall.cell[hi] == None; hi-- {
			if n := b.room2hall(ri, hi); n != nil {
				s = append(s, n)
			}
		}
		for hi := b.room[ri].hallEntry() + 1; hi < len(b.hall.cell) && b.hall.cell[hi] == None; hi++ {
			if n := b.room2hall(ri, hi); n != nil {
				s = append(s, n)
			}
		}
	}

	return
}

func (b *burrow) hall2room(hi int) *burrow {
	p := b.hall.cell[hi]
	if p == None {
		return nil
	}

	r := &b.room[p.home()]
	if !r.clean() {
		return nil
	}

	if !b.hall.free(hi, r.hallEntry()) {
		return nil
	}

	n := b.next()
	r = &n.room[p.home()]
	if d, ok := r.push(p); ok {
		dist := Abs(r.hallEntry()-hi) + d
		n.hall.cell[hi] = None
		n.cost += dist * p.cost()
		return n
	}
	return nil
}

func (b *burrow) room2room(ri int) *burrow {
	p := b.room[ri].top()
	if p == None || p.home() == ri {
		return nil
	}

	if !b.room[p.home()].clean() {
		return nil
	}

	e1, e2 := b.room[ri].hallEntry(), b.room[p.home()].hallEntry()
	if !b.hall.free(e1, e2) {
		return nil
	}

	n := b.next()
	_, d1, _ := n.room[ri].take()
	if d2, ok := n.room[p.home()].push(p); ok {
		dist := Abs(e2-e1) + d1 + d2
		n.cost += dist * p.cost()
		return n
	}
	return nil
}

func (b *burrow) room2hall(ri, hi int) *burrow {
	// expects that the room is not clean and hall path is free
	if b.hall.isEntry(hi) {
		return nil
	}

	n := b.next()
	p, d, ok := n.room[ri].take()
	if !ok {
		return nil
	}

	n.hall.cell[hi] = p
	d += Abs(hi - n.room[ri].hallEntry())
	n.cost += d * p.cost()
	return n
}

func (b *burrow) print() {
	fmt.Printf("#%d cost: %d\n#############\n#", b.step, b.cost)
	for _, c := range b.hall.cell {
		fmt.Printf("%v", c)
	}
	fmt.Printf("#\n")

	for i := 3; i >= 0; i-- {
		if i == 3 {
			fmt.Printf("###")
		} else {
			fmt.Printf("  #")
		}
		for ri := 0; ri < len(b.room); ri++ {
			fmt.Printf("%v#", b.room[ri].cell[i])
		}
		if i == 3 {
			fmt.Printf("##\n")
		} else {
			fmt.Printf("  \n")
		}
	}
	fmt.Printf("  #########\n")
}

func (b *burrow) printPath() {
	if b.prev != nil {
		b.prev.printPath()
	}
	b.print()
	fmt.Println()
}

func emptyBurrow() *burrow {
	b := &burrow{}
	for i := 0; i < len(b.room); i++ {
		b.room[i].id = i
	}
	return b
}

func parseBurrow(s string) *burrow {
	var r []pod
	for i, p := 0, 0; i < 8; i++ {
		for s[p] < 'A' {
			p++
		}
		r = append(r, pod(s[p]-'A'+1))
		p++
	}

	b := emptyBurrow()
	b.room[0].cell = [...]pod{Block, Block, r[4], r[0]}
	b.room[1].cell = [...]pod{Block, Block, r[5], r[1]}
	b.room[2].cell = [...]pod{Block, Block, r[6], r[2]}
	b.room[3].cell = [...]pod{Block, Block, r[7], r[3]}
	return b
}

var folded = [4][2]pod{{D, D}, {B, C}, {A, B}, {C, A}}

func (b *burrow) unfold() *burrow {
	n := *b
	for ri := 0; ri < len(n.room); ri++ {
		cell := &n.room[ri].cell
		cell[0] = cell[2]
		cell[1] = folded[ri][0]
		cell[2] = folded[ri][1]
	}
	return &n
}

func search(start *burrow) *burrow {
	q := NewHeap[*burrow](func(b1, b2 *burrow) bool { return b1.cost < b2.cost })
	q.Push(start)
	v := map[key]struct{}{}

	for q.Len() > 0 {
		b := q.Pop()
		k := b.key()
		if _, ok := v[k]; ok {
			continue
		}
		v[k] = struct{}{}

		if b.done() {
			return b
		}

		for _, n := range b.steps() {
			q.Push(n)
		}
	}
	return nil
}

func solve(p *Problem) {
	s1 := parseBurrow(p.ReadAll())
	b1 := search(s1)
	p.PartOne(b1.cost)

	s2 := s1.unfold()
	b2 := search(s2)
	p.PartTwo(b2.cost)
}
