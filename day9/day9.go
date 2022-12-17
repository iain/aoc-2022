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

type Knot struct {
	index int
	point Point
}

type KnotList [10]*Knot

type History map[Point]bool

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func draw(knots KnotList, hist History) {
	var minX int = -11
	var maxX int = 15
	var minY int = -15
	var maxY int = 10

	// fmt.Printf("Drawing from %d,%d to %d,%d\n", minX, minY, maxX, maxY)
	fmt.Print("\n")

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			curr := Point{x, y}
			switch {
			case knots[0].point == curr:
				fmt.Print("H")
			case knots[1].point == curr:
				fmt.Print("1")
			case knots[2].point == curr:
				fmt.Print("2")
			case knots[3].point == curr:
				fmt.Print("3")
			case knots[4].point == curr:
				fmt.Print("4")
			case knots[5].point == curr:
				fmt.Print("5")
			case knots[6].point == curr:
				fmt.Print("6")
			case knots[7].point == curr:
				fmt.Print("7")
			case knots[8].point == curr:
				fmt.Print("8")
			case knots[9].point == curr:
				fmt.Print("9")
			case x == 0 && y == 0:
				fmt.Print("s")
			case hist[curr]:
				fmt.Print("\033[31m#\033[0m")
			default:
				fmt.Print("\033[36m.\033[0m")
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
		// fmt.Println("touching")
		return tail

	// right
	case head.x-tail.x > 1 && head.y-tail.y == 0:
		// fmt.Println("right")
		return Point{tail.x + 1, tail.y}

	// left
	case head.x-tail.x < -1 && head.y-tail.y == 0:
		// fmt.Println("left")
		return Point{tail.x - 1, tail.y}

	// up
	case head.x-tail.x == 0 && head.y-tail.y < -1:
		// fmt.Println("up")
		return Point{tail.x, tail.y - 1}

	// down
	case head.x-tail.x == 0 && head.y-tail.y > 1:
		// fmt.Println("down")
		return Point{tail.x, tail.y + 1}

	// up-right
	case head.x > tail.x && head.y < tail.y:
		// fmt.Println("up-right")
		return Point{tail.x + 1, tail.y - 1}

	// up-left
	case head.x < tail.x && head.y < tail.y:
		// fmt.Println("up-left")
		return Point{tail.x - 1, tail.y - 1}

	// down-right
	case head.x > tail.x && head.y > tail.y:
		// fmt.Println("down-left")
		return Point{tail.x + 1, tail.y + 1}

	// down-left
	case head.x < tail.x && head.y > tail.y:
		// fmt.Println("down-left")
		return Point{tail.x - 1, tail.y + 1}

	default:
		panic("unknown")
	}
}

func Main() {
	fmt.Println("Day 9")

	knots := KnotList{}

	for j := 0; j < cap(knots); j++ {
		knots[j] = &Knot{point: Point{0, 0}, index: j}
	}

	history := make(History)

	lines := shared.ReadLines("day9/full.input")
	for _, line := range lines {
		fields := strings.Fields(line)
		direction := fields[0]
		amount := shared.StringToInt(fields[1])
		fmt.Println("==", direction, amount, "==")
		for i := 0; i < amount; i++ {

			head := knots[0].point

			switch direction {
			case "R":
				head = Point{head.x + 1, head.y}
			case "L":
				head = Point{head.x - 1, head.y}
			case "U":
				head = Point{head.x, head.y - 1}
			case "D":
				head = Point{head.x, head.y + 1}
			default:
				panic("No such command: " + line)
			}
			knots[0].point = head

			for j := 0; j < len(knots)-1; j++ {
				knots[j+1].point = move(knots[j].point, knots[j+1].point)
			}

			// knots[j].point = head
			// knots[j+1].point = tail
			history[knots[9].point] = true
		}

		draw(knots, history)

	}

	fmt.Println(len(history))
}
