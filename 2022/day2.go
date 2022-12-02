package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

func runDay2(ctx context.Context, args []string) error {
	rounds, err := readInputDay2("day2.input")
	if err != nil {
		return fmt.Errorf("unable to read input: %w", err)
	}

	total := 0

	for _, r := range rounds {
		shapeScore, outcomeScore := r.Score()
		total += shapeScore + outcomeScore
		fmt.Printf("Shape: %d, Outcome %d, Total: %d, %#v\n", shapeScore, outcomeScore, shapeScore+outcomeScore, r)
	}

	fmt.Printf("\nTotal: %d\n", total)

	return nil
}

type round struct {
	OpponentPlay   string
	DesiredOutcome string
}

func (r *round) Score() (int, int) {
	outcomeScore := 0
	suggestedPlay := ""
	switch r.DesiredOutcome {
	case "X": // lose
		outcomeScore = 0
		switch r.OpponentPlay {
		case "A": // rock
			suggestedPlay = "scissors"
		case "B": // paper
			suggestedPlay = "rock"
		case "C": // scissors
			suggestedPlay = "paper"
		}
	case "Y": // draw
		outcomeScore = 3
		switch r.OpponentPlay {
		case "A": // rock
			suggestedPlay = "rock"
		case "B": // paper
			suggestedPlay = "paper"
		case "C": // scissors
			suggestedPlay = "scissors"
		}
	case "Z": // win
		outcomeScore = 6
		switch r.OpponentPlay {
		case "A": // rock
			suggestedPlay = "paper"
		case "B": // paper
			suggestedPlay = "scissors"
		case "C": // scissors
			suggestedPlay = "rock"
		}
	}

	shapeScore := 0
	switch suggestedPlay {
	case "rock":
		shapeScore = 1
	case "paper":
		shapeScore = 2
	case "scissors":
		shapeScore = 3
	}

	return shapeScore, outcomeScore
}

func readInputDay2(path string) ([]round, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	rounds := []round{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		r := round{}
		line := scanner.Text()

		_, err := fmt.Sscan(line, &r.OpponentPlay, &r.DesiredOutcome)
		if err != nil {
			return nil, fmt.Errorf("unable to parse line: %w", err)
		}

		rounds = append(rounds, r)
	}

	return rounds, nil
}
