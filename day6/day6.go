package day6

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

func Main() {
	fmt.Println("Day 6")

	lines := shared.ReadLines("data/day6.input")
	chars := strings.Split(lines[0], "")

	start := findMarker(chars)
	fmt.Println("found", start)
}

func findMarker(chars []string) int {
	for i := 0; i < len(chars)-4; i++ {
		if isMarker(chars, i) {
			return i + 4
		}
	}
	return -1
}

func isMarker(chars []string, i int) bool {
	// fmt.Println(chars[i+0], chars[i+1], chars[i+2], chars[i+3])
	return true &&
		chars[i+0] != chars[i+1] &&
		chars[i+0] != chars[i+2] &&
		chars[i+0] != chars[i+3] &&
		chars[i+1] != chars[i+0] &&
		chars[i+1] != chars[i+2] &&
		chars[i+1] != chars[i+3] &&
		chars[i+2] != chars[i+1] &&
		chars[i+2] != chars[i+0] &&
		chars[i+2] != chars[i+3]
}
