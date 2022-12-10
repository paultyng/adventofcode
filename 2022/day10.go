package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
)

func runDay10Part1(ctx context.Context, args []string) (string, error) {
	path := "day10.input"
	if len(args) > 0 {
		path = args[0]
	}

	cycles, err := readInputDay10(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	total := 0
	var sumDuring = []int{20, 60, 100, 140, 180, 220}
	for _, cn := range sumDuring {
		total += cycles[cn-1] * cn
	}
	return fmt.Sprintf("%d", total), nil
}

func runDay10Part2(ctx context.Context, args []string) (string, error) {
	path := "day10.input"
	if len(args) > 0 {
		path = args[0]
	}
	_, err := readInputDay10(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	panic("not implemented")
}

func readInputDay10(path string) ([]int, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	return processCycles(input)
}

func processCycles(r io.Reader) ([]int, error) {
	cycles := []int{1}
	err := readLines(r, func(line string) error {
		lastX := cycles[len(cycles)-1]
		switch {
		case line == "noop":
			cycles = append(cycles, lastX)
		case strings.HasPrefix(line, "addx"):
			var v int
			_, err := fmt.Sscanf(line, "addx %d", &v)
			if err != nil {
				return fmt.Errorf("unable to parse addx: %q %w", line, err)
			}

			cycles = append(cycles, lastX, lastX+v)
		default:
			return fmt.Errorf("unknown instruction: %q", line)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	return cycles, nil
}
