package day14

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

type Point struct {
	x int
	y int
}

type Coords map[Point]bool

type Formation struct {
	rocks Coords
	sands Coords
	minX  int
	minY  int
	maxX  int
	maxY  int
}

func note(f *Formation, x, y int) {
	if x < f.minX {
		f.minX = x
	}
	if x > f.maxX {
		f.maxX = x
	}
	if y < f.minY {
		f.minY = y
	}
	if y > f.maxY {
		f.maxY = y
	}
}

func (f Formation) rock(x, y int) bool {
	// fmt.Println("check rock at", x, y)
	return f.rocks[Point{x, y}]
}

func (f Formation) sand(x, y int) bool {
	return f.sands[Point{x, y}]
}

func (f Formation) floor(x, y int) bool {
	return y == f.maxY+2
}

func (f Formation) occupied(x, y int) bool {
	return f.sand(x, y) || f.rock(x, y) || f.floor(x, y)
}

func Main() {
	fmt.Println("Day 14")
	formation := parse()
	draw(formation)
	for {
		// time.Sleep(1 * time.Millisecond)
		tick(formation)
		// draw(formation)
	}
}

func tick(f Formation) {
	s := Point{500, 0}
	for {
		// if s.y > f.maxY {
		// 	panic(fmt.Sprintf("done! %d sands", len(f.sands)))
		// }
		if f.occupied(s.x, s.y+1) {
			if f.occupied(s.x-1, s.y+1) {
				if f.occupied(s.x+1, s.y+1) {
					// no more place to go
					f.sands[s] = true
					if s.x == 500 && s.y == 0 {
						panic(fmt.Sprintf("done! %d sands", len(f.sands)))
					}
					break
				} else {
					s = Point{s.x + 1, s.y + 1}
				}
			} else {
				s = Point{s.x - 1, s.y + 1}
			}
		} else {
			s = Point{s.x, s.y + 1}
		}
	}
}

func draw(f Formation) {
	fmt.Print("\033[2J")
	for y := f.minY - 1; y <= f.maxY+2; y++ {
		for x := f.minX - 1; x <= f.maxX+1; x++ {
			if x == 500 && y == 0 {
				fmt.Print("\033[32m+\033[0m")
			} else if f.sand(x, y) {
				fmt.Print("\033[33m░\033[0m")
			} else if f.floor(x, y) {
				fmt.Print("\033[32m█\033[0m")
			} else if f.rock(x, y) {
				fmt.Print("\033[31m█\033[0m")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func parse() Formation {
	f := Formation{minX: 500, minY: 0, rocks: make(Coords), sands: make(Coords)}
	for _, line := range shared.ReadLines("day14/full.input") {
		coords := strings.Split(line, " -> ")
		for i := 1; i < len(coords); i++ {

			var x1, y1 int
			fmt.Sscanf(coords[i-1], "%d,%d", &x1, &y1)

			var x2, y2 int
			fmt.Sscanf(coords[i-0], "%d,%d", &x2, &y2)

			// fmt.Println("Line", x1, y1, "to", x2, y2)

			for x := min(x1, x2); x <= max(x1, x2); x++ {
				for y := min(y1, y2); y <= max(y1, y2); y++ {
					point := Point{x, y}
					note(&f, x, y)
					// fmt.Println("adding rock", point)
					f.rocks[point] = true
				}
			}
		}
	}
	return f
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
