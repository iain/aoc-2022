package day10

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

func Main() {
	fmt.Println("Day 10")

	var values []int = []int{1}

	// current := 1

	lines := shared.ReadLines("day10/full.input")
	for _, line := range lines {
		fields := strings.Fields(line)
		switch fields[0] {
		case "addx":
			num := shared.StringToInt(fields[1])
			values = append(values, values[len(values)-1])
			values = append(values, values[len(values)-1]+num)
		case "noop":
			values = append(values, values[len(values)-1])
		default:
			panic("unknown command" + fields[0])
		}
	}

	var wanted []int = []int{20, 60, 100, 140, 180, 220}

	fmt.Println("ticks:", len(values))

	sum := 0

	for _, i := range wanted {
		strength := values[i-1] * i
		sum += strength
		fmt.Println("at:", i, "X:", values[i-1], "strength:", strength)
	}

	fmt.Println("sum:", sum)

}
