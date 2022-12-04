package main

import (
	"context"
	"fmt"
	"os"
)

func runDay4Part1(ctx context.Context, args []string) (string, error) {
	path := "day4.input"
	if len(args) > 0 {
		path = args[0]
	}
	_, err := readInputDay4(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	panic("not implemented")
}

func runDay4Part2(ctx context.Context, args []string) (string, error) {
	path := "day4.input"
	if len(args) > 0 {
		path = args[0]
	}
	_, err := readInputDay4(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	panic("not implemented")
}

func readInputDay4(path string) ([]struct{}, error) {
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
