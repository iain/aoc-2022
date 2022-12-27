package day15

import (
	"aoc-2022/shared"
	"fmt"
)

type Point struct {
	x int
	y int
}

type Coords map[Point]bool

type Area struct {
	beacons Coords
	sensors Coords
	signals Coords
}

func Main() {
	fmt.Println("Day 15")
	targetRow := 2_000_000
	a := parse("day15/full.input", targetRow)
	draw(a)
	count := 0
	for point := range a.signals {
		if point.y == targetRow {
			count++
		}
	}
	for point := range a.beacons {
		if point.y == targetRow {
			count--
		}
	}
	fmt.Println(count)
}

func parse(filename string, targetRow int) Area {
	a := Area{beacons: make(Coords), sensors: make(Coords), signals: make(Coords)}
	for _, line := range shared.ReadLines(filename) {
		var sx, sy, bx, by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensor := Point{sx, sy}
		beacon := Point{bx, by}
		a.sensors[sensor] = true
		a.beacons[beacon] = true
		d := distance(sensor, beacon)
		fmt.Println("Found sensor", sx, sy, "found beacon", bx, by, "distance", d)

		y := targetRow
		// for y := (sensor.y - d); y <= (sensor.y + d); y++ {
		for x := (sensor.x - d); x <= (sensor.x + d); x++ {
			signal := Point{x, y}
			if distance(sensor, signal) <= d {
				a.signals[signal] = true
			}
		}
		// }
	}
	return a
}

func draw(a Area) {
	fmt.Println("sensors", a.sensors)
	fmt.Println("beacons", a.beacons)
	for y := -5; y < 30; y++ {
		for x := -5; x < 30; x++ {
			p := Point{x, y}
			switch {
			case a.beacons[p]:
				pr("B", 31)
			case a.sensors[p]:
				pr("S", 32)
			case a.signals[p]:
				pr("#", 33)
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
