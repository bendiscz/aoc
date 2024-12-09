package aoc

var SET = struct{}{}

type Set[T comparable] map[T]struct{}

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
