package aoc

import (
	"fmt"
	"testing"
)

func TestRational_Integer(t *testing.T) {
	tests := []int{
		0,
		1,
		-1,
		1000,
	}

	for _, want := range tests {
		t.Run(fmt.Sprintf("Integer(%d)", want), func(t *testing.T) {
			if got := Integer(want); got.D() != 1 || got.N() != want {
				t.Errorf("Integer(%d) got = %v, want = %v", want, got, Rational{n: want})
			}
		})
	}
}

func TestRational_Ratio(t *testing.T) {
	tests := []struct {
		x, y int
		want string
	}{
		{1, 1, "1/1"},
		{-1, 1, "-1/1"},
		{-1, -1, "1/1"},
		{1, -1, "-1/1"},
		{1, 2, "1/2"},
		{2, 4, "1/2"},
		{500, 1000, "1/2"},
		{13, 97, "13/97"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Ratio(%d,%d)", tc.x, tc.y), func(t *testing.T) {
			if got := Ratio(tc.x, tc.y).String(); got != tc.want {
				t.Errorf("Ratio(%d, %d) got = %v, want = %v", tc.x, tc.y, got, tc.want)
			}
		})
	}
}

func TestRational_Ratio_DivideByZero(t *testing.T) {
	defer func() {
		if got, want := recover(), "divide by zero"; got != want {
			t.Errorf("Ratio(1, 0) got = %v, want = %v", got, want)
		}
	}()

	Ratio(1, 0)
}

func TestRational_Neg(t *testing.T) {
	tests := []struct {
		r    Rational
		want string
	}{
		{Ratio(1, 1), "-1/1"},
		{Ratio(1, 2), "-1/2"},
		{Ratio(-1, 2), "1/2"},
		{Ratio(0, 2), "0/1"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v).Neg()", tc.r), func(t *testing.T) {
			if got := tc.r.Neg().String(); got != tc.want {
				t.Errorf("(%v).Neg() got = %v, want = %v", tc.r, got, tc.want)
			}
		})
	}
}

func TestRational_Inv(t *testing.T) {
	tests := []struct {
		r    Rational
		want string
	}{
		{Ratio(1, 1), "1/1"},
		{Ratio(1, 2), "2/1"},
		{Ratio(-1, 2), "-2/1"},
		{Ratio(30, 10), "1/3"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v).Inv()", tc.r), func(t *testing.T) {
			if got := tc.r.Inv().String(); got != tc.want {
				t.Errorf("(%v).Inv() got = %v, want = %v", tc.r, got, tc.want)
			}
		})
	}
}

func TestRational_Inv_DivideByZero(t *testing.T) {
	defer func() {
		if got, want := recover(), "divide by zero"; got != want {
			t.Errorf("(0/1).Inv() got = %v, want = %v", got, want)
		}
	}()

	Ratio(0, 1).Inv()
}

func TestRational_Add(t *testing.T) {
	tests := []struct {
		a, b Rational
		want string
	}{
		{Ratio(1, 1), Ratio(1, 1), "2/1"},
		{Ratio(1, 1), Ratio(-1, 1), "0/1"},
		{Ratio(1, 2), Ratio(1, 2), "1/1"},
		{Ratio(1, 2), Ratio(1, 3), "5/6"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)+(%v)", tc.a, tc.b), func(t *testing.T) {
			if got := tc.a.Add(tc.b).String(); got != tc.want {
				t.Errorf("%v + %v got = %v, want = %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestRational_Sub(t *testing.T) {
	tests := []struct {
		a, b Rational
		want string
	}{
		{Integer(1), Integer(1), "0/1"},
		{Integer(1), Integer(-1), "2/1"},
		{Integer(1), Ratio(1, 2), "1/2"},
		{Ratio(1, 2), Ratio(1, 3), "1/6"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)-(%v)", tc.a, tc.b), func(t *testing.T) {
			if got := tc.a.Sub(tc.b).String(); got != tc.want {
				t.Errorf("%v - %v got = %v, want = %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestRational_Mul(t *testing.T) {
	tests := []struct {
		a, b Rational
		want string
	}{
		{Integer(1), Integer(1), "1/1"},
		{Integer(1), Integer(-1), "-1/1"},
		{Ratio(1, 2), Ratio(1, 2), "1/4"},
		{Ratio(1, 3), Ratio(3, 1), "1/1"},
		{Ratio(3, 2), Ratio(5, 3), "5/2"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)*(%v)", tc.a, tc.b), func(t *testing.T) {
			if got := tc.a.Mul(tc.b).String(); got != tc.want {
				t.Errorf("%v * %v got = %v, want = %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestRational_Div(t *testing.T) {
	tests := []struct {
		a, b Rational
		want string
	}{
		{Integer(1), Integer(1), "1/1"},
		{Integer(1), Integer(-1), "-1/1"},
		{Ratio(1, 2), Ratio(1, 2), "1/1"},
		{Ratio(1, 3), Ratio(3, 1), "1/9"},
		{Ratio(3, 2), Ratio(5, 3), "9/10"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)/(%v)", tc.a, tc.b), func(t *testing.T) {
			if got := tc.a.Div(tc.b).String(); got != tc.want {
				t.Errorf("%v / %v got = %v, want = %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestRational_IsInt(t *testing.T) {
	tests := []struct {
		a    Rational
		want bool
	}{
		{Integer(0), true},
		{Integer(1), true},
		{Ratio(10, 2), true},
		{Ratio(3, 2), false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v).IsInt()", tc.a), func(t *testing.T) {
			if got := tc.a.IsInt(); got != tc.want {
				t.Errorf("(%v).IsInt() got = %v, want = %v", tc.a, got, tc.want)
			}
		})
	}
}

func TestRational_ToFloat32(t *testing.T) {
	tests := []struct {
		a    Rational
		want float32
	}{
		{Integer(0), 0},
		{Integer(1), 1},
		{Ratio(10, 2), 5},
		{Ratio(3, 2), 1.5},
		{Ratio(-1, 3), -1.0 / 3},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v).ToFloat32()", tc.a), func(t *testing.T) {
			if got := tc.a.ToFloat32(); got != tc.want {
				t.Errorf("(%v).ToFloat32() got = %v, want = %v", tc.a, got, tc.want)
			}
		})
	}
}

func TestRational_ToFloat64(t *testing.T) {
	tests := []struct {
		a    Rational
		want float64
	}{
		{Integer(0), 0},
		{Integer(1), 1},
		{Ratio(10, 2), 5},
		{Ratio(3, 2), 1.5},
		{Ratio(-1, 3), -1.0 / 3},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v).ToFloat64()", tc.a), func(t *testing.T) {
			if got := tc.a.ToFloat64(); got != tc.want {
				t.Errorf("(%v).ToFloat64() got = %v, want = %v", tc.a, got, tc.want)
			}
		})
	}
}
