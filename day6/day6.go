package day6

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

const markerSize = 14

func Main() {
	fmt.Println("Day 6")

	lines := shared.ReadLines("data/day6.input")
	chars := strings.Split(lines[0], "")

	start := findMarker(chars)
	fmt.Println("found", start)
}

func findMarker(chars []string) int {
	for i := 0; i < len(chars)-markerSize; i++ {
		if isMarker(chars, i) {
			return i + markerSize
		}
	}
	return -1
}

func isMarker(chars []string, i int) bool {
	fmt.Println(chars[i : markerSize+i])
	uniq := map[string]bool{}

	for _, char := range chars[i : markerSize+i] {
		uniq[char] = true
	}

	return len(uniq) == markerSize
}
