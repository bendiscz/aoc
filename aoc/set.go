package aoc

import "maps"

var SET = struct{}{}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Empty() bool {
	return len(s) == 0
}

func (s Set[T]) Clone() Set[T] {
	return maps.Clone(s)
}

func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) Intersect(s2 Set[T]) Set[T] {
	r := Set[T]{}
	for e := range s {
		if s2.Contains(e) {
			r[e] = SET
		}
	}
	return r
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	r := s.Clone()
	for e := range s2 {
		r[e] = SET
	}
	return r
}
