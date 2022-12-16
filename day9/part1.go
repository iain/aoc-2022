package day9

import (
	"aoc-2022/shared"
	"fmt"
	"math"
	"strings"
)

type Point struct {
	x int
	y int
}

type History map[Point]bool

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func draw(head, tail Point, hist History) {
	var minX int = 0
	var maxX int = 0
	var minY int = 0
	var maxY int = 0

	for p := range hist {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	if head.x < minX {
		minX = head.x
	}
	if head.x > maxX {
		maxX = head.x
	}
	if head.y < minY {
		minY = head.y
	}
	if head.y > maxY {
		maxY = head.y
	}
	if tail.x < minX {
		minX = tail.x
	}
	if tail.x > maxX {
		maxX = tail.x
	}
	if tail.y < minY {
		minY = tail.y
	}
	if tail.y > maxY {
		maxY = tail.y
	}

	// fmt.Printf("Drawing from %d,%d to %d,%d\n", minX, minY, maxX, maxY)
	fmt.Print("\n")

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			switch {
			case head == Point{x, y}:
				fmt.Print("H")
			case tail == Point{x, y}:
				fmt.Print("T")
			case x == 0 && y == 0:
				fmt.Print("s")
			default:
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func norm(val int) int {
	if val > 2 {
		return 1
	}
	if val < -2 {
		return -1
	}
	return 0
}

func move(head Point, tail Point) Point {

	switch {

	// still touching
	case abs(head.x-tail.x) < 2 && abs(head.y-tail.y) < 2:
		fmt.Println("touching")
		return tail

	// right
	case head.x-tail.x > 1 && head.y-tail.y == 0:
		fmt.Println("right")
		return Point{tail.x + 1, tail.y}

	// left
	case head.x-tail.x < -1 && head.y-tail.y == 0:
		fmt.Println("left")
		return Point{tail.x - 1, tail.y}

	// up
	case head.x-tail.x == 0 && head.y-tail.y < -1:
		fmt.Println("up")
		return Point{tail.x, tail.y - 1}

	// down
	case head.x-tail.x == 0 && head.y-tail.y > 1:
		fmt.Println("down")
		return Point{tail.x, tail.y + 1}

	// up-right
	case head.x > tail.x && head.y < tail.y:
		fmt.Println("up-right")
		return Point{tail.x + 1, tail.y - 1}

	// up-left
	case head.x < tail.x && head.y < tail.y:
		fmt.Println("up-left")
		return Point{tail.x - 1, tail.y - 1}

	// down-right
	case head.x > tail.x && head.y > tail.y:
		fmt.Println("down-left")
		return Point{tail.x + 1, tail.y + 1}

	// down-left
	case head.x < tail.x && head.y > tail.y:
		fmt.Println("down-left")
		return Point{tail.x - 1, tail.y + 1}

	default:
		panic("unknown")
	}
}

func Main() {
	fmt.Println("Day 9")

	head := Point{0, 0}
	tail := Point{0, 0}
	history := make(History)

	lines := shared.ReadLines("day9/full.input")
	for _, line := range lines {
		fields := strings.Fields(line)
		direction := fields[0]
		amount := shared.StringToInt(fields[1])
		fmt.Println("dir", direction, "amount", amount)
		for i := 0; i < amount; i++ {
			switch direction {
			case "R":
				head = Point{head.x + 1, head.y}
				tail = move(head, tail)
			case "L":
				head = Point{head.x - 1, head.y}
				tail = move(head, tail)
			case "U":
				head = Point{head.x, head.y - 1}
				tail = move(head, tail)
			case "D":
				head = Point{head.x, head.y + 1}
				tail = move(head, tail)
			default:
				panic("No such command: " + line)
			}
			history[tail] = true
			// draw(head, tail, history)
		}
	}

	fmt.Println(len(history))
}
