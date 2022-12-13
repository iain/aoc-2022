package day1

import (
	"aoc-2022/shared"
	"fmt"
	"sort"
)

func Main() {
	ch := make(chan Elf)

	go GetCalories("data/day1.input", ch)

	var elves Expedition

	for v := range ch {
		// fmt.Println(v)
		elves = append(elves, v)
	}

	max := maxCalories(elves)
	fmt.Println("Max:", max)

	best3 := topThree(elves)
	fmt.Println("Top", len(best3), best3)
	fmt.Println("Max Top 3:", sumCalories(best3))
}

type FoodItem struct {
	calories int
}

type Elf struct {
	foodItems []FoodItem
}

type Expedition []Elf

func (elf Elf) Calories() int {
	sum := 0
	for _, foodItem := range elf.foodItems {
		sum += foodItem.calories
	}
	return sum
}

func maxCalories(elves Expedition) int {
	var max int
	for _, elf := range elves {
		if cal := elf.Calories(); cal > max {
			max = cal
		}
	}
	return max
}

func sumCalories(elves Expedition) int {
	var sum int
	for _, elf := range elves {
		sum += elf.Calories()
	}
	return sum
}

func topThree(elves Expedition) []Elf {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Calories() > elves[j].Calories()
	})
	return elves[0:3]
}

func GetCalories(filename string, ch chan Elf) {
	lines := shared.ReadLines(filename)

	elf := Elf{}

	for _, line := range lines {
		if line == "" {
			// send complete elf data
			ch <- elf
			// make a new elf
			elf = Elf{}
		} else {
			i := shared.StringToInt(line)
			foodItem := FoodItem{calories: i}
			elf.foodItems = append(elf.foodItems, foodItem)
		}
	}

	// flush last elf
	if len(elf.foodItems) > 0 {
		ch <- elf
	}

	close(ch)
}
