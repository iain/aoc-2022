package day4

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

func Main() {
	fmt.Println("Day 4")

	assignments := []Assignment{}

	lines := shared.ReadLines("data/day4.input")
	for _, line := range lines {
		assignment := lineToAssignment(line)
		assignments = append(assignments, assignment)
	}
	// fmt.Println("assignment", assignments)

	count := shared.Reduce(assignments, func(acc int, current Assignment) int {
		if current.overlaps() {
			fmt.Println("Overlap:", current)
			return acc + 1
		} else {
			fmt.Println("No Overlap:", current)
			return acc
		}
	}, 0)

	fmt.Println("count", count)
}

type Range struct {
	from int
	to   int
}

type Assignment struct {
	one Range
	two Range
}

func (a Assignment) covers() bool {
	return a.one.covers(a.two) || a.two.covers(a.one)
}

func (a Assignment) overlaps() bool {
	return a.one.overlaps(a.two) || a.two.overlaps(a.one)
}

func (r Range) covers(o Range) bool {
	return r.from <= o.from && r.to >= o.to
}

func (r Range) overlaps(o Range) bool {
	return r.from <= o.to && r.to >= o.from
}

func lineToAssignment(line string) Assignment {
	ranges := strings.Split(line, ",")
	return Assignment{
		one: toRange(ranges[0]),
		two: toRange(ranges[1]),
	}
}

func toRange(str string) Range {
	nums := strings.Split(str, "-")
	return Range{
		from: shared.StringToInt(nums[0]),
		to:   shared.StringToInt(nums[1]),
	}
}
