package main

import (
	"bufio"
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
		dupe := r.DuplicateItem()

		total += itemPriority(dupe)
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

	group := []map[string]bool{}
	total := 0
	for _, r := range rucksacks {
		allItems := map[string]bool{}
		maps.Copy(allItems, r.compartment1)
		maps.Copy(allItems, r.compartment2)
		group = append(group, allItems)
		if len(group) == 3 {
			badge := duplicateKey(group...)
			if badge == nil {
				panic("no badge")
			}
			total += itemPriority(*badge)
			group = []map[string]bool{}
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
	compartment1 map[string]bool
	compartment2 map[string]bool
}

func (r *rucksack) DuplicateItem() string {
	d := duplicateKey(r.compartment1, r.compartment2)
	if d == nil {
		return ""
	}
	return *d
}

func itemPriority(i string) int {
	asc := int(i[0])
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
		compartment1: characterMap(line[:compartmentCount]),
		compartment2: characterMap(line[compartmentCount:]),
	}
}

func characterMap(s string) map[string]bool {
	m := map[string]bool{}
	for _, c := range s {
		m[string(c)] = true
	}
	return m
}

func readInputDay3(path string) ([]rucksack, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	rucksacks := []rucksack{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		rucksacks = append(rucksacks, readRucksack(scanner.Text()))
	}

	return rucksacks, nil
}
