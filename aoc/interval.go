package aoc

import (
	"cmp"
	"fmt"
	"sort"
)

type Interval[T cmp.Ordered] struct {
	X1, X2 T
}

func (i Interval[T]) NonEmpty() bool {
	return i.X2 > i.X1
}

func (i Interval[T]) Set() IntervalSet[T] {
	s := [...]Interval[T]{i}
	return s[:]
}

func (i Interval[T]) String() string {
	return fmt.Sprintf("[%v; %v]", i.X1, i.X2)
}

type IntervalSet[T cmp.Ordered] []Interval[T]

func NewIntervalSet[T cmp.Ordered](x1, x2 T) IntervalSet[T] {
	return Interval[T]{x1, x2}.Set()
}

func MergeIntervals[T cmp.Ordered](intervals ...Interval[T]) IntervalSet[T] {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].X1 == intervals[j].X1 {
			return intervals[i].X2 < intervals[j].X2
		} else {
			return intervals[i].X1 < intervals[j].X1
		}
	})

	r := IntervalSet[T](nil)
	c := intervals[0]
	for _, i := range intervals[1:] {
		if i.X1 > c.X2 {
			if c.NonEmpty() {
				r = append(r, c)
			}
			c = i
		} else {
			c.X2 = max(c.X2, i.X2)
		}
	}
	if c.NonEmpty() {
		r = append(r, c)
	}
	return r
}
