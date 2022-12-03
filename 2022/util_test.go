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

func TestDuplicateKey(t *testing.T) {
	if actual := duplicateKey(map[string]bool{}, map[string]bool{}); actual != nil {
		t.Fatalf("expected nil")
	}

	if actual := duplicateKey(map[string]bool{"a": true, "b": true}, map[string]bool{"b": false}); *actual != "b" {
		t.Fatalf("expected b")
	}

	if actual := duplicateKey(
		map[string]bool{"a": true, "b": true, "c": true},
		map[string]bool{"a": false, "b": true},
		map[string]bool{"b": false},
	); *actual != "b" {
		t.Fatalf("expected b")
	}
}
