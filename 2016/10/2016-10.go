package main

//import (
//	"fmt"
//	"strconv"
//
//	. "github.com/bendiscz/aoc/aoc"
//)
//
//const (
//	year    = 2016
//	day     = 10
//	example = `
//
//value 5 goes to bot 2
//bot 2 gives low to bot 1 and high to bot 0
//value 3 goes to bot 1
//bot 1 gives low to output 1 and high to bot 0
//bot 0 gives low to output 2 and high to output 0
//value 2 goes to bot 2
//
//`
//)
//
//func main() {
//	Run(year, day, example, solve)
//}
//
//type chip int
//
//func (c *chip) String() string {
//	if c == nil {
//		return "nil"
//	} else {
//		return strconv.Itoa(int(*c))
//	}
//}
//
//type inputBin struct {
//	chip   chip
//	target target
//}
//
//type outputBin struct {
//	id   int
//	chip chip
//}
//
//func (out *outputBin) String() string {
//	return fmt.Sprintf("{%d chip=%v}", out.id, out.chip)
//}
//
//type target interface {
//	accept(chip) bool
//}
//
//func makeTarget(kind, id string) target {
//	switch kind {
//	case "bot":
//		return botTarget(aoc.ParseInt(id))
//	case "output":
//		return outputTarget(aoc.ParseInt(id))
//	default:
//		panic("unknown target type: " + kind)
//	}
//}
//
//type outputTarget int
//
//func (t outputTarget) accept(c chip) bool {
//	id := int(t)
//	if _, ok := outputs[id]; ok {
//		return false
//	}
//
//	outputs[id] = &outputBin{
//		id:   id,
//		chip: c,
//	}
//	return true
//}
//
//func (t outputTarget) String() string {
//	return fmt.Sprintf("output(%d)", int(t))
//}
//
//type botTarget int
//
//func (t botTarget) accept(c chip) bool {
//	b, ok := bots[int(t)]
//	if !ok {
//		return false
//	}
//	return b.accept(c)
//}
//
//func (t botTarget) String() string {
//	return fmt.Sprintf("bot(%d)", int(t))
//}
//
//type bot struct {
//	id           int
//	chip1, chip2 *chip
//	low, high    target
//}
//
//func newBot(id int, low, high target) *bot {
//	return &bot{
//		id:   id,
//		low:  low,
//		high: high,
//	}
//}
//
//func (b *bot) complete() bool {
//	return b.chip1 != nil && b.chip2 != nil
//}
//
//func (b *bot) accept(c chip) bool {
//	if b.chip1 == nil {
//		b.chip1 = &c
//		return true
//	}
//	if b.chip2 == nil {
//		b.chip2 = &c
//		return true
//	}
//	return false
//}
//
//func (b *bot) work() bool {
//	if !b.complete() {
//		return false
//	}
//
//	if int(*b.chip1) > int(*b.chip2) {
//		b.chip1, b.chip2 = b.chip2, b.chip1
//	}
//
//	if *b.chip1 == chip(17) && *b.chip2 == chip(61) {
//		result = b
//	}
//
//	lowAccepted := b.low.accept(*b.chip1)
//	highAccepted := b.high.accept(*b.chip2)
//
//	fmt.Printf("bot %d gives %v to %v and %v to %v\n", b.id, *b.chip1, b.low, *b.chip2, b.high)
//
//	if lowAccepted {
//		b.chip1 = nil
//	}
//	if highAccepted {
//		b.chip2 = nil
//	}
//
//	return lowAccepted || highAccepted
//}
//
//func (b *bot) compares(c1, c2 chip) bool {
//	if !b.complete() {
//		return false
//	}
//	return c1 == *b.chip1 && c2 == *b.chip2 || c1 == *b.chip2 && c2 == *b.chip1
//}
//
//func (b *bot) String() string {
//	return fmt.Sprintf("{%d chips=[%v;%v] low=%v high=%v}", b.id, b.chip1, b.chip2, b.low, b.high)
//}
//
//func solve(p *Problem) {
//	for p.NextLine() {
//		p.Printf("%s", p.Line())
//	}
//
//	p.PartOne("TODO")
//
//	p.PartTwo("TODO")
//}
