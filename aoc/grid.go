package aoc

import (
	"fmt"
	"iter"

	"golang.org/x/exp/constraints"
)

var (
	XY0 = XY{0, 0}

	PosX = XY{1, 0}
	NegX = XY{-1, 0}
	PosY = XY{0, 1}
	NegY = XY{0, -1}

	Q1 = XY{1, 1}
	Q2 = XY{-1, 1}
	Q3 = XY{-1, -1}
	Q4 = XY{1, -1}

	HVDirs   = [...]XY{PosX, NegX, PosY, NegY}
	DiagDirs = [...]XY{Q1, Q2, Q3, Q4}
	AllDirs  = [...]XY{PosX, NegX, PosY, NegY, Q1, Q2, Q3, Q4}
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

func (c XY) ForEach(fn func(xy XY)) {
	for y := 0; y < c.Y; y++ {
		for x := 0; x < c.X; x++ {
			fn(XY{x, y})
		}
	}
}

func MaxXY(xy XY, xys ...XY) XY {
	for _, xy2 := range xys {
		xy = XY{max(xy.X, xy2.X), max(xy.Y, xy2.Y)}
	}
	return xy
}

func MinXY(xy XY, xys ...XY) XY {
	for _, xy2 := range xys {
		xy = XY{min(xy.X, xy2.X), min(xy.Y, xy2.Y)}
	}
	return xy
}

func (c XY) All() iter.Seq[XY] {
	return func(yield func(XY) bool) {
		for y := 0; y < c.Y; y++ {
			for x := 0; x < c.X; x++ {
				yield(XY{X: x, Y: y})
			}
		}
	}
}

type Grid[T any] interface {
	Size() XY
	At(c XY) *T
	AtXY(x, y int) *T
	Row(y int) Vector[T]
	Column(x int) Vector[T]
	Trans() Grid[T]
	FlipX() Grid[T]
	FlipY() Grid[T]
	RotateR() Grid[T]
	RotateL() Grid[T]
	SubGrid(off, dim XY) Grid[T]
	All() iter.Seq2[XY, *T]
}

func iterateGrid[T any](g Grid[T]) iter.Seq2[XY, *T] {
	return func(yield func(XY, *T) bool) {
		dim := g.Size()
		for y := 0; y < dim.Y; y++ {
			for x := 0; x < dim.X; x++ {
				xy := XY{X: x, Y: y}
				if !yield(xy, g.At(xy)) {
					return
				}
			}
		}
	}
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

func (m *Matrix[T]) Size() XY                    { return m.Dim }
func (m *Matrix[T]) At(c XY) *T                  { return m.AtXY(c.X, c.Y) }
func (m *Matrix[T]) AtXY(x, y int) *T            { return &m.Data[y][x] }
func (m *Matrix[T]) Row(y int) Vector[T]         { return matrixRow[T]{m, y} }
func (m *Matrix[T]) Column(x int) Vector[T]      { return matrixColumn[T]{m, x} }
func (m *Matrix[T]) Trans() Grid[T]              { return gridView[T]{grid: m}.Trans() }
func (m *Matrix[T]) FlipX() Grid[T]              { return gridView[T]{grid: m}.FlipX() }
func (m *Matrix[T]) FlipY() Grid[T]              { return gridView[T]{grid: m}.FlipY() }
func (m *Matrix[T]) RotateR() Grid[T]            { return gridView[T]{grid: m}.RotateR() }
func (m *Matrix[T]) RotateL() Grid[T]            { return gridView[T]{grid: m}.RotateL() }
func (m *Matrix[T]) SubGrid(off, dim XY) Grid[T] { return subGrid[T]{m, off, dim} }
func (m *Matrix[T]) All() iter.Seq2[XY, *T]      { return iterateGrid(m) }

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

func CloneGrid[T any](g Grid[T]) *Matrix[T] {
	m := NewMatrix[T](g.Size())
	for x := 0; x < m.Dim.X; x++ {
		for y := 0; y < m.Dim.Y; y++ {
			m.Data[y][x] = *g.AtXY(x, y)
		}
	}
	return m
}

func CollectGrid[T any, V any](cells map[XY]V, fn func(V) T) (*Matrix[T], XY) {
	l, h := XY0, XY0
	for k := range cells {
		l = MinXY(l, k)
		h = MaxXY(h, k)
	}
	dim := h.Sub(l).Add(Square(1))
	m := NewMatrix[T](dim)
	for c, v := range cells {
		*m.At(c.Sub(l)) = fn(v)
	}
	return m, l
}

func CopyGrid[T any](dst, src Grid[T]) {
	dim := MinXY(dst.Size(), src.Size())
	for x := 0; x < dim.X; x++ {
		for y := 0; y < dim.Y; y++ {
			*dst.AtXY(x, y) = *src.AtXY(x, y)
		}
	}
}

const (
	SymbolEmpty  = "·"
	SymbolBullet = "•"
	SymbolFull   = "█"
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

type BoolCell struct {
	V bool
}

func (c BoolCell) String() string {
	if c.V {
		return SymbolFull
	} else {
		return SymbolEmpty
	}
}

type ByteCell struct {
	V byte
}

func (c ByteCell) String() string {
	return fmt.Sprintf("%c", c.V)
}

type gridView[T any] struct {
	grid Grid[T]
	negX bool
	negY bool
	swap bool
}

func (g gridView[T]) Size() XY {
	if g.swap {
		return g.grid.Size().Trans()
	}
	return g.grid.Size()
}

func (g gridView[T]) At(c XY) *T {
	return g.AtXY(c.X, c.Y)
}

func (g gridView[T]) AtXY(x, y int) *T {
	dim := g.grid.Size()
	if g.swap {
		x, y = y, x
	}
	if g.negX {
		x = dim.X - 1 - x
	}
	if g.negY {
		y = dim.Y - 1 - y
	}
	return g.grid.AtXY(x, y)
}

func (g gridView[T]) Row(y int) Vector[T]    { return gridViewRow[T]{g, y} }
func (g gridView[T]) Column(x int) Vector[T] { return gridViewColumn[T]{g, x} }

func (g gridView[T]) Trans() Grid[T] {
	return gridView[T]{
		grid: g.grid,
		negX: g.negX,
		negY: g.negY,
		swap: !g.swap,
	}
}

func (g gridView[T]) FlipX() Grid[T] {
	return gridView[T]{
		grid: g.grid,
		negX: !g.negX != g.swap,
		negY: g.negY != g.swap,
		swap: g.swap,
	}
}

func (g gridView[T]) FlipY() Grid[T] {
	return gridView[T]{
		grid: g.grid,
		negX: g.negX != g.swap,
		negY: !g.negY != g.swap,
		swap: g.swap,
	}
}

func (g gridView[T]) RotateR() Grid[T]            { return g.Trans().FlipX() }
func (g gridView[T]) RotateL() Grid[T]            { return g.Trans().FlipY() }
func (g gridView[T]) SubGrid(off, dim XY) Grid[T] { return subGrid[T]{g, off, dim} }
func (g gridView[T]) All() iter.Seq2[XY, *T]      { return iterateGrid(g) }

type gridViewRow[T any] struct {
	gridView[T]
	y int
}

func (g gridViewRow[T]) Len() int    { return g.Size().X }
func (g gridViewRow[T]) At(i int) *T { return g.AtXY(i, g.y) }

type gridViewColumn[T any] struct {
	gridView[T]
	x int
}

func (g gridViewColumn[T]) Len() int    { return g.Size().X }
func (g gridViewColumn[T]) At(i int) *T { return g.AtXY(g.x, i) }

type subGrid[T any] struct {
	grid Grid[T]
	off  XY
	dim  XY
}

func (g subGrid[T]) Size() XY                    { return g.dim }
func (g subGrid[T]) At(c XY) *T                  { return g.grid.At(c.Add(g.off)) }
func (g subGrid[T]) AtXY(x, y int) *T            { return g.grid.AtXY(x+g.off.X, y+g.off.Y) }
func (g subGrid[T]) Row(y int) Vector[T]         { return subGridRow[T]{g, y + g.off.Y} }
func (g subGrid[T]) Column(x int) Vector[T]      { return subGridColumn[T]{g, x + g.off.X} }
func (g subGrid[T]) Trans() Grid[T]              { return gridView[T]{grid: g}.Trans() }
func (g subGrid[T]) FlipX() Grid[T]              { return gridView[T]{grid: g}.FlipX() }
func (g subGrid[T]) FlipY() Grid[T]              { return gridView[T]{grid: g}.FlipY() }
func (g subGrid[T]) RotateR() Grid[T]            { return gridView[T]{grid: g}.RotateR() }
func (g subGrid[T]) RotateL() Grid[T]            { return gridView[T]{grid: g}.RotateL() }
func (g subGrid[T]) SubGrid(off, dim XY) Grid[T] { return subGrid[T]{g.grid, g.off.Add(off), dim} }
func (g subGrid[T]) All() iter.Seq2[XY, *T]      { return iterateGrid(g) }

type subGridRow[T any] struct {
	subGrid[T]
	y int
}

func (g subGridRow[T]) Len() int    { return g.dim.X }
func (g subGridRow[T]) At(i int) *T { return g.grid.AtXY(i+g.off.X, g.y) }

type subGridColumn[T any] struct {
	subGrid[T]
	x int
}

func (g subGridColumn[T]) Len() int    { return g.dim.Y }
func (g subGridColumn[T]) At(i int) *T { return g.grid.AtXY(g.x, i+g.off.Y) }
