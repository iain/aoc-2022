package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Hand struct {
	points int
	name   string
}

var Rock = Hand{points: 1, name: "Rock"}
var Paper = Hand{points: 2, name: "Paper"}
var Scissors = Hand{points: 3, name: "Scissors"}

var mappings = map[string]Hand{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var beats = map[Hand]Hand{
	Rock:     Scissors,
	Scissors: Paper,
	Paper:    Rock,
}

type Round struct {
	Opponent Hand
	You      Hand
}

func (r Round) points() int {
	return r.winPoints() + r.You.points
}

func (r Round) winPoints() int {
	if r.Opponent == r.You {
		return 3
	} else {
		if beats[r.You] == r.Opponent {
			return 6
		} else {
			return 0
		}
	}
}

func Main() {
	rounds := getRounds("data/day2.input")
	fmt.Println("Rounds:", rounds)

	sum := 0

	for _, round := range rounds {
		points := round.points()
		fmt.Println("Round:", round, "Points:", points)
		sum += points
	}

	fmt.Println("Sum:", sum)
}

func getRounds(filename string) []Round {
	rounds := []Round{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if line := scanner.Text(); line == "" {
			// one
		} else {
			// two
			abs := strings.Fields(line)
			opp := mappings[abs[0]]
			you := mappings[abs[1]]
			rounds = append(rounds, Round{Opponent: opp, You: you})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return rounds
}
