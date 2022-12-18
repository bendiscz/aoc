package aoc

import (
	"log"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Signed](x T) T {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func Sign[T constraints.Signed](x T) T {
	if x == 0 {
		return 0
	} else if x > 0 {
		return 1
	} else {
		return -1
	}
}

func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	} else {
		return y
	}
}

func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	} else {
		return y
	}
}

func ParseInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("invalid int: %#v", s)
	}
	return x
}

func ParseUint(s string) uint {
	x, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Panicf("invalid uint: %#v", s)
	}
	return uint(x)
}

func SplitFields(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool { return unicode.IsSpace(r) || r == ',' || r == ';' })
}

func ParseInts(s string) []int {
	f := SplitFields(s)
	n := make([]int, len(f))
	for i, x := range f {
		n[i] = ParseInt(x)
	}
	return n
}

func ParseUints(s string) []uint {
	f := SplitFields(s)
	n := make([]uint, len(f))
	for i, x := range f {
		n[i] = ParseUint(x)
	}
	return n
}
