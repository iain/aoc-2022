package day10

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

func Main() {
	fmt.Println("Day 10")

	var values []int = []int{}

	current := 1

	lines := shared.ReadLines("day10/full.input")
	for _, line := range lines {
		fields := strings.Fields(line)
		switch fields[0] {
		case "addx":
			values = append(values, current)
			values = append(values, current)
			num := shared.StringToInt(fields[1])
			current += num
		case "noop":
			values = append(values, current)
		default:
			panic("unknown command" + fields[0])
		}
	}

	fmt.Println("ticks:", len(values), values)

	Part2(values)

}

func Part2(values []int) {
	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {
			spriteLocation := values[(row*40)+col]

			if col+1 < spriteLocation+3 && col+1 >= spriteLocation {
				// fmt.Println("# col", col+1, "loc", spriteLocation)
				fmt.Print("#")
			} else {
				// fmt.Println(". col", col+1, "loc", spriteLocation)
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func Part1(values []int) {
	var wanted []int = []int{20, 60, 100, 140, 180, 220}
	sum := 0

	for _, i := range wanted {
		strength := values[i-1] * i
		sum += strength
		fmt.Println("at:", i, "X:", values[i-1], "strength:", strength)
	}

	fmt.Println("sum:", sum)
}
