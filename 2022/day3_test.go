package main

import (
	"testing"

	"golang.org/x/exp/maps"
)

func TestReadRucksack(t *testing.T) {
	rucksack := readRucksack("vJrwpWtwJgWrhcsFMMfFFhFp")

	if !maps.Equal(rucksack.compartment1, newSet('v', 'J', 'r', 'w', 'p', 'W', 't', 'g')) {
		t.Fatal("compartment 1 mismatch")
	}

	if !maps.Equal(rucksack.compartment2, newSet('h', 'c', 's', 'F', 'M', 'f', 'p')) {
		t.Fatal("compartment 1 mismatch")
	}
}

func TestItemPriority(t *testing.T) {
	for item, expectedPriority := range map[rune]int{
		'a': 1,
		'z': 26,
		'A': 27,
		'Z': 52,
	} {
		t.Run(string(item), func(t *testing.T) {
			actual := itemPriority(item)
			if expectedPriority != actual {
				t.Errorf("Expected %d, got %d", expectedPriority, actual)
			}
		})
	}
}
