package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Start")
	fmt.Println("")

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

	fmt.Println("")
	fmt.Println("Exit")
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

func to_int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func GetCalories(filename string, ch chan Elf) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elf := Elf{}

	for scanner.Scan() {
		if line := scanner.Text(); line == "" {
			// send complete elf data
			ch <- elf
			// make a new elf
			elf = Elf{}
		} else {
			i := to_int(line)
			foodItem := FoodItem{calories: i}
			elf.foodItems = append(elf.foodItems, foodItem)
		}
	}

	// flush last elf
	if len(elf.foodItems) > 0 {
		ch <- elf
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	close(ch)
}
