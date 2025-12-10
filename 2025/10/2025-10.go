package main

import (
	"fmt"
	"math"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2025
	day     = 10
	example = `

[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}

`
)

func main() {
	Run(year, day, example, solve)
}

type machine struct {
	lights  int
	buttons []int
	jolt    []int
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	machines := []machine(nil)
	for p.NextLine() {
		f := SplitFieldsDelim(p.Line(), " ")
		m := machine{}
		for i := 1; i < len(f[0])-1; i++ {
			if f[0][i] == '#' {
				m.lights |= 1 << (i - 1)
			}
		}
		for i := 1; i < len(f)-1; i++ {
			b := 0
			for _, x := range ParseInts(f[i]) {
				b |= 1 << x
			}
			m.buttons = append(m.buttons, b)
		}
		m.jolt = ParseInts(f[len(f)-1])

		machines = append(machines, m)
	}

	for _, m := range machines {
		s1 += count1(m)
		s2 += count2(m)
		//p.Printf("%d %d", s1, s2)
	}

	p.PartOne(s1)
	p.PartTwo(s2)

}

func count1(m machine) int {
	best := math.MaxInt
	for c := 0; c < 1<<len(m.buttons); c++ {
		x, n := 0, 0
		for i, b := range m.buttons {
			if c&(1<<i) != 0 {
				x = x ^ b
				n++
			}
		}
		if x == m.lights {
			best = min(best, n)
		}
	}
	return best
}

const N = 13

// a_0 * x_1 + a_1 * x_2 + ... + a_N * x_N + b = 0
type linear struct {
	a [N]rational
	b rational
}

func newLinear() linear {
	l := linear{b: integer(0)}
	for i := 0; i < N; i++ {
		l.a[i] = integer(0)
	}
	return l
}

func extract(lin linear, index int) (linear, bool) {
	a := lin.a[index].mul(integer(-1))
	if a.n == 0 {
		return linear{}, false
	}

	a = a.inv()

	r := newLinear()
	r.b = lin.b.mul(a)
	for i := 0; i < N; i++ {
		if i != index {
			r.a[i] = lin.a[i].mul(a)
		}
	}
	return r, true
}

func substitute(lin linear, index int, expr linear) linear {
	r := newLinear()

	a := lin.a[index]
	lin.a[index] = integer(0)

	for i := 0; i < N; i++ {
		r.a[i] = lin.a[i].plus(a.mul(expr.a[i]))
	}
	r.b = lin.b.plus(a.mul(expr.b))
	return r
}

type variable struct {
	expr linear
	free bool
	val  int
	max  int
}

func eval(v variable, vals [N]int) rational {
	if v.free {
		return integer(v.val)
	}

	x := v.expr.b
	for i := 0; i < N; i++ {
		x = x.plus(v.expr.a[i].mul(integer(vals[i])))
	}
	return x
}

func count2(m machine) int {
	vars := make([]variable, len(m.buttons))
	for i := range vars {
		vars[i].max = math.MaxInt
	}

	eqs := make([]linear, len(m.jolt))
	for i, jolt := range m.jolt {
		eq := newLinear()
		eq.b = integer(-jolt)
		for j, b := range m.buttons {
			if b&(1<<i) != 0 {
				eq.a[j] = integer(1)
				vars[j].max = min(vars[j].max, jolt)
			}
		}
		eqs[i] = eq
	}

	for i := range vars {
		vars[i].free = true

		for _, eq := range eqs {
			if expr, ok := extract(eq, i); ok {
				vars[i].free = false
				vars[i].expr = expr

				for j := range eqs {
					eqs[j] = substitute(eqs[j], i, expr)
				}

				break
			}
		}
	}

	free := []int(nil)
	for i, v := range vars {
		if v.free {
			free = append(free, i)
		}
	}

	best, _ := evalRecursive(vars, free, 0)
	return best
}

func evalRecursive(vars []variable, free []int, index int) (int, bool) {
	if index == len(free) {
		vals := [N]int{}
		total := 0

		for i := len(vars) - 1; i >= 0; i-- {
			x := eval(vars[i], vals)
			if x.n < 0 || x.d != 1 {
				return 0, false
			}
			vals[i] = x.n
			total += vals[i]
		}

		return total, true
	}

	best, found := math.MaxInt, false
	for x := 0; x <= vars[free[index]].max; x++ {
		vars[free[index]].val = x
		total, ok := evalRecursive(vars, free, index+1)

		if ok {
			found = true
			best = min(best, total)
		}
	}

	if found {
		return best, true
	} else {
		return 0, false
	}
}

type rational struct {
	n, d int
}

func integer(x int) rational {
	return rational{x, 1}
}

func (r rational) norm() rational {
	c := GCD(Abs(r.n), r.d)
	if c == 1 {
		return r
	}
	return rational{r.n / c, r.d / c}
}

func (r rational) plus(r2 rational) rational {
	d2 := LCM(r.d, r2.d)
	res := rational{r.n*(d2/r.d) + r2.n*(d2/r2.d), d2}
	return res.norm()
}

func (r rational) mul(r2 rational) rational {
	res := rational{r.n * r2.n, r.d * r2.d}
	return res.norm()
}

func (r rational) inv() rational {
	if r.n == 0 {
		panic("divide by zero")
	}
	s := Sign(r.n)
	return rational{s * r.d, Abs(r.n)}
}

func (r rational) String() string {
	return fmt.Sprintf("%d/%d", r.n, r.d)
}

//func solve2(m machine) int {
//	ctx := z3.NewContext(nil)
//	s := z3.NewSolver(ctx)
//	zero := number(ctx, 0)
//
//	counts := []z3.Int(nil)
//	for i := 0; i < len(m.buttons); i++ {
//		b := ctx.IntConst(fmt.Sprintf("b%d", i))
//		s.Assert(b.GE(zero))
//		counts = append(counts, b)
//	}
//
//	for i, jolt := range m.jolt {
//		press := []z3.Int(nil)
//		for j, b := range m.buttons {
//			if b&(1<<i) != 0 {
//				press = append(press, counts[j])
//			}
//		}
//		s.Assert(zero.Add(press...).Eq(number(ctx, jolt)))
//	}
//
//	total := ctx.IntConst("total")
//	s.Assert(total.Eq(zero.Add(counts...)))
//
//	best := 0
//	for {
//		if sat, err := s.Check(); !sat || err != nil {
//			return best
//		}
//
//		model := s.Model()
//
//		t, _, _ := model.Eval(total, true).(z3.Int).AsInt64()
//		best = int(t)
//
//		s.Assert(total.LT(number(ctx, best)))
//	}
//}
//
//func number(ctx *z3.Context, x int) z3.Int {
//	return ctx.FromInt(int64(x), ctx.IntSort()).(z3.Int)
//}
