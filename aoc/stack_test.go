package aoc

import "testing"

func TestStack(t *testing.T) {
	type entry struct {
		x int
	}
	s := Stack[entry]{}

	if want, got := 0, s.Len(); got != want {
		t.Errorf("Len want: %v, got: %v", want, got)
	}

	assertPanic(t, "stack underflow", func() { _ = s.Top() })
	assertPanic(t, "stack underflow", func() { _ = s.Pop() })

	s.Push(entry{1})
	s.Push(entry{2})

	if want, got := 2, s.Len(); got != want {
		t.Errorf("Len want: %v, got: %v", want, got)
	}

	if want, got := (entry{2}), s.Top(); got != want {
		t.Errorf("Top want: %v, got: %v", want, got)
	}
	if want, got := (entry{2}), s.Pop(); got != want {
		t.Errorf("Pop want: %v, got: %v", want, got)
	}

	if want, got := 1, s.Len(); got != want {
		t.Errorf("Len want: %v, got: %v", want, got)
	}

	if want, got := (entry{1}), s.Pop(); got != want {
		t.Errorf("Pop want: %v, got: %v", want, got)
	}

	s.Push(entry{3})
	s.Clear()

	if want, got := 0, s.Len(); got != want {
		t.Errorf("Len want: %v, got: %v", want, got)
	}
}
