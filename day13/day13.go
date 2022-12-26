package day13

import (
	"aoc-2022/shared"
	"encoding/json"
	"fmt"
)

type Pair struct {
	left  Packet
	right Packet
	index int
}

func (p Pair) bothIntegers() bool {
	return p.left.isInteger() && p.right.isInteger()
}

type Packet struct {
	list []*Packet
	val  *int
	desc string
}

func (p Packet) isInteger() bool {
	return p.val != nil
}

func (p Packet) isList() bool {
	return p.val == nil
}

func (p Packet) to_i() int {
	return *p.val
}

func Main() {
	fmt.Println("Day 13")
	pairs := parse()
	fmt.Println("----------------------")
	sum := compareAll(pairs)
	fmt.Println(sum)
}

func compareAll(pair []Pair) int {
	sum := 0
	for _, pair := range pair {
		fmt.Println("== Pair", pair.index, "==")
		if isCorrectOrder(pair.left, pair.right) {
			fmt.Println("\033[32mPair", pair.index, "is in correct order\033[0m")
			sum += pair.index
		} else {
			fmt.Println("\033[31mPair", pair.index, "is in the wrong order\033[0m")
		}
	}
	return sum
}

func isCorrectOrder(left, right Packet) bool {
	fmt.Println("- Compare", left.desc, "vs", right.desc)

	if left.isList() && len(left.list) == 0 {
		fmt.Println("Left is empty, so in order")
		return true
	}

	for i := range left.list {
		if len(right.list) <= i {
			fmt.Println("  - Right ran out of items")
			return false
		}

		leftItem := *left.list[i]
		rightItem := *right.list[i]

		switch {
		case leftItem.isInteger() && rightItem.isInteger():
			if leftItem.to_i() < rightItem.to_i() {
				fmt.Println("   - Correct order")
				return true
			}
			if leftItem.to_i() > rightItem.to_i() {
				fmt.Println("   - Incorrect order")
				return false
			}
		case !leftItem.isInteger() && rightItem.isInteger():
			newRight := Packet{val: rightItem.val, desc: fmt.Sprintf("{%d}", rightItem.to_i())}
			return isCorrectOrder(leftItem, newRight)
		case leftItem.isInteger() && !rightItem.isInteger():
			newLeft := Packet{val: leftItem.val, desc: fmt.Sprintf("{%d}", leftItem.to_i())}
			return isCorrectOrder(newLeft, rightItem)
		default:
			fmt.Println("Both are lists")
			return isCorrectOrder(leftItem, rightItem)
		}
	}

	fmt.Println("   - Same item?", left.desc, "vs", right.desc)
	return true
}

func parse() []Pair {
	lines := shared.ReadLines("day13/sample.input")

	size := int((len(lines) + 1) / 3)

	var pairs []Pair = make([]Pair, size)

	for i := 0; i < size; i++ {
		left := lines[i*3]
		right := lines[(i*3)+1]

		pairs[i] = Pair{
			left:  parsePacket(left),
			right: parsePacket(right),
			index: i + 1,
		}
	}
	return pairs
}

func parsePacket(input string) Packet {
	var result []interface{}

	err := json.Unmarshal([]byte(input), &result)

	if err != nil {
		panic(err)
	}

	return *recursivePacket(result, input)
}

func recursivePacket(item interface{}, input string) *Packet {
	switch t := item.(type) {
	case float64:
		val := int(item.(float64))
		return &Packet{
			val:  &val,
			desc: v(val),
		}
	case []interface{}:
		list := item.([]interface{})
		var packets []*Packet = make([]*Packet, len(list))
		for i, subitem := range list {
			packets[i] = recursivePacket(subitem, v(subitem))
		}
		return &Packet{list: packets, desc: input}
	default:
		panic(fmt.Sprintf("oops, unknown type: %T", t))
	}
}

func v(x interface{}) string {
	return fmt.Sprintf("%v", x)
}
