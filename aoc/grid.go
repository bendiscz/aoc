package aoc

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

var (
	PosX = XY{1, 0}
	NegX = XY{-1, 0}
	PosY = XY{0, 1}
	NegY = XY{0, -1}

	HVDirs = [...]XY{PosX, NegX, PosY, NegY}
)

type XY struct {
	X, Y int
}

func (c XY) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func (c XY) Trans() XY {
	return XY{c.Y, c.X}
}

func (c XY) Add(d XY) XY { return XY{c.X + d.X, c.Y + d.Y} }
func (c XY) Sub(d XY) XY { return XY{c.X - d.X, c.Y - d.Y} }
func (c XY) Neg() XY     { return XY{-c.X, -c.Y} }

func Rectangle(x, y int) XY {
	return XY{x, y}
}

func Square(a int) XY {
	return XY{a, a}
}

func (c XY) HasInsideX(x int) bool { return x >= 0 && x < c.X }
func (c XY) HasInsideY(y int) bool { return y >= 0 && y < c.Y }
func (c XY) HasInside(a XY) bool   { return c.HasInsideX(a.X) && c.HasInsideY(a.Y) }

type Grid[T any] interface {
	Size() XY
	At(c XY) *T
	AtXY(x, y int) *T
	Row(y int) Vector[T]
	Column(x int) Vector[T]
	TransView() Grid[T]
}

type Vector[T any] interface {
	Len() int
	At(i int) *T
}

type Matrix[T any] struct {
	Dim  XY
	Data [][]T
}

func NewMatrix[T any](dim XY) *Matrix[T] {
	m := &Matrix[T]{
		Dim:  dim,
		Data: make([][]T, dim.Y),
	}
	for y := range m.Data {
		m.Data[y] = make([]T, dim.X)
	}
	return m
}

func (m *Matrix[T]) Size() XY               { return m.Dim }
func (m *Matrix[T]) At(c XY) *T             { return m.AtXY(c.X, c.Y) }
func (m *Matrix[T]) AtXY(x, y int) *T       { return &m.Data[y][x] }
func (m *Matrix[T]) Row(y int) Vector[T]    { return matrixRow[T]{m, y} }
func (m *Matrix[T]) Column(x int) Vector[T] { return matrixColumn[T]{m, x} }
func (m *Matrix[T]) TransView() Grid[T]     { return transMatrixView[T]{m} }

func (m *Matrix[T]) AppendRow() Vector[T] {
	m.Data = append(m.Data, make([]T, m.Dim.X))
	m.Dim.Y++
	return m.Row(m.Dim.Y - 1)
}

type matrixRow[T any] struct {
	m *Matrix[T]
	y int
}

func (r matrixRow[T]) Len() int    { return r.m.Size().X }
func (r matrixRow[T]) At(i int) *T { return r.m.AtXY(i, r.y) }

type matrixColumn[T any] struct {
	m *Matrix[T]
	x int
}

func (c matrixColumn[T]) Len() int    { return c.m.Size().Y }
func (c matrixColumn[T]) At(i int) *T { return c.m.AtXY(c.x, i) }

type transMatrixView[T any] struct {
	m *Matrix[T]
}

func (t transMatrixView[T]) Size() XY               { return t.m.Size().Trans() }
func (t transMatrixView[T]) At(c XY) *T             { return t.m.At(c.Trans()) }
func (t transMatrixView[T]) AtXY(x, y int) *T       { return t.m.AtXY(y, x) }
func (t transMatrixView[T]) Row(y int) Vector[T]    { return t.m.Column(y) }
func (t transMatrixView[T]) Column(x int) Vector[T] { return t.m.Row(x) }
func (t transMatrixView[T]) TransView() Grid[T]     { return t.m }

type Slice[T any] []T

func (s Slice[T]) Len() int    { return len(s) }
func (s Slice[T]) At(i int) *T { return &s[i] }

func CopyVector[T ~byte](s Vector[T], line string) {
	for i, ch := range []byte(line) {
		*s.At(i) = T(ch)
	}
}

func ParseVectorDigits[T constraints.Integer](s Vector[T], line string) {
	for i, ch := range []byte(line) {
		*s.At(i) = T(ch - '0')
	}
}

func ParseVectorMap[T any](s Vector[T], line string, mapping map[byte]T) {
	for i, ch := range []byte(line) {
		*s.At(i) = mapping[ch]
	}
}

func ParseVectorFunc[T any](s Vector[T], line string, mapping func(byte) T) {
	for i, ch := range []byte(line) {
		*s.At(i) = mapping(ch)
	}
}

const (
	SymbolEmpty = "??"
	SymbolFull  = "???"
)

func PrintGridFunc[T any](g Grid[T], fn func(T) string) {
	m := g.Size()
	for y := 0; y < m.Y; y++ {
		for x := 0; x < m.X; x++ {
			fmt.Print(fn(*g.AtXY(x, y)))
		}
		fmt.Println()
	}
}

func PrintGridBool(g Grid[bool]) {
	PrintGridFunc[bool](g, func(x bool) string {
		if x {
			return SymbolFull
		} else {
			return SymbolEmpty
		}
	})
}

func PrintGrid[T fmt.Stringer](g Grid[T]) {
	PrintGridFunc[T](g, func(x T) string { return x.String() })
}
