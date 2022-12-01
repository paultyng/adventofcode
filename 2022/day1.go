package main

import (
	"context"
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func runDay1(ctx context.Context, args []string) error { 
	elves, err := readInput("day1.input")
	if err != nil {
		return fmt.Errorf("unable to read input: %w", err)
	}

	maxCalories := 0
	for _, elf := range elves {
		total := 0
		for _, calories := range elf {
			total += calories
		}
		if total > maxCalories {
			maxCalories = total
		}
	}

	fmt.Printf("Max calories: %d\n", maxCalories)

	return nil
}

func readInput(path string) ([][]int, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	elves := [][]int{}
	elf := []int{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, elf)
			elf = []int{}
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("unable to parse calories (%q): %w", line, err)
		}

		elf = append(elf, calories)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	return elves, nil
}