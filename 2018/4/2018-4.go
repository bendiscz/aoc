package main

import (
	"strings"
	"time"

	. "github.com/bendiscz/aoc/aoc"
	"golang.org/x/exp/slices"
)

const (
	year    = 2018
	day     = 4
	example = `

[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up

`
)

func main() {
	Run(year, day, example, solve)
}

type entry struct {
	t   time.Time
	msg string
}

type guard struct {
	id       int
	sleeping int
	minutes  [60]int
}

func solve(p *Problem) {
	log := []entry(nil)
	for p.NextLine() {
		f := p.Parse(`^\[(\d+)-(\d+)-(\d+) (\d+):(\d+)\] (.*)$`)
		e := entry{
			t:   time.Date(ParseInt(f[1]), time.Month(ParseInt(f[2])), ParseInt(f[3]), ParseInt(f[4]), ParseInt(f[5]), 0, 0, time.UTC),
			msg: f[6],
		}
		log = append(log, e)
	}

	slices.SortFunc(log, func(a, b entry) int {
		return a.t.Compare(b.t)
	})

	guards, g, sleepingSince := map[int]*guard{}, (*guard)(nil), time.Time{}
	for _, e := range log {
		switch {
		case strings.HasPrefix(e.msg, "Guard"):
			id := ParseInts(e.msg)[0]
			g = guards[id]
			if g == nil {
				g = &guard{id: id}
				guards[id] = g
			}

		case e.msg == "falls asleep":
			sleepingSince = e.t

		case e.msg == "wakes up":
			g.sleeping += int(e.t.Sub(sleepingSince).Minutes())
			for t := sleepingSince; t.Before(e.t); t = t.Add(time.Minute) {
				g.minutes[t.Minute()]++
			}
		}
	}

	g1, g2, m2, n2 := (*guard)(nil), (*guard)(nil), 0, 0
	for _, g = range guards {
		if g1 == nil || g.sleeping > g1.sleeping {
			g1 = g
		}

		m, n := maxMinute(g)
		if g2 == nil || n > n2 {
			g2 = g
			m2, n2 = m, n
		}
	}

	m1, _ := maxMinute(g1)
	p.PartOne(g1.id * m1)
	p.PartTwo(g2.id * m2)
}

func maxMinute(g *guard) (int, int) {
	m, n := 0, 0
	for i := 0; i < 60; i++ {
		if g.minutes[i] > n {
			n = g.minutes[i]
			m = i
		}
	}
	return m, n
}
