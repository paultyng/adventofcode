package main

import (
	"context"
	"fmt"
	"os"
)

func runDay19Part1(ctx context.Context, args []string) (string, error) {
	path := "day19.input"
	if len(args) > 0 {
		path = args[0]
	}
	_, err := readInputDay19(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	panic("not implemented")
}

func runDay19Part2(ctx context.Context, args []string) (string, error) {
	path := "day19.input"
	if len(args) > 0 {
		path = args[0]
	}
	_, err := readInputDay19(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	panic("not implemented")
}

func readInputDay19(path string) ([]rucksack, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	err = readLines(input, func(line string) error {
		var localVar int
		_, err := fmt.Sscan(line, &localVar)
		if err != nil {
			return fmt.Errorf("unable to parse line: %w", err)
		}

		panic("not implemented")
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	panic("not implemented")
}
