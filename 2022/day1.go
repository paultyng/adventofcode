package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const topNCalorieHolders = 3

func runDay1Part2(ctx context.Context, args []string) (string, error) {
	path := "day1.input"
	if len(args) > 0 {
		path = args[0]
	}
	elves, err := readInputDay1(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	totals := make([]int, len(elves))
	for i, elf := range elves {
		totals[i] = sum(elf)
	}

	// fmt.Printf("%#v\n", totals)

	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})

	// fmt.Printf("%#v\n", totals)

	totals = totals[:topNCalorieHolders]

	return fmt.Sprintf("%d", sum(totals)), nil
}

func runDay1Part1(ctx context.Context, args []string) (string, error) {
	path := "day1.input"
	if len(args) > 0 {
		path = args[0]
	}
	elves, err := readInputDay1(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
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

	return fmt.Sprintf("%d", maxCalories), nil
}

func readInputDay1(path string) ([][]int, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	elves := [][]int{}
	elf := []int{}

	err = readLines(input, func(line string) error {
		if line == "" {
			elves = append(elves, elf)
			elf = []int{}
			return nil
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("unable to parse calories (%q): %w", line, err)
		}

		elf = append(elf, calories)

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	// if any remaining, make sure to include
	if len(elf) > 0 {
		elves = append(elves, elf)
	}

	return elves, nil
}
