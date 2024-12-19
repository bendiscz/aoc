package aoc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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
	desc     string
	exam     bool
	data     io.ReadSeeker
	lines    *bufio.Scanner
	line     string
	nextLine *string
	silent   bool
}

func Example(s string) *Problem {
	return &Problem{
		desc: "example",
		exam: true,
		data: strings.NewReader(strings.Trim(s, "\n")),
	}
}

func Day(year, day int) *Problem {
	dir, err := getProjectRoot()
	if err != nil {
		log.Panicf("failed to detect project root: %v", err)
	}

	f := fetcher{dir: dir}
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
	p.nextLine = nil
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

func (p *Problem) PeekLine() string {
	if p.nextLine != nil {
		return *p.nextLine
	}

	if line, ok := p.readNextLine(); ok {
		p.nextLine = &line
		return line
	} else {
		return ""
	}
}

func (p *Problem) NextLine() bool {
	if p.nextLine != nil {
		p.line = *p.nextLine
		p.nextLine = nil
		return true
	}

	line, ok := p.readNextLine()
	p.line = line
	return ok
}

func (p *Problem) readNextLine() (string, bool) {
	s := p.scanner()
	if ok := s.Scan(); !ok {
		return "", false
	}

	return s.Text(), true
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

func getProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get current working directory: %v", err)
	}
	for {
		_, err = os.Stat(filepath.Join(dir, "go.mod"))
		if err == nil {
			return dir, nil
		}
		if !os.IsNotExist(err) {
			return "", fmt.Errorf("failed to stat go.mod: %v", err)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find go.mod in working directory and its parents")
		}
		dir = parent
	}
}
