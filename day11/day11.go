package day11

import (
	"aoc-2022/shared"
	"fmt"
	"sort"
	"strings"
)

type MonkeyIndex int64
type Operation string
type Operand int64
type Test int64
type MonkeyTarget map[bool]MonkeyIndex
type WorryLevel uint64

type Monkey struct {
	index       MonkeyIndex
	items       []WorryLevel
	operation   Operation
	operand     Operand
	test        Test
	targets     MonkeyTarget
	inspections int
}

type MonkeyList map[MonkeyIndex]*Monkey

func Main() {
	fmt.Println("Day 11")

	monkeyList := readMonkeys("day11/full.input")

	modulation := findModulation(&monkeyList)
	fmt.Println(modulation)

	for i := 0; i < 10_000; i++ {
		playRound(&monkeyList, modulation)
		if i%1000 == 0 {
			fmt.Println("Playing round", i)
			for _, monkey := range monkeyList {
				fmt.Println(monkey)
			}
		}
	}

	fmt.Println("\nDone playing")
	var business []int
	for _, monkey := range monkeyList {
		fmt.Println(monkey)
		business = append(business, monkey.inspections)
	}
	sort.Ints(business)
	fmt.Println(business)
	fmt.Println(business[len(business)-1] * business[len(business)-2])
}

func findModulation(monkeyList *MonkeyList) WorryLevel {
	product := 1
	for _, monkey := range *monkeyList {
		product *= int(monkey.test)
	}
	return WorryLevel(product)
}

func playRound(monkeyList *MonkeyList, modulation WorryLevel) {
	for i := 0; i < len(*monkeyList); i++ {
		monkey := (*monkeyList)[MonkeyIndex(i)]
		// fmt.Println("-- monkey", monkey.index)
		for len(monkey.items) > 0 {
			var item, newItem WorryLevel

			monkey.inspections++
			item, monkey.items = monkey.items[0], monkey.items[1:]
			operand := WorryLevel(monkey.operand)
			if monkey.operand == Operand(-1) {
				operand = WorryLevel(item)
			}
			switch monkey.operation {
			case "*":
				newItem = item * operand
			case "+":
				newItem = item + operand
			default:
				panic("unknown operation " + monkey.operation)
			}
			// item = item / 3
			if item > newItem {
				panic("wraparound, old: " + fmt.Sprint(item) + ", new: " + fmt.Sprint(newItem))
			}
			newItem = newItem % modulation
			testResult := uint64(newItem) % uint64(monkey.test)
			targetIndex := monkey.targets[testResult == 0]
			// fmt.Println(">", testResult, targetIndex, item)
			targetMonkey := (*monkeyList)[targetIndex]
			targetMonkey.items = append(targetMonkey.items, newItem)
		}
	}
}

func readMonkeys(filename string) MonkeyList {
	monkeyList := make(MonkeyList)
	currentMonkeyIndex := MonkeyIndex(0)

	for _, line := range shared.ReadLines(filename) {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		switch fields[0] {
		case "Monkey":
			currentMonkeyIndex = MonkeyIndex(shared.StringToInt(strings.ReplaceAll(fields[1], ":", "")))
			monkey := Monkey{index: currentMonkeyIndex, targets: make(MonkeyTarget)}
			monkeyList[currentMonkeyIndex] = &monkey
		case "Starting":
			monkey := monkeyList[currentMonkeyIndex]
			monkey.items = parseWorryList(fields[2:])
		case "Operation:":
			monkey := monkeyList[currentMonkeyIndex]
			operation := fields[4]
			operand := fields[5]
			if operand == "old" {
				monkey.operand = Operand(-1)
			} else {
				monkey.operand = Operand(shared.StringToInt(operand))
			}
			monkey.operation = Operation(operation)
		case "Test:":
			monkey := monkeyList[currentMonkeyIndex]
			monkey.test = Test(shared.StringToInt(fields[3]))
		case "If":
			monkey := monkeyList[currentMonkeyIndex]
			index := MonkeyIndex(shared.StringToInt(fields[5]))
			if fields[1] == "true:" {
				monkey.targets[true] = index
			} else {
				monkey.targets[false] = index
			}
		default:
			panic("unknown line: " + line)
		}
	}

	return monkeyList
}

func parseWorryList(words []string) []WorryLevel {
	list := make([]WorryLevel, len(words))
	for i, word := range words {
		list[i] = parseSingleWorry(word)
	}
	return list
}

func parseSingleWorry(word string) WorryLevel {
	var level int64 = int64(shared.StringToInt(strings.ReplaceAll(word, ",", "")))
	return WorryLevel(level)
}
