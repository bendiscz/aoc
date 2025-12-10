package aoc

import "fmt"

type Rational struct {
	n, d int
}

func Integer(x int) Rational {
	return Rational{n: x}
}

func Ratio(x, y int) Rational {
	if y == 0 {
		panic("divide by zero")
	}
	if y < 0 {
		x, y = -x, -y
	}
	return Rational{n: x, d: y - 1}.norm()
}

func (r Rational) N() int { return r.n }
func (r Rational) D() int { return r.d + 1 }

func (r Rational) IsInt() bool { return r.d == 0 }

func (r Rational) norm() Rational {
	gcd := GCD(Abs(r.n), r.d+1)
	if gcd == 1 {
		return r
	}
	return Rational{r.n / gcd, (r.d+1)/gcd - 1}
}

func (r Rational) String() string {
	return fmt.Sprintf("%d/%d", r.n, r.d+1)
}

func (r Rational) Add(r2 Rational) Rational {
	d2 := LCM(r.d, r2.d+1)
	res := Rational{r.n*(d2/(r.d+1)) + r2.n*(d2/(r2.d+1)), d2 - 1}
	return res.norm()
}

func (r Rational) Mul(r2 Rational) Rational {
	res := Rational{r.n * r2.n, (r.d+1)*(r2.d+1) - 1}
	return res.norm()
}

func (r Rational) Inv() Rational {
	if r.n == 0 {
		panic("divide by zero")
	}
	s := Sign(r.n)
	return Rational{s * (r.d + 1), Abs(r.n) - 1}
}
