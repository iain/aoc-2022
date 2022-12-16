package day7

import (
	"aoc-2022/shared"
	"fmt"
	"regexp"
	"strings"
)

const diskSize = 70_000_000
const spaceNeeded = 30_000_000

type Directory struct {
	name    string
	parent  *Directory
	subdirs map[string]*Directory
	files   map[string]*File
}

func (d Directory) addSubdir(name string) {
	subdir := Directory{
		name:    name,
		parent:  &d,
		subdirs: make(map[string]*Directory),
		files:   make(map[string]*File),
	}
	allDirs = append(allDirs, &subdir)
	d.subdirs[name] = &subdir
}

func (d Directory) addFile(size int, name string) {
	file := File{name: name, size: size, directory: &d}
	d.files[name] = &file
}

func (d Directory) getSubdir(name string) *Directory {
	return d.subdirs[name]
}

func (d Directory) getSize() int {
	sum := 0
	for _, file := range d.files {
		sum += file.size
	}
	for _, dir := range d.subdirs {
		sum += dir.getSize()
	}
	return sum
}

type File struct {
	name      string
	size      int
	directory *Directory
}

var rootDir = Directory{
	name:    "/",
	subdirs: make(map[string]*Directory),
	files:   make(map[string]*File),
}

var allDirs []*Directory

func Main() {
	fmt.Println("Day 7")

	currentDir := rootDir

	lines := shared.ReadLines("data/day7.input")
	for _, line := range lines {
		// fmt.Println(line)
		switch {
		case line == "$ cd /":
			currentDir = rootDir
		case line == "$ ls":
			// fmt.Println("listing")
		case match(`^dir \w+`, line):
			words := strings.Fields(line)
			currentDir.addSubdir(words[1])
		case match(`^\d+ \w+`, line):
			words := strings.Fields(line)
			currentDir.addFile(shared.StringToInt(words[0]), words[1])
		case match(`^\$ cd \w+`, line):
			words := strings.Fields(line)
			currentDir = *currentDir.getSubdir(words[2])
		case match(`^\$ cd \.\.`, line):
			currentDir = *currentDir.parent
		default:
			panic("unknown line: " + line)
		}
	}

	totalUsed := rootDir.getSize()
	fmt.Println("total size used", totalUsed)
	fmt.Println("total size free", diskSize-totalUsed)

	deleteSize := diskSize

	for _, dir := range allDirs {
		size := dir.getSize()
		sizeFree := diskSize - totalUsed + size
		if sizeFree > spaceNeeded {
			fmt.Println("potential", dir.name, size)
			if size < deleteSize {
				deleteSize = size
			}
		}
	}

	fmt.Println("size to delete", deleteSize)
}

func match(pattern string, line string) bool {
	matched, err := regexp.Match(pattern, []byte(line))
	if err != nil {
		panic(err)
	}
	return matched
}
