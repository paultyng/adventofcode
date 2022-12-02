package main

import (
	"context"
	"testing"
)

func TestRunDay2Part1(t *testing.T) {
	ctx := context.TODO()

	answer, err := runDay2Part1(ctx, []string{"day2.input.test"})
	if err != nil {
		t.Fatal(err)
	}
	if answer != "15" {
		t.Fatalf("Expected answer to be 15, got %s", answer)
	}

	answer, err = runDay2Part1(ctx, []string{"day2.input"})
	if err != nil {
		t.Fatal(err)
	}
	if answer != "11150" {
		t.Fatalf("Expected answer to be 11150, got %s", answer)
	}
}

func TestRunDay2Part2(t *testing.T) {
	ctx := context.TODO()

	answer, err := runDay2Part2(ctx, []string{"day2.input.test"})
	if err != nil {
		t.Fatal(err)
	}
	if answer != "12" {
		t.Fatalf("Expected answer to be 12, got %s", answer)
	}

	answer, err = runDay2Part2(ctx, []string{"day2.input"})
	if err != nil {
		t.Fatal(err)
	}
	if answer != "8295" {
		t.Fatalf("Expected answer to be 8295, got %s", answer)
	}
}
