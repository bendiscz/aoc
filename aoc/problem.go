package aoc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"time"
)

var (
	reCache = map[string]*regexp.Regexp{}
)

type LineReader interface {
	ReadAll() string
	ReadLine() string
	ReadLineBytes() []byte
	SkipLines(n int)
	NextLine() bool
	Line() string
	LineBytes() []byte
	Parse(pattern string) []string
	Scanf(format string, v ...any)
}

type Problem struct {
	desc   string
	exam   bool
	data   io.ReadSeeker
	lines  *bufio.Scanner
	line   string
	silent bool
}

func Example(s string) *Problem {
	return &Problem{
		desc: "example",
		exam: true,
		data: strings.NewReader(strings.Trim(s, "\n")),
	}
}

func Day(year, day int) *Problem {
	f := fetcher{}
	return &Problem{
		desc: "actual",
		data: f.fetchInput(year, day),
	}
}

func Run(year, day int, example string, solver func(*Problem)) {
	log.Printf("🎄 Advent of Code %d/%d 🎄", year, day)
	if len(example) > 0 {
		log.Printf("---")
		Example(example).Run(solver)
	}
	log.Printf("---")
	Day(year, day).Run(solver)
}

func (p *Problem) Run(solver func(*Problem)) {
	p.Printf("🚀 solving problem for %s input", p.desc)
	st := time.Now()
	solver(p)
	p.Printf("⏱ problem solved in %v", time.Since(st))
}

func (p *Problem) Example() bool {
	return p.exam
}

func (p *Problem) Silence() {
	p.silent = true
}

func (p *Problem) Reset() {
	_, err := p.data.Seek(0, io.SeekStart)
	if err != nil {
		log.Panicf("failed to reset input: %v", err)
	}

	p.lines = nil
}

func (p *Problem) ReadAll() string {
	if p.lines != nil {
		log.Panicf("cannot ReadAll after scanning lines")
	}

	b, err := io.ReadAll(p.data)
	if err != nil {
		log.Panicf("failed to read input data: %v", err)
	}
	return string(b)
}

func (p *Problem) SkipLines(n int) {
	for i := 0; i < n; i++ {
		p.NextLine()
	}
}

func (p *Problem) ReadLine() string {
	p.NextLine()
	return p.Line()
}

func (p *Problem) ReadLineBytes() []byte {
	p.NextLine()
	return p.LineBytes()
}

func (p *Problem) NextLine() bool {
	s := p.scanner()
	if ok := s.Scan(); !ok {
		p.line = ""
		return false
	}

	p.line = s.Text()
	return true
}

func (p *Problem) Line() string {
	return p.line
}

func (p *Problem) LineBytes() []byte {
	return []byte(p.line)
}

func (p *Problem) Parse(pattern string) []string {
	re, ok := reCache[pattern]
	if !ok {
		re = regexp.MustCompile(pattern)
		reCache[pattern] = re
	}
	return re.FindStringSubmatch(p.Line())
}

func (p *Problem) Scanf(format string, v ...any) {
	_, _ = fmt.Sscanf(p.Line(), format, v...)
}

func (p *Problem) Printf(format string, v ...any) {
	if p.silent {
		return
	}
	log.Printf(format, v...)
}

func (p *Problem) PartOne(v any) {
	p.Printf("✔ part one: %v", v)
}

func (p *Problem) PartTwo(v any) {
	p.Printf("✔ part two: %v", v)
}

func (p *Problem) scanner() *bufio.Scanner {
	if p.lines == nil {
		p.lines = bufio.NewScanner(p.data)
	}
	return p.lines
}

type blockLineReader struct {
	*Problem
	finished bool
}

func (r *blockLineReader) readLine() string {
	if r.finished {
		return ""
	}

	line := r.ReadLine()
	if line == "" {
		r.finished = true
	}

	return line
}

func (r *blockLineReader) ReadAll() string {
	all := strings.Builder{}
	for line := r.readLine(); line != ""; line = r.readLine() {
		all.WriteString(line)
	}
	return all.String()
}

func (r *blockLineReader) ReadLine() string {
	return r.readLine()
}

func (r *blockLineReader) ReadLineBytes() []byte {
	return []byte(r.ReadLine())
}

//x	ReadAll() string
//x	ReadLine() string
//x	ReadLineBytes() []byte
//	SkipLines(n int)
//	NextLine() bool
//	Line() string
//	LineBytes() []byte
//	Parse(pattern string) []string
//	Scanf(format string, v ...any)
