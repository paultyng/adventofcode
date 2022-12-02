package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

func runDay3Part1(ctx context.Context, args []string) (string, error) {
	path := "day3.input"
	if len(args) > 0 {
		path = args[0]
	}
	_, err := readInputDay3(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	panic("not implemented")
}

func runDay3Part2(ctx context.Context, args []string) (string, error) {
	path := "day3.input"
	if len(args) > 0 {
		path = args[0]
	}
	_, err := readInputDay3(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	panic("not implemented")
}

func readInputDay3(path string) ([]int, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	// rounds := []round{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		// r := round{}
		// line := scanner.Text()

		// _, err := fmt.Sscan(line, &r.OpponentPlay, &r.Suggestion)
		// if err != nil {
		// 	return nil, fmt.Errorf("unable to parse line: %w", err)
		// }

		// rounds = append(rounds, r)
	}

	panic("not implemented")
}
