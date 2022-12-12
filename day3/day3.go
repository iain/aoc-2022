package day3

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

// var priorities = map[string]int{
// 	"a": 1,
// }

func Main() {
	fmt.Println("Day 3")

	dupes := []string{}

	lines := shared.ReadLines("data/day3.input")
	for _, line := range lines {
		inventory := strings.Split(line, "")
		mid := len(inventory) / 2
		left := inventory[0:mid]
		right := inventory[mid:]

		dupes = append(dupes, findDupes(left, right)...)
	}

	fmt.Println("Dupes:", dupes)

	sum := 0
	for _, dupe := range dupes {
		priority := toPriority(dupe)
		fmt.Println("Priority:", dupe, priority)
		sum += priority
	}

	fmt.Println("Sum:", sum)
}

func toPriority(str string) int {
	runes := []rune(str)
	ascii := int(runes[0])
	priority := ascii - 96
	if priority < 0 {
		priority += 58
	}

	return priority
}

func contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func findDupes(left, right []string) []string {
	dupes := []string{}

	for _, needle := range left {
		if contains(right, needle) && !contains(dupes, needle) {
			dupes = append(dupes, needle)
		}
	}

	// for _, needle := range right {
	// 	if contains(left, needle) && !contains(dupes, needle) {
	// 		fmt.Println("found2:", needle)
	// 		dupes = append(dupes, needle)
	// 	}
	// }

	return dupes
}
