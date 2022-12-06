// Code generated by "go generate" - DO NOT EDIT.

package main

import "fmt"

func runPartFactory(currentDay int, part string) runPart {
	switch fmt.Sprintf("Day%dPart%s", currentDay, part) {
	default:
		panic(fmt.Sprintf("day %d part %s not implemented", currentDay, part))
	case "Day1Part1":
		return runDay1Part1
	case "Day1Part2":
		return runDay1Part2
	case "Day2Part1":
		return runDay2Part1
	case "Day2Part2":
		return runDay2Part2
	case "Day3Part1":
		return runDay3Part1
	case "Day3Part2":
		return runDay3Part2
	case "Day4Part1":
		return runDay4Part1
	case "Day4Part2":
		return runDay4Part2
	case "Day5Part1":
		return runDay5Part1
	case "Day5Part2":
		return runDay5Part2
	case "Day6Part1":
		return runDay6Part1
	case "Day6Part2":
		return runDay6Part2
	}
}
