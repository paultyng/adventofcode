package main

import "testing"

func TestSumInts(t *testing.T) {
	if actual := sum([]int{1, 2, 3, 4, 5}); actual != 15 {
		t.Errorf("Expected %d, got %d", 15, actual)
	}

	if actual := sum([]int{}); actual != 0 {
		t.Errorf("Expected %d, got %d", 0, actual)
	}

	if actual := sum(nil); actual != 0 {
		t.Errorf("Expected %d, got %d", 0, actual)
	}
}
