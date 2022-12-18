package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type xy struct {
	x, y int
}

type sensor struct {
	xy
	d int
}

func (s sensor) vertices() [4]xy {
	d := s.d + 1
	return [...]xy{
		{s.x + d, s.y},
		{s.x, s.y + d},
		{s.x - d, s.y},
		{s.x, s.y - d},
	}
}

func (s sensor) contains(c xy) bool {
	dx, dy := abs(s.x-c.x), abs(s.y-c.y)
	return dx+dy <= s.d
}

func intersectSensors(s1, s2 sensor) []xy {
	s := make([]xy, 0, 2)
	v1, v2 := s1.vertices(), s2.vertices()
	for _, i := range [...][4]int{
		{1, 0, 2, 1},
		{1, 0, 3, 0},
		{2, 3, 2, 1},
		{2, 3, 3, 0},
	} {
		if c, ok := intersectLines(v1[i[0]], v1[i[1]], v2[i[2]], v2[i[3]]); ok {
			s = append(s, c)
		}
		if c, ok := intersectLines(v2[i[0]], v2[i[1]], v1[i[2]], v1[i[3]]); ok {
			s = append(s, c)
		}
	}
	return s
}

func intersectLines(a1, a2, b1, b2 xy) (xy, bool) {
	// coordinates: x points right, y points up
	// a1...a2 is from top left to bottom right
	// b1...b2 is from bottom left to top right
	// ---
	// a.x + a.y = c1
	// b.x - b.y = c2
	// ---
	// x = (c1 + c2) / 2
	// y = c1 - x
	// ---
	// c1 = a.x1 + a.y1
	// c2 = b.x1 - b.y1
	x := a1.x + a1.y + b1.x - b1.y
	if x%2 == 1 {
		return xy{}, false
	}
	x /= 2
	if x < a1.x || x > a2.x || x < b1.x || x > b2.x {
		return xy{}, false
	}

	y := a1.x + a1.y - x
	return xy{x, y}, true
}

func solve(d0 int, input string) int {
	sensors := []sensor(nil)
	scanner := bufio.NewScanner(strings.NewReader(strings.TrimSpace(input)))
	pattern := regexp.MustCompile(`Sensor at x=(\d+), y=(\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	for scanner.Scan() {
		m := pattern.FindStringSubmatch(scanner.Text())
		sx, sy, bx, by := parseInt(m[1]), parseInt(m[2]), parseInt(m[3]), parseInt(m[4])
		sensors = append(sensors, sensor{xy{sx, sy}, abs(bx-sx) + abs(by-sy)})
	}

	intersections := map[xy]int{}
	for i := 0; i < len(sensors); i++ {
		for j := i + 1; j < len(sensors); j++ {
		loop:
			for _, s := range intersectSensors(sensors[i], sensors[j]) {
				if s.x < 0 || s.x > d0 || s.y < 0 || s.y > d0 {
					continue
				}

				for k, s2 := range sensors {
					if k != i && k != j && s2.contains(s) {
						continue loop
					}
				}

				intersections[s]++
				if intersections[s] == 4 {
					return 4000000*s.x + s.y
				}
			}
		}
	}
	return 0
}

func main() {
	t0 := time.Now()
	r := solve(max, input)
	dt := time.Since(t0)
	fmt.Printf("part two: %d\n", r)
	fmt.Printf("time: %v\n", dt)
}

const max = 4000000

const input = `
Sensor at x=3859432, y=2304903: closest beacon is at x=3677247, y=3140958
Sensor at x=2488890, y=2695345: closest beacon is at x=1934788, y=2667279
Sensor at x=3901948, y=701878: closest beacon is at x=4095477, y=368031
Sensor at x=2422190, y=1775708: closest beacon is at x=1765036, y=2000000
Sensor at x=2703846, y=3282799: closest beacon is at x=2121069, y=3230302
Sensor at x=172003, y=2579074: closest beacon is at x=-77667, y=3197309
Sensor at x=1813149, y=1311283: closest beacon is at x=1765036, y=2000000
Sensor at x=1704453, y=2468117: closest beacon is at x=1934788, y=2667279
Sensor at x=1927725, y=2976002: closest beacon is at x=1934788, y=2667279
Sensor at x=3176646, y=1254463: closest beacon is at x=2946873, y=2167634
Sensor at x=2149510, y=3722117: closest beacon is at x=2121069, y=3230302
Sensor at x=3804434, y=251015: closest beacon is at x=4095477, y=368031
Sensor at x=2613561, y=3932220: closest beacon is at x=2121069, y=3230302
Sensor at x=3997794, y=3291220: closest beacon is at x=3677247, y=3140958
Sensor at x=98328, y=3675176: closest beacon is at x=-77667, y=3197309
Sensor at x=2006541, y=2259601: closest beacon is at x=1934788, y=2667279
Sensor at x=663904, y=122919: closest beacon is at x=1618552, y=-433244
Sensor at x=1116472, y=3349728: closest beacon is at x=2121069, y=3230302
Sensor at x=2810797, y=2300748: closest beacon is at x=2946873, y=2167634
Sensor at x=1760767, y=2024355: closest beacon is at x=1765036, y=2000000
Sensor at x=3098487, y=2529092: closest beacon is at x=2946873, y=2167634
Sensor at x=1716839, y=634872: closest beacon is at x=1618552, y=-433244
Sensor at x=9323, y=979154: closest beacon is at x=-245599, y=778791
Sensor at x=1737623, y=2032367: closest beacon is at x=1765036, y=2000000
Sensor at x=26695, y=3049071: closest beacon is at x=-77667, y=3197309
Sensor at x=3691492, y=3766350: closest beacon is at x=3677247, y=3140958
Sensor at x=730556, y=1657010: closest beacon is at x=1765036, y=2000000
Sensor at x=506169, y=3958647: closest beacon is at x=-77667, y=3197309
Sensor at x=2728744, y=23398: closest beacon is at x=1618552, y=-433244
Sensor at x=3215227, y=3077078: closest beacon is at x=3677247, y=3140958
Sensor at x=2209379, y=3030851: closest beacon is at x=2121069, y=3230302
`
