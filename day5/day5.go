package day5

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

type Crate struct {
	name  string
	below *Crate
}

var stacks = [9]*Crate{}

func Main() {
	fmt.Println("Day 5")

	mode := 1

	lines := shared.ReadLines("data/day5.input")
	for _, line := range lines {
		if getChar(line, 1) == "1" {
			draw()
			fmt.Println("----")
			mode = 2
		}

		if getChar(line, 0) == "m" {
			mode = 3
		}

		if mode == 1 {
			readStacks(line)
		}

		if mode == 3 {
			doMoveFromLine(line)
		}
	}

	draw()
	outcome()
}

func outcome() {
	str := ""
	for i := 0; i < cap(stacks); i++ {
		str += (*stacks[i]).name
	}
	fmt.Println(str)
}

func readStacks(line string) {
	for i := 0; i < cap(stacks); i++ {
		if name := getChar(line, (i*4)+1); name != " " {
			if stacks[i] == nil {
				stacks[i] = &Crate{name: name}
			} else {
				addToTail(stacks[i], Crate{name: name})
			}
		}
	}
}

func getChar(line string, i int) string {
	if len(line) > i {
		return string(line[i])
	} else {
		return " "
	}
}

func addToTail(c *Crate, new_tail Crate) {
	tail := c
	for tail.below != nil {
		tail = *&tail.below
	}
	tail.below = &new_tail
}

func draw() {
	for i := 0; i < cap(stacks); i++ {
		if stacks[i] == nil {
			fmt.Println("||")
		} else {
			tail := *stacks[i]
			str := tail.name
			for tail.below != nil {
				tail = *tail.below
				str += tail.name
			}
			fmt.Println("|", str)
		}
	}
}

func doMoveFromLine(line string) {
	fields := strings.Fields(line)
	count := shared.StringToInt(fields[1])
	source := shared.StringToInt(fields[3])
	target := shared.StringToInt(fields[5])
	doMove(count, source-1, target-1)
}

func doMove(count, source, target int) {
	// fmt.Println("moving", count, "from", source, "to", target)
	for i := 0; i < count; i++ {
		doSingleMove(source, target)
	}
	// draw()
}

func doSingleMove(source, target int) {
	currentCrate := *stacks[source]
	stacks[source] = currentCrate.below
	currentCrate.below = stacks[target]
	stacks[target] = &currentCrate
}
