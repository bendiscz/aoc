package main

import (
	"fmt"
	"math"
	"math/big"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2023
	day     = 24
	example = `

19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3

`
)

func main() {
	Run(year, day, example, solve)
}

type stone struct {
	px, py, pz int
	vx, vy, vz int
}

type bigStone struct {
	px, py, pz *big.Int
	vx, vy, vz *big.Int
}

func newBigStone(s stone) bigStone {
	return bigStone{
		px: big.NewInt(int64(s.px)),
		py: big.NewInt(int64(s.py)),
		pz: big.NewInt(int64(s.pz)),
		vx: big.NewInt(int64(s.vx)),
		vy: big.NewInt(int64(s.vy)),
		vz: big.NewInt(int64(s.vz)),
	}
}

func (s stone) abc() (float64, float64, float64) {
	return float64(s.vy), float64(-s.vx), float64(s.vx*s.py - s.vy*s.px)
}

func solve(p *Problem) {
	stones := []stone(nil)
	for p.NextLine() {
		f := ParseInts(p.Line())
		stones = append(stones, stone{f[0], f[1], f[2], f[3], f[4], f[5]})
	}

	sMin, sMax := float64(200000000000000), float64(400000000000000)
	if p.Example() {
		sMin, sMax = 7, 27
	}

	count1 := 0
	for i, s1 := range stones {
		for _, s2 := range stones[i+1:] {
			// x1 = px1 + t1 * vx1
			// px1 + t1 * vx1 = px2 + t2 * vx2
			// t1 = (px2 - px1 + t2 * vx2) / vx1

			a1, b1, c1 := s1.abc()
			a2, b2, c2 := s2.abc()
			x := (b2*c1 - b1*c2) / (b1*a2 - b2*a1)
			y := -(a1*x + c1) / b1
			t1 := (x - float64(s1.px)) / float64(s1.vx)
			t2 := (x - float64(s2.px)) / float64(s2.vx)

			//p.Printf("%v %v %v", x, y, t1)
			if t1 >= 0 && t2 >= 0 && x >= sMin && x <= sMax && y >= sMin && y <= sMax {
				//p.Printf("%v %v %v %v", x, y, t1, t2)
				count1++
			}
		}
	}
	p.PartOne(count1)

	if p.Example() {
		return
	}
	//
	//s0 := stone{}
	//for s0.px = 0; s0.px <= 400000; s0.px++ {
	//	//p.Printf("%d", s0.px)
	//loop:
	//	for s0.vx = -200000; s0.vx <= 200000; s0.vx++ {
	//		for i, s := range stones {
	//			d := s.vx - s0.vx
	//			if d != 0 && (s0.px-s.px)%d != 0 {
	//				if i > 2 {
	//					p.Printf("fail at %d", i)
	//				}
	//				continue loop
	//			}
	//		}
	//		p.Printf("possible px, vx: %d, %d", s0.px, s0.vx)
	//	}
	//}
	//
	//t := make([]float64, len(stones))
	//r := make([]int, len(stones))
	//for vx := -3; vx <= -3; vx++ {
	//	for i, s := range stones {
	//		d := s.vx - vx
	//		if d == 0 {
	//			continue
	//		}
	//		t[i] = float64(s.px) / float64(d)
	//		r[i] = Mod(0-s.px, d)
	//	}
	//	p.Printf("%+v", r)
	//}

	//s0 := stone{}
	//tx := make([]float64, len(stones))
	//ty := make([]float64, len(stones))
	//for s0.vx = -3; s0.vx <= -3; s0.vx++ {
	//	for i, s := range stones {
	//		//d := s.vx - s0.vx
	//		//if d == 0 {
	//		//	tx[i] = -1
	//		//} else {
	//		tx[i] = float64(-s.px) / float64(s.vx-s0.vx)
	//		//}
	//	}
	//
	//	//loop:
	//	for s0.vy = 1; s0.vy <= 1; s0.vy++ {
	//		for i, s := range stones {
	//			//d := s.vy - s0.vy
	//			//if d == 0 {
	//			//	ty[i] = -1
	//			//} else {
	//			ty[i] = float64(000-s.pz) / float64(s.vz-s0.vy)
	//			//}
	//			//
	//			//if i > 0 {
	//			//	if math.Abs((tx[i]-tx[i-1])-(ty[i]-ty[i-1])) > 0.00001 {
	//			//		continue loop
	//			//	}
	//			//}
	//		}
	//
	//		//p.Printf("vx, vy = %d, %d", s0.vx, s0.vy)
	//		p.Printf("tx: %+v", tx)
	//		p.Printf("ty: %+v", ty)
	//
	//		fail := false
	//		for j := 1; j < len(stones); j++ {
	//			dx := tx[j] - tx[j-1]
	//			dy := ty[j] - ty[j-1]
	//			if dx != dy {
	//				fail = true
	//			}
	//		}
	//
	//		if !fail {
	//			p.Printf("vx, vy = %d, %d", s0.vx, s0.vy)
	//		}
	//	}
	//
	//}

	bigStones := make([]bigStone, len(stones))
	for i, s := range stones {
		bigStones[i] = newBigStone(s)
	}

	const d = 400
	for vx := -d; vx <= d; vx++ {
		if vx == 0 {
			continue
		}
		fmt.Printf("testing vx %d\n", vx)
		for vy := -d; vy <= d; vy++ {
			if vy == 0 || GCD(vx, vy) > 1 {
				continue
			}
			for vz := -d; vz <= d; vz++ {
				if vz == 0 || GCD(vx, vz) > 1 || GCD(vy, vz) > 1 {
					//fmt.Printf("skipping %d\n", vz)
					continue
				}
				test(vx, vy, vz, stones)
				if x, y, z, ok := test2(vx, vy, vz, stones); ok {
					p.PartTwo(x + y + z)
					return
				}
			}
		}
	}

	//p.PartTwo("TODO")
	//os.Exit(1)
}

func cross(x1, y1, z1, x2, y2, z2 int) (int, int, int) {
	return y1*z2 - z1*y2, z1*x2 - x1*z2, x1*y2 - y1*x2
}

func test(vx, vy, vz int, stones []stone) {
	//t0 := time.Now()
	//defer func() {
	//	fmt.Printf("done1 in %v\n", time.Since(t0))
	//}()

	//fmt.Printf("testing %d %d %d\n", vx, vy, vz)
	i := 0
	var a, b, c int
	for i < len(stones) {
		a, b, c = cross(vx, vy, vz, stones[i].vx, stones[i].vy, stones[i].vz)
		if a != 0 || b != 0 || c != 0 {
			break
		}
		i++
	}

	if i == len(stones) {
		panic("!")
	}

	//g := GCD(GCD(a, b), c)
	//a /= g
	//b /= g
	//c /= g
	d := -(a*stones[i].px + b*stones[i].py + c*stones[i].pz)
	//fmt.Printf("%v\n", d)

	var x0, y0, z0 float64
	found := false

	for _, s := range stones {
		q := a*s.vx + b*s.vy + c*s.vz
		if q == 0 {
			continue
		}

		p := a*s.px + b*s.py + c*s.pz + d
		t := float64(-p) / float64(q)

		x := float64(s.px) + t*float64(s.vx)
		y := float64(s.py) + t*float64(s.vy)
		z := float64(s.pz) + t*float64(s.vz)

		if !found {
			x0, y0, z0 = x, y, z
			found = true
		} else {
			tx := (x - x0) / float64(vx)
			ty := (y - y0) / float64(vy)
			tz := (z - z0) / float64(vz)

			if !eq(tx, ty) || !eq(tx, tz) {
				return
			}
		}
	}

	//fmt.Printf("found!!! %d %d %d\n", vx, vy, vz)
	//fmt.Printf("%d %d %d %d %d\n", i, a, b, c, d)
}

func eq(a, b float64) bool {
	return math.Abs(a-b) < 0.01
}

func test2(vxi, vyi, vzi int, stones []stone) (int, int, int, bool) {
	//t0 := time.Now()
	//defer func() {
	//	fmt.Printf("test2 done in %v\n", time.Since(t0))
	//}()

	//fmt.Printf("testing %d %d %d\n", vxi, vyi, vzi)
	i := 0
	var ai, bi, ci int
	for i < len(stones) {
		ai, bi, ci = cross(vxi, vyi, vzi, stones[i].vx, stones[i].vy, stones[i].vz)
		if ai != 0 || bi != 0 || ci != 0 {
			break
		}
		i++
	}

	if i == len(stones) {
		panic("!")
	}

	vx := big.NewRat(int64(vxi), 1)
	vy := big.NewRat(int64(vyi), 1)
	vz := big.NewRat(int64(vzi), 1)
	//fmt.Printf("v: %v %v %v\n", vx, vy, vz)

	a := big.NewRat(int64(ai), 1)
	b := big.NewRat(int64(bi), 1)
	c := big.NewRat(int64(ci), 1)
	//a := new(big.Rat).SetInt64(int64(ai))
	//b := new(big.Rat).SetInt64(int64(bi))
	//c := new(big.Rat).SetInt64(int64(ci))

	d := new(big.Rat).Mul(a, new(big.Rat).SetInt64(int64(stones[i].px)))
	d.Add(d, new(big.Rat).Mul(b, new(big.Rat).SetInt64(int64(stones[i].py))))
	d.Add(d, new(big.Rat).Mul(c, new(big.Rat).SetInt64(int64(stones[i].pz))))
	d.Neg(d)
	//fmt.Printf("%v %v %v %v\n", a, b, c, d)

	var px0, py0, pz0 *big.Rat
	found := false

	for _, s := range stones {
		q := new(big.Rat).Mul(a, big.NewRat(int64(s.vx), 1))
		q.Add(q, new(big.Rat).Mul(b, big.NewRat(int64(s.vy), 1)))
		q.Add(q, new(big.Rat).Mul(c, big.NewRat(int64(s.vz), 1)))
		if q.Num().BitLen() == 0 {
			continue
		}

		//p := a*s.px + b*s.py + c*s.pz + d
		p := new(big.Rat).Mul(a, new(big.Rat).SetInt64(int64(s.px)))
		p.Add(p, new(big.Rat).Mul(b, new(big.Rat).SetInt64(int64(s.py))))
		p.Add(p, new(big.Rat).Mul(c, new(big.Rat).SetInt64(int64(s.pz))))
		p.Add(p, d)

		t := new(big.Rat).Neg(p)
		t.Quo(t, q)
		//t := float64(-p) / float64(q)

		//fmt.Printf("t: %v\n", t)

		x := new(big.Rat).SetInt64(int64(s.vx))
		x.Mul(x, t)
		x.Add(x, new(big.Rat).SetInt64(int64(s.px)))

		y := new(big.Rat).SetInt64(int64(s.vy))
		y.Mul(y, t)
		y.Add(y, new(big.Rat).SetInt64(int64(s.py)))

		z := new(big.Rat).SetInt64(int64(s.vz))
		z.Mul(z, t)
		z.Add(z, new(big.Rat).SetInt64(int64(s.pz)))

		//px := x - t*vx
		px := new(big.Rat).Sub(x, new(big.Rat).Mul(t, vx))
		py := new(big.Rat).Sub(y, new(big.Rat).Mul(t, vy))
		pz := new(big.Rat).Sub(z, new(big.Rat).Mul(t, vz))
		//fmt.Printf("%v %v %v\n", px, py, pz)
		//x := float64(s.px) + t*float64(s.vx)
		//y := float64(s.py) + t*float64(s.vy)
		//z := float64(s.pz) + t*float64(s.vz)

		if !found {
			px0, py0, pz0 = px, py, pz
			found = true
		} else {
			//tx := new(big.Rat).Quo(new(big.Rat).Sub(x, x0), vx)
			//ty := new(big.Rat).Quo(new(big.Rat).Sub(y, y0), vy)
			//tz := new(big.Rat).Quo(new(big.Rat).Sub(z, z0), vz)

			//tx := (x - x0) / float64(vxi)
			//ty := (y - y0) / float64(vyi)
			//tz := (z - z0) / float64(vzi)

			// px := x - t * vx

			if px0.Cmp(px) != 0 || py0.Cmp(py) != 0 || pz0.Cmp(pz) != 0 {
				return 0, 0, 0, false
			}
			//if !eq(tx, ty) || !eq(tx, tz) {
			//	return
			//}
		}
	}

	//fmt.Printf("found!!! %v %v %v @ %d %d %d\n", px0, py0, pz0, vxi, vyi, vzi)
	//fmt.Printf("%d %d %d %d %d\n", i, a, b, c, d)
	return int(px0.Num().Int64()), int(py0.Num().Int64()), int(pz0.Num().Int64()), true
}

//
//func test3(vxi, vyi, vzi int, stones []bigStone) (int, int, int, bool) {
//	t0 := time.Now()
//	defer func() {
//		fmt.Printf("test3 done in %v\n", time.Since(t0))
//	}()
//
//	//fmt.Printf("testing %d %d %d\n", vxi, vyi, vzi)
//	i := 0
//	var ai, bi, ci int
//	for i < len(stones) {
//		ai, bi, ci = cross(vxi, vyi, vzi, int(stones[i].vx.Int64()), int(stones[i].vy.Int64()), int(stones[i].vz.Int64()))
//		if ai != 0 || bi != 0 || ci != 0 {
//			break
//		}
//		i++
//	}
//
//	if i == len(stones) {
//		panic("!")
//	}
//
//	vx := big.NewInt(int64(vxi))
//	vy := big.NewInt(int64(vyi))
//	vz := big.NewInt(int64(vzi))
//	//fmt.Printf("v: %v %v %v\n", vx, vy, vz)
//
//	a := big.NewInt(int64(ai))
//	b := big.NewInt(int64(bi))
//	c := big.NewInt(int64(ci))
//	//a := new(big.Rat).SetInt64(int64(ai))
//	//b := new(big.Rat).SetInt64(int64(bi))
//	//c := new(big.Rat).SetInt64(int64(ci))
//
//	d := new(big.Int).Mul(a, stones[i].px)
//	d.Add(d, new(big.Int).Mul(b, stones[i].py))
//	d.Add(d, new(big.Int).Mul(c, stones[i].pz))
//	d.Neg(d)
//	//fmt.Printf("%v %v %v %v\n", a, b, c, d)
//
//	var px0, py0, pz0 *big.Int
//	found := false
//
//	for _, s := range stones {
//		q := new(big.Int).Mul(a, s.vx)
//		q.Add(q, new(big.Int).Mul(b, s.vy))
//		q.Add(q, new(big.Int).Mul(c, s.vz))
//		if q.BitLen() == 0 {
//			continue
//		}
//
//		//p := -(a*s.px + b*s.py + c*s.pz + d)
//		p := new(big.Int).Mul(a, s.px)
//		p.Add(p, new(big.Int).Mul(b, s.py))
//		p.Add(p, new(big.Int).Mul(c, s.pz))
//		p.Add(p, d)
//		p.Neg(p)
//
//		//t := p / q
//		t := new(big.Rat).SetInt(p)
//		t = t.Quo(t, new(big.Rat).Set(q))
//
//		//fmt.Printf("t: %v\n", t)
//
//		x := new(big.Rat).SetInt(s.vx)
//		x.Mul(x, t)
//		x.Add(x, new(big.Rat).SetInt64(int64(s.px)))
//
//		y := new(big.Rat).SetInt64(int64(s.vy))
//		y.Mul(y, t)
//		y.Add(y, new(big.Rat).SetInt64(int64(s.py)))
//
//		z := new(big.Rat).SetInt64(int64(s.vz))
//		z.Mul(z, t)
//		z.Add(z, new(big.Rat).SetInt64(int64(s.pz)))
//
//		//px := x - t*vx
//		px := new(big.Rat).Sub(x, new(big.Rat).Mul(t, vx))
//		py := new(big.Rat).Sub(y, new(big.Rat).Mul(t, vy))
//		pz := new(big.Rat).Sub(z, new(big.Rat).Mul(t, vz))
//		//fmt.Printf("%v %v %v\n", px, py, pz)
//		//x := float64(s.px) + t*float64(s.vx)
//		//y := float64(s.py) + t*float64(s.vy)
//		//z := float64(s.pz) + t*float64(s.vz)
//
//		if !found {
//			px0, py0, pz0 = px, py, pz
//			found = true
//		} else {
//			//tx := new(big.Rat).Quo(new(big.Rat).Sub(x, x0), vx)
//			//ty := new(big.Rat).Quo(new(big.Rat).Sub(y, y0), vy)
//			//tz := new(big.Rat).Quo(new(big.Rat).Sub(z, z0), vz)
//
//			//tx := (x - x0) / float64(vxi)
//			//ty := (y - y0) / float64(vyi)
//			//tz := (z - z0) / float64(vzi)
//
//			// px := x - t * vx
//
//			if px0.Cmp(px) != 0 || py0.Cmp(py) != 0 || pz0.Cmp(pz) != 0 {
//				return 0, 0, 0, false
//			}
//			//if !eq(tx, ty) || !eq(tx, tz) {
//			//	return
//			//}
//		}
//	}
//
//	fmt.Printf("found!!! %v %v %v @ %d %d %d\n", px0, py0, pz0, vxi, vyi, vzi)
//	//fmt.Printf("%d %d %d %d %d\n", i, a, b, c, d)
//	return int(px0.Num().Int64()), int(py0.Num().Int64()), int(pz0.Num().Int64()), true
//}

//func test2(vx, vy, vz int, stones []stone) {
//	for i := 1; i < len(stones); i++ {
//		if !intersect(stones[0], stones[i], vx, vy, vz) {
//			return
//		}
//	}
//	fmt.Printf("found!!! %d %d %d\n", vx, vy, vz)
//}
//
//func intersect(s1, s2 stone, vx, vy, vz int) bool {
//	tx := float64(s2.px-s2.px) / float64(s1.vx-s2.vx)
//	ty := float64(s2.py-s2.py) / float64(s1.vy-s2.vy)
//	tz := float64(s2.pz-s2.pz) / float64(s1.vz-s2.vz)
//
//	return tx == ty && tx == tz
//}
