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

func (f Forest) visibleToTheLeft(p Point) int {
	count := 0
	current := f.heightAt(p.x, p.y)

	for xx := p.x - 1; f.hasTree(xx, p.y); xx-- {
		count++
		if height := f.heightAt(xx, p.y); height < current {
		} else {
			return count
		}
	}
	return count
}

func (f Forest) visibleToTheRight(p Point) int {
	count := 0
	current := f.heightAt(p.x, p.y)

	for xx := p.x + 1; f.hasTree(xx, p.y); xx++ {
		count++
		if height := f.heightAt(xx, p.y); height < current {
		} else {
			return count
		}
	}
	return count
}

func (f Forest) visibleToTheTop(p Point) int {
	count := 0
	current := f.heightAt(p.x, p.y)

	for yy := p.y - 1; f.hasTree(p.x, yy); yy-- {
		count++
		if height := f.heightAt(p.x, yy); height < current {
		} else {
			return count
		}
	}
	return count
}

func (f Forest) visibleToTheBottom(p Point) int {
	count := 0
	current := f.heightAt(p.x, p.y)

	for yy := p.y + 1; f.hasTree(p.x, yy); yy++ {
		count++
		if height := f.heightAt(p.x, yy); height < current {
		} else {
			return count
		}
	}
	return count
}

func (f Forest) scenicScore(p Point) int {
	left := f.visibleToTheLeft(p)
	right := f.visibleToTheRight(p)
	top := f.visibleToTheTop(p)
	bottom := f.visibleToTheBottom(p)
	fmt.Println(p, f.heightAt(p.x, p.y), ":", "left:", left, "right:", right, "top:", top, "bottom:", bottom, "=", left*right*top*bottom)
	return left * right * top * bottom
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

	fmt.Println("Scanned forest", len(forest.trees), width, height)
	fmt.Println("")

	maxScenic := 0

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			scenicScore := forest.scenicScore(Point{x, y})
			if scenicScore > maxScenic {
				maxScenic = scenicScore
			}
			// fmt.Print(forest.heightAt(x, y))
		}
		// fmt.Print("\n")
	}

	forest.scenicScore(Point{2, 3})

	fmt.Println("\nMax scenic score:", maxScenic)

}
