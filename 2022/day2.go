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
	OpponentPlay  string
	SuggestedPlay string
}

func (r *round) Score() (int, int) {
	shapeScore := 0
	switch r.SuggestedPlay {
	case "X": // rock
		shapeScore = 1
	case "Y": // paper
		shapeScore = 2
	case "Z": // scissors
		shapeScore = 3
	}

	outcomeScore := 0
	if r.SuggestedPlay == "X" && r.OpponentPlay == "A" || r.SuggestedPlay == "Y" && r.OpponentPlay == "B" || r.SuggestedPlay == "Z" && r.OpponentPlay == "C" {
		outcomeScore = 3
	} else if r.SuggestedPlay == "X" && r.OpponentPlay == "C" || r.SuggestedPlay == "Y" && r.OpponentPlay == "A" || r.SuggestedPlay == "Z" && r.OpponentPlay == "B" {
		outcomeScore = 6
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

		_, err := fmt.Sscan(line, &r.OpponentPlay, &r.SuggestedPlay)
		if err != nil {
			return nil, fmt.Errorf("unable to parse line: %w", err)
		}

		rounds = append(rounds, r)
	}

	return rounds, nil
}
