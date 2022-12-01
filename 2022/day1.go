package main

import (
	"context"
	"os"
	"fmt"
	"bufio"
	"strconv"
	"sort"
)

const topNCalorieHolders = 3

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func runDay1(ctx context.Context, args []string) error { 
	elves, err := readInput("day1.input")
	if err != nil {
		return fmt.Errorf("unable to read input: %w", err)
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
		
	fmt.Printf("Max calories: %#v, total: %d\n", totals, sum(totals))

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

	if len(elf) > 0 {
		elves = append(elves, elf)
	}

	return elves, nil
}