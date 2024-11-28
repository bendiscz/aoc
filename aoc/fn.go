package aoc

import (
	"log"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/constraints"
)

const Hex = "0123456789abcdef"

func Mod[T constraints.Integer](a, b T) T {
	a %= b
	if a >= 0 {
		return a
	}
	if b < 0 {
		return a - b
	}
	return a + b
}

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

func GCD[T constraints.Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM[T constraints.Integer](a, b T) T {
	return a / GCD(a, b) * b
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

func SplitFieldsDelim(s, delim string) []string {
	return strings.FieldsFunc(s, func(r rune) bool { return strings.ContainsRune(delim, r) })
}

func ParseInts(s string) []int {
	f := strings.FieldsFunc(s, func(r rune) bool { return r != '-' && !unicode.IsDigit(r) })
	n := make([]int, len(f))
	for i, x := range f {
		n[i] = ParseInt(x)
	}
	return n
}

func ParseUints(s string) []uint {
	f := strings.FieldsFunc(s, func(r rune) bool { return !unicode.IsDigit(r) })
	n := make([]uint, len(f))
	for i, x := range f {
		n[i] = ParseUint(x)
	}
	return n
}
