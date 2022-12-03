package main

import (
	"testing"

	"golang.org/x/exp/maps"
)

func TestReadRucksack(t *testing.T) {
	rucksack := readRucksack("vJrwpWtwJgWrhcsFMMfFFhFp")

	if !maps.Equal(rucksack.compartment1, map[string]bool{
		"v": true,
		"J": true,
		"r": true,
		"w": true,
		"p": true,
		"W": true,
		"t": true,
		"g": true,
	}) {
		t.Fatal("compartment 1 mismatch")
	}

	if !maps.Equal(rucksack.compartment2, map[string]bool{
		"h": true,
		"c": true,
		"s": true,
		"F": true,
		"M": true,
		"f": true,
		"p": true,
	}) {
		t.Fatal("compartment 1 mismatch")
	}
}

func TestItemPriority(t *testing.T) {
	for item, expectedPriority := range map[string]int{
		"a": 1,
		"z": 26,
		"A": 27,
		"Z": 52,
	} {
		t.Run(item, func(t *testing.T) {
			actual := itemPriority(item)
			if expectedPriority != actual {
				t.Errorf("Expected %d, got %d", expectedPriority, actual)
			}
		})
	}
}
