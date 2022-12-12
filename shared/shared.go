package shared

import (
	"bufio"
	"os"
)

// quickly get all lines of the file
func ReadLines(filename string) []string {
	lines := []string{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
