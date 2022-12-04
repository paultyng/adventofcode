package main

import (
	"testing"

	"golang.org/x/exp/maps"
)

func TestReadRucksack(t *testing.T) {
	rucksack := readRucksack("vJrwpWtwJgWrhcsFMMfFFhFp")

	var empty struct{}

	if !maps.Equal(rucksack.compartment1, map[string]struct{}{
		"v": empty,
		"J": empty,
		"r": empty,
		"w": empty,
		"p": empty,
		"W": empty,
		"t": empty,
		"g": empty,
	}) {
		t.Fatal("compartment 1 mismatch")
	}

	if !maps.Equal(rucksack.compartment2, map[string]struct{}{
		"h": empty,
		"c": empty,
		"s": empty,
		"F": empty,
		"M": empty,
		"f": empty,
		"p": empty,
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
