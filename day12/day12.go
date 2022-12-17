package day12

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

type Point struct {
	x int
	y int
}

type Elevation int

type Grid map[Point]Elevation

type Path struct {
	point  Point
	parent *Path
}

type Reached map[Point]bool
type Queue []Path

func Main() {
	fmt.Println("Day 12")

	grid := make(Grid)
	var end Point
	var starts []Point

	for y, line := range shared.ReadLines("day12/full.input") {
		for x, char := range strings.Split(line, "") {
			point := Point{x, y}
			if char == "S" || char == "a" {
				starts = append(starts, point)
			}
			if char == "E" {
				end = point
			}
			grid[point] = charToElevation(char)
		}
	}

	var shortest int = 1000000

	for _, start := range starts {
		steps := countSteps(grid, start, end)
		if steps == 0 {
			continue
		}
		if shortest > steps {
			shortest = steps
		}
		fmt.Println(steps, "steps from", start)
	}
	fmt.Println(shortest)
}

func countSteps(grid Grid, start, end Point) int {
	frontier := Queue{Path{point: start}}
	reached := make(Reached)

	var winning Path

	for len(frontier) > 0 {
		var current Path
		current, frontier = frontier[0], frontier[1:]
		for _, next := range getNeighbors(current, grid) {
			if next.point == end {
				winning = next
				break
			}
			if !reached[next.point] {
				frontier = append(frontier, next)
				reached[next.point] = true
			}
		}
	}

	size := 0
	var path *Path = &winning
	for path != nil {
		size++
		// fmt.Println(path.point)
		path = path.parent
	}
	return size - 1
}

func getNeighbors(current Path, grid Grid) []Path {
	var paths []Path
	from := current.point

	if to := (Point{from.x - 1, from.y + 0}); isAllowed(current, to, grid) {
		paths = append(paths, Path{point: to, parent: &current})
	}
	if to := (Point{from.x + 1, from.y + 0}); isAllowed(current, to, grid) {
		paths = append(paths, Path{point: to, parent: &current})
	}
	if to := (Point{from.x + 0, from.y + 1}); isAllowed(current, to, grid) {
		paths = append(paths, Path{point: to, parent: &current})
	}
	if to := (Point{from.x + 0, from.y - 1}); isAllowed(current, to, grid) {
		paths = append(paths, Path{point: to, parent: &current})
	}

	return paths
}

func isAllowed(current Path, to Point, grid Grid) bool {
	if _, ok := grid[to]; !ok {
		return false
	}
	return grid[to]-grid[current.point] < 2
}

func charToElevation(char string) Elevation {
	if char == "S" {
		return Elevation(0)
	}
	if char == "E" {
		return Elevation(26)
	}
	runes := []rune(char)
	ascii := int(runes[0])
	return Elevation(ascii - 97)
}
