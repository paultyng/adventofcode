package main

import (
	"context"
	"testing"
)

func TestRunDay1Part1(t *testing.T) {
	ctx := context.TODO()

	answer, err := runDay1Part1(ctx, []string{"day1.input.test"})
	if err != nil {
		t.Fatal(err)
	}
	if answer != "24000" {
		t.Fatalf("Expected answer to be 24000, got %s", answer)
	}

	answer, err = runDay1Part1(ctx, []string{"day1.input"})
	if err != nil {
		t.Fatal(err)
	}
	if answer != "66719" {
		t.Fatalf("Expected answer to be 66719, got %s", answer)
	}
}

func TestRunDay1Part2(t *testing.T) {
	ctx := context.TODO()

	answer, err := runDay1Part2(ctx, []string{"day1.input.test"})
	if err != nil {
		t.Fatal(err)
	}
	if answer != "45000" {
		t.Fatalf("Expected answer to be 45000, got %s", answer)
	}

	answer, err = runDay1Part2(ctx, []string{"day1.input"})
	if err != nil {
		t.Fatal(err)
	}
	if answer != "198551" {
		t.Fatalf("Expected answer to be 198551, got %s", answer)
	}
}
