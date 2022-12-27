package day15

import (
	"aoc-2022/shared"
	"fmt"
)

const input = "day15/full.input"
const max_x = 4_000_000
const max_y = 4_000_000

type Point struct {
	x int
	y int
}

type Coords map[Point]bool

type Sensor map[Point]Point

type Area struct {
	beacons Coords
	sensors Sensor
	signals Coords
}

func Main() {
	fmt.Println("Day 15")
	a := parse(input)

	// draw(a)

	for sensor, beacon := range a.sensors {
		d := distance(sensor, beacon)
		for i := 0; i <= d; i++ {
			for _, p := range points(sensor, d+1, i) {
				// fmt.Println("Checking", p)
				if a.isHiddenBeacon(p) {
					panic(fmt.Sprintf("FOUND IT: %v, freq: %d", p, tuningFrequency(p)))
				}
			}
		}
	}
}

func tuningFrequency(p Point) int64 {
	return int64(p.x)*int64(max_x) + int64(p.y)
}

func points(s Point, d, i int) [4]Point {
	o := d - i

	var points = [4]Point{
		(Point{s.x + i, s.y - o}),
		(Point{s.x + i, s.y + o}),
		(Point{s.x - i, s.y - o}),
		(Point{s.x - i, s.y + o}),
	}

	return points
}

func (a *Area) isSensor(p Point) bool {
	_, ok := a.sensors[p]
	return ok
}
func (a *Area) isBeacon(p Point) bool {
	_, ok := a.beacons[p]
	return ok
}

func (a *Area) isHiddenBeacon(p Point) bool {
	if p.x < 0 || p.y < 0 || p.x > max_x || p.y > max_y {
		return false
	}
	for sensor, beacon := range a.sensors {
		d1 := distance(sensor, beacon)
		d2 := distance(sensor, p)
		if d2 <= d1 {
			return false
		}
	}
	return true
}

func parse(filename string) Area {
	a := Area{beacons: make(Coords), sensors: make(Sensor), signals: make(Coords)}
	for _, line := range shared.ReadLines(filename) {
		var sx, sy, bx, by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensor := Point{sx, sy}
		beacon := Point{bx, by}
		a.sensors[sensor] = beacon
		a.beacons[beacon] = true
		// d := distance(sensor, beacon)
		// fmt.Println("Found sensor", sx, sy, "found beacon", bx, by, "distance", d)

		// for y := (sensor.y - d); y <= (sensor.y + d); y++ {
		// 	for x := (sensor.x - d); x <= (sensor.x + d); x++ {
		// 		signal := Point{x, y}
		// 		if distance(sensor, signal) <= d {
		// 			a.signals[signal] = true
		// 		}
		// 	}
		// }
	}
	return a
}

func draw(a Area) {
	fmt.Println("sensors", a.sensors)
	fmt.Println("beacons", a.beacons)
	o := Point{14, 11}
	for y := -10; y < 30; y++ {
		for x := -10; x < 30; x++ {
			p := Point{x, y}
			switch {
			case a.isBeacon(p):
				pr("B", 31)
			case a.isSensor(p):
				pr("S", 32)
			case a.signals[p]:
				pr("░", 33)
			case p == o:
				pr("╳", 34)
			default:
				pr(" ", 0)
			}
		}
		fmt.Print("\n")
	}
}

func pr(str string, color int) {
	fmt.Printf("\033[%dm%s\033[0m", color, str)
}

func distance(a, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	} else {
		return v
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
