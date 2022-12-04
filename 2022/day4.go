package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
)

func runDay4Part1(ctx context.Context, args []string) (string, error) {
	path := "day4.input"
	if len(args) > 0 {
		path = args[0]
	}
	pairings, err := readInputDay4(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	total := 0
	for _, pair := range pairings {
		if len(pair) != 2 {
			panic("unexpected pairing size")
		}

		if pair.HasFullDuplication() {
			total++
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func runDay4Part2(ctx context.Context, args []string) (string, error) {
	path := "day4.input"
	if len(args) > 0 {
		path = args[0]
	}
	pairings, err := readInputDay4(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	total := 0
	for _, pair := range pairings {
		if len(pair) != 2 {
			panic("unexpected pairing size")
		}

		if pair.Overlaps() {
			total++
		}
	}

	return fmt.Sprintf("%d", total), nil
}

type assignment struct {
	Start int
	End   int
}

func (a *assignment) FullyContains(b assignment) bool {
	return a.Start <= b.Start && b.End <= a.End
}

func (a *assignment) Overlaps(b assignment) bool {
	return a.Start <= b.Start && b.Start <= a.End || b.Start <= a.Start && a.Start <= b.End
}

type pairing []assignment

func (p pairing) Overlaps() bool {
	if len(p) != 2 {
		panic("unexpected pairing size")
	}

	return p[0].Overlaps(p[1])
}

func (p pairing) HasFullDuplication() bool {
	if len(p) != 2 {
		panic("unexpected pairing size")
	}

	return p[0].FullyContains(p[1]) || p[1].FullyContains(p[0])
}

func readInputDay4(path string) ([]pairing, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	lines, err := csv.NewReader(input).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	pairs := []pairing{}

	for _, line := range lines {
		pair := []assignment{}
		for _, col := range line {
			a := assignment{}
			_, err := fmt.Sscanf(col, "%d-%d", &a.Start, &a.End)
			if err != nil {
				return nil, fmt.Errorf("unable to parse assignment (%q): %w", col, err)
			}
			pair = append(pair, a)
		}
		pairs = append(pairs, pair)
	}

	return pairs, nil
}
