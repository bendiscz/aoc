package aoc

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"0", 0},
		{"-0", 0},
		{"123", 123},
		{"-9223372036854775808", math.MinInt64},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseInt(test.name)
			if got != test.want {
				t.Errorf("expected %v, got: %v", test.want, got)
			}
		})
	}
}

func TestParseUint(t *testing.T) {
	tests := []struct {
		name string
		want uint
	}{
		{"0", 0},
		{"123", 123},
		{"18446744073709551615", math.MaxUint64},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if got := ParseUint(test.name); got != test.want {
				t.Errorf("expected %v, got: %v", test.want, got)
			}
		})
	}
}

func TestParseInt_Panics(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"", "invalid int: \"\""},
		{"abc", "invalid int: \"abc\""},
		{"123a", "invalid int: \"123a\""},
		{"-1.0", "invalid int: \"-1.0\""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertPanic(t, test.want, func() { ParseInt(test.name) })
		})
	}
}

func TestParseUint_Panics(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"", "invalid uint: \"\""},
		{"abc", "invalid uint: \"abc\""},
		{"123a", "invalid uint: \"123a\""},
		{"-1", "invalid uint: \"-1\""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertPanic(t, test.want, func() { ParseUint(test.name) })
		})
	}
}

func TestMinMax(t *testing.T) {
	tests := []struct {
		x, y     int
		min, max int
	}{
		{0, 0, 0, 0},
		{1, 2, 1, 2},
		{1, -1, -1, 1},
		{-2, -1, -2, -1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d,%d", test.x, test.y), func(t *testing.T) {
			if got := min(test.x, test.y); got != test.min {
				t.Errorf("Min want: %v, got: %v", test.min, got)
			}
			if got := max(test.x, test.y); got != test.max {
				t.Errorf("Max want: %v, got: %v", test.max, got)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{0, 0},
		{1, 1},
		{-1, 1},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.x), func(t *testing.T) {
			if got := Abs(test.x); got != test.want {
				t.Errorf("Abs want: %v, got: %v", test.want, got)
			}
		})
	}
}

func TestSign(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{0, 0},
		{10, 1},
		{-20, -1},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.x), func(t *testing.T) {
			if got := Sign(test.x); got != test.want {
				t.Errorf("Sign want: %v, got: %v", test.want, got)
			}
		})
	}
}

func assertPanic(t *testing.T, want any, f func()) {
	defer func() {
		if r := recover(); r != want {
			t.Errorf("want: %v; got: %v", want, r)
		}
	}()
	f()
}
