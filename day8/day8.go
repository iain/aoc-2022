package day8

import (
	"aoc-2022/shared"
	"fmt"
	"strings"
)

type Point struct {
	x int
	y int
}
type Height int

type Forest struct {
	trees map[Point]Height
}

func (f Forest) heightAt(x, y int) Height {
	return f.trees[Point{x, y}]
}

func (f Forest) hasTree(x, y int) bool {
	_, ok := f.trees[Point{x, y}]
	return ok
}

func (f Forest) visibleToTheLeft(p Point) bool {
	for xx := p.x - 1; f.hasTree(xx, p.y); xx-- {
		if f.trees[p] <= f.heightAt(xx, p.y) {
			return false
		}
	}
	return true
}

func (f Forest) visibleToTheRight(p Point) bool {
	for xx := p.x + 1; f.hasTree(xx, p.y); xx++ {
		if f.trees[p] <= f.heightAt(xx, p.y) {
			return false
		}
	}
	return true
}

func (f Forest) visibleToTheTop(p Point) bool {
	for yy := p.y - 1; f.hasTree(p.x, yy); yy-- {
		if f.trees[p] <= f.heightAt(p.x, yy) {
			return false
		}
	}
	return true
}

func (f Forest) visibleToTheBottom(p Point) bool {
	for yy := p.y + 1; f.hasTree(p.x, yy); yy++ {
		if f.trees[p] <= f.heightAt(p.x, yy) {
			return false
		}
	}
	return true
}

func (f Forest) isVisible(p Point) bool {
	return f.visibleToTheLeft(p) ||
		f.visibleToTheRight(p) ||
		f.visibleToTheTop(p) ||
		f.visibleToTheBottom(p)
}

func Main() {
	fmt.Println("Day 8")

	forest := Forest{}
	forest.trees = make(map[Point]Height)
	width := 0
	height := 0

	lines := shared.ReadLines("data/day8.input")
	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			forest.trees[Point{x, y}] = Height(shared.StringToInt(char))
			width = x + 1
			height = y + 1
		}
	}

	fmt.Println("Scanned forest", len(forest.trees))
	fmt.Println("")

	numVisible := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if forest.isVisible(Point{x, y}) {
				fmt.Print("\033[32m")
				numVisible++
			} else {
				fmt.Print("\033[31m")
			}
			fmt.Print(forest.heightAt(x, y))
		}
		fmt.Print("\033[0m\n")
	}

	fmt.Println("\nFound visible trees in forest:", numVisible)

}
