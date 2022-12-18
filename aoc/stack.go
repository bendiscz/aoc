package aoc

import "log"

type Stack[T any] struct {
	entries []T
}

func (s *Stack[T]) Len() int {
	return len(s.entries)
}

func (s *Stack[T]) Clear() {
	s.entries = nil
}

func (s *Stack[T]) Top() T {
	l := s.Len()
	if l == 0 {
		log.Panicf("stack underflow")
	}
	return s.entries[l-1]
}

func (s *Stack[T]) Push(value T) {
	s.entries = append(s.entries, value)
}

func (s *Stack[T]) Pop() T {
	value := s.Top()
	s.entries = s.entries[:s.Len()-1]
	return value
}
