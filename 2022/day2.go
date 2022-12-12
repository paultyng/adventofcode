package main

import (
	"context"
	"fmt"
	"os"
)

func runDay2Part1(ctx context.Context, args []string) (string, error) {
	path := "day2.input"
	if len(args) > 0 {
		path = args[0]
	}
	rounds, err := readInputDay2(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	total := 0

	for _, r := range rounds {
		shapeScore, outcomeScore := r.Score1()
		total += shapeScore + outcomeScore
		// fmt.Printf("Shape: %d, Outcome %d, Total: %d, %#v\n", shapeScore, outcomeScore, shapeScore+outcomeScore, r)
	}

	return fmt.Sprintf("%d", total), nil
}

func runDay2Part2(ctx context.Context, args []string) (string, error) {
	path := "day2.input"
	if len(args) > 0 {
		path = args[0]
	}
	rounds, err := readInputDay2(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	total := 0

	for _, r := range rounds {
		shapeScore, outcomeScore := r.Score2()
		total += shapeScore + outcomeScore
		// fmt.Printf("Shape: %d, Outcome %d, Total: %d, %#v\n", shapeScore, outcomeScore, shapeScore+outcomeScore, r)
	}

	return fmt.Sprintf("%d", total), nil
}

type round struct {
	OpponentPlay string
	Suggestion   string
}

func (r *round) Score1() (int, int) {
	shapeScore := 0
	switch r.Suggestion {
	case "X": // rock
		shapeScore = 1
	case "Y": // paper
		shapeScore = 2
	case "Z": // scissors
		shapeScore = 3
	}

	outcomeScore := 0
	if r.Suggestion == "X" && r.OpponentPlay == "A" || r.Suggestion == "Y" && r.OpponentPlay == "B" || r.Suggestion == "Z" && r.OpponentPlay == "C" {
		outcomeScore = 3
	} else if r.Suggestion == "X" && r.OpponentPlay == "C" || r.Suggestion == "Y" && r.OpponentPlay == "A" || r.Suggestion == "Z" && r.OpponentPlay == "B" {
		outcomeScore = 6
	}

	return shapeScore, outcomeScore
}

func (r *round) Score2() (int, int) {
	outcomeScore := 0
	suggestedPlay := ""
	switch r.Suggestion {
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
	err = readLines(input, func(_ int, line string) error {
		r := round{}
		_, err := fmt.Sscan(line, &r.OpponentPlay, &r.Suggestion)
		if err != nil {
			return fmt.Errorf("unable to parse line: %w", err)
		}

		rounds = append(rounds, r)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	return rounds, nil
}
