package main

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/exp/maps"
)

func runDay3Part1(ctx context.Context, args []string) (string, error) {
	path := "day3.input"
	if len(args) > 0 {
		path = args[0]
	}
	rucksacks, err := readInputDay3(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	total := 0
	for _, r := range rucksacks {
		dupe := duplicateKey(r.compartment1, r.compartment2)
		if dupe == nil {
			panic("no duplicate found")
		}

		total += itemPriority(*dupe)
	}

	return fmt.Sprintf("%d", total), nil
}

func runDay3Part2(ctx context.Context, args []string) (string, error) {
	path := "day3.input"
	if len(args) > 0 {
		path = args[0]
	}
	rucksacks, err := readInputDay3(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	group := []set[rune]{}
	total := 0
	for _, r := range rucksacks {
		allItems := set[rune]{}
		maps.Copy(allItems, r.compartment1)
		maps.Copy(allItems, r.compartment2)
		group = append(group, allItems)
		if len(group) == 3 {
			badge := duplicateKey(group...)
			if badge == nil {
				panic("no badge")
			}
			total += itemPriority(*badge)
			group = []set[rune]{}
		}
	}

	if len(group) > 0 {
		badge := duplicateKey(group...)
		if badge == nil {
			panic("no badge")
		}
		total += itemPriority(*badge)
	}

	return fmt.Sprintf("%d", total), nil
}

type rucksack struct {
	compartment1 set[rune]
	compartment2 set[rune]
}

func itemPriority(i rune) int {
	asc := int(i)
	if asc >= 97 {
		return asc - 96
	}

	return asc - 38
}

func readRucksack(line string) rucksack {
	if len(line)%2 != 0 {
		panic(fmt.Sprintf("invalid rucksack: %q", line))
	}

	compartmentCount := len(line) / 2

	return rucksack{
		compartment1: newSet([]rune(line[:compartmentCount])...),
		compartment2: newSet([]rune(line[compartmentCount:])...),
	}
}

func readInputDay3(path string) ([]rucksack, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	rucksacks := []rucksack{}
	err = readLines(input, func(_ int, line string) error {
		rucksacks = append(rucksacks, readRucksack(line))
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	return rucksacks, nil
}
