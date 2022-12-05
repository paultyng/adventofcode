package main

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestPastAnswers(t *testing.T) {
	for _, part := range []struct {
		run      runPart
		args     []string
		expected string
	}{
		{runDay1Part1, []string{"day1.input.test"}, "24000"},
		{runDay1Part1, []string{"day1.input"}, "66719"},
		{runDay1Part2, []string{"day1.input.test"}, "45000"},
		{runDay1Part2, []string{"day1.input"}, "198551"},

		{runDay2Part1, []string{"day2.input.test"}, "15"},
		{runDay2Part1, []string{"day2.input"}, "11150"},
		{runDay2Part2, []string{"day2.input.test"}, "12"},
		{runDay2Part2, []string{"day2.input"}, "8295"},

		{runDay3Part1, []string{"day3.input.test"}, "157"},
		{runDay3Part1, []string{"day3.input"}, "7850"},
		{runDay3Part2, []string{"day3.input.test"}, "70"},
		{runDay3Part2, []string{"day3.input"}, "2581"},

		{runDay4Part1, []string{"day4.input.test"}, "2"},
		{runDay4Part1, []string{"day4.input"}, "305"},
		{runDay4Part2, []string{"day4.input.test"}, "4"},
		{runDay4Part2, []string{"day4.input"}, "811"},

		{runDay5Part1, []string{"day5.input.test"}, "CMZ"},
		{runDay5Part1, []string{"day5.input"}, "QPJPLMNNR"},
		{runDay5Part2, []string{"day5.input.test"}, "MCD"},
		{runDay5Part2, []string{"day5.input"}, "BQDNWJPVJ"},
	} {
		t.Run(fmt.Sprintf("%s %#v", runtime.FuncForPC(reflect.ValueOf(part.run).Pointer()).Name(), part.args), func(t *testing.T) {
			ctx := context.TODO()
			answer, err := part.run(ctx, part.args)
			if err != nil {
				t.Fatal(err)
			}
			if answer != part.expected {
				t.Fatalf("Expected answer to be %s, got %s", part.expected, answer)
			}
		})
	}
}
