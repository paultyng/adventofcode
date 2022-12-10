package main

import (
	"os"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestProcessCycles(t *testing.T) {
	r := strings.NewReader("noop\naddx 3\naddx -5\n")
	actual, err := processCycles(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []int{1, 1, 1, 4, 4, -1}
	if !slices.Equal(actual, expected) {
		t.Fatalf("expected does not match actual cycles:\nexp: %v\nact: %v", expected, actual)
	}

	f, err := os.Open("day10.input.test")
	if err != nil {
		t.Fatalf("unable to open input: %v", err)
	}
	defer f.Close()

	actual, err = processCycles(f)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected = []int{21, 19, 18, 21, 16, 18}
	actual = []int{actual[19], actual[59], actual[99], actual[139], actual[179], actual[219]}
	if !slices.Equal(actual, expected) {
		t.Fatalf("expected does not match actual cycles:\nexp: %v\nact: %v", expected, actual)
	}
}
