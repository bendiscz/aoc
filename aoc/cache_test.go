package aoc

import "testing"

type state string

func (s state) Key() string {
	return string(s)
}

func TestCachedFunc(t *testing.T) {
	called := map[string]bool{}
	c := NewCachedFunc(func(in state) string {
		s := string(in)
		if called[s] {
			t.Errorf("fn called repeatedly for %s", s)
		}
		return s
	})

	c("")
	c("a")
	c("b")
	c("a")
	c("")
}
