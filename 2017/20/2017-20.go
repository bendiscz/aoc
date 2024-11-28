package main

import (
	"math"
	"sort"

	. "github.com/bendiscz/aoc/aoc"
)

const (
	year    = 2017
	day     = 20
	example = `

p=<-6,0,0>, v=< 3,0,0>, a=< 0,0,0>
p=<-4,0,0>, v=< 2,0,0>, a=< 0,0,0>
p=<-2,0,0>, v=< 1,0,0>, a=< 0,0,0>
p=< 3,0,0>, v=<-1,0,0>, a=< 0,0,0>

`
)

//p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
//p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>

func main() {
	Run(year, day, example, solve)
}

type coord [3]int

func (c coord) abs() int {
	return Abs(c[0]) + Abs(c[1]) + Abs(c[2])
}

const never = math.MaxInt

type particle struct {
	p, v, a coord
	dt      int
}

type collision struct {
	t    int
	i, j int
}

func solve(p *Problem) {
	particles := []particle(nil)
	for i := 0; p.NextLine(); i++ {
		f := ParseInts(p.Line())
		particles = append(particles, particle{
			p:  coord{f[0], f[1], f[2]},
			v:  coord{f[3], f[4], f[5]},
			a:  coord{f[6], f[7], f[8]},
			dt: never,
		})
	}

	p.PartOne(find(particles))

	collisions := []collision(nil)
	for i := 0; i < len(particles); i++ {
		for j := i + 1; j < len(particles); j++ {
			if t := collide(particles[i], particles[j]); t != never {
				collisions = append(collisions, collision{t, i, j})
			}
		}
	}
	sort.Slice(collisions, func(i, j int) bool { return collisions[i].t < collisions[j].t })

	for _, c := range collisions {
		if particles[c.i].dt < c.t || particles[c.j].dt < c.t {
			continue
		}

		particles[c.i].dt = c.t
		particles[c.j].dt = c.t
	}

	count := 0
	for _, pt := range particles {
		if pt.dt == never {
			count++
		}
	}
	p.PartTwo(count)
}

func find(particles []particle) int {
	mv, ma, mi := math.MaxInt, math.MaxInt, 0
	for i, pt := range particles {
		v := pt.v.abs()
		a := pt.a.abs()
		if a < ma {
			mv = v
			ma = a
			mi = i
		} else if a == ma && v < mv {
			mv = v
			mi = i
		}
	}
	return mi
}

func collide(pt1, pt2 particle) int {
	t0, t1, all := [3]float64{}, [3]float64{}, [3]bool{}
	for i := 0; i < 3; i++ {
		// p1 + v1*t + 1/2*a1*t*(t+1)
		// 1/2*a1*t^2 + (v1 + 1/2*a1)*t + p1
		a := float64(pt1.a[i]-pt2.a[i]) / 2
		b := float64(pt1.v[i]-pt2.v[i]) + a
		c := float64(pt1.p[i] - pt2.p[i])

		if a == 0 {
			if b == 0 {
				if c != 0 {
					return never
				}
				all[i] = true
			} else {
				t0[i] = -c / b
				t1[i] = -c / b
			}
		} else {
			d := math.Sqrt(b*b - 4*a*c)
			if math.IsNaN(d) {
				return never
			}
			t0[i] = (-b - d) / (2 * a)
			t1[i] = (-b + d) / (2 * a)
		}
	}

	for i := 0; i < 3; i++ {
		if !all[i] {
			if check(pt1, pt2, t0[i]) {
				return int(t0[i])
			}
			if check(pt1, pt2, t1[i]) {
				return int(t1[i])
			}
			return never
		}
	}

	return 0
}

func check(pt1, pt2 particle, t float64) bool {
	if t < 0 || t != float64(int(t)) {
		return false
	}

	t0 := int(t)
	for i := 0; i < 3; i++ {
		p1 := pt1.p[i] + t0*pt1.v[i] + t0*(t0+1)*pt1.a[i]/2
		p2 := pt2.p[i] + t0*pt2.v[i] + t0*(t0+1)*pt2.a[i]/2
		if p1 != p2 {
			return false
		}
	}

	return true
}
