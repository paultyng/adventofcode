package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func runDay11Part1(ctx context.Context, args []string) (string, error) {
	path := "day11.input"
	if len(args) > 0 {
		path = args[0]
	}
	monkies, err := readInputDay11(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	// fmt.Printf("Monkies: %v\n", monkies)

	inspections := make([]int, len(monkies))

	for round := 1; round <= 20; round++ {
		// items := [][]int{}
		// for i := range monkies {
		// 	items = append(items, monkies[i].StartingItems)
		// 	monkies[i].StartingItems = []int{}
		// }

		for monkeyIndex, m := range monkies {
			originalItems := m.StartingItems
			monkies[monkeyIndex].StartingItems = []int{}

			// fmt.Printf("Monkey %d:\n", monkeyIndex)

			for _, worry := range originalItems {
				inspections[monkeyIndex]++
				// fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", worry)

				worry = m.Operation.Apply(worry)
				// fmt.Printf("    Worry level after operation is now %d.\n", worry)

				worry = int(worry / 3)
				// fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", worry)

				if worry%m.Test.DivisibleBy == 0 {
					// fmt.Printf("    Current worry level is divisible by %d.\n", m.Test.DivisibleBy)
					// fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", worry, m.Test.TrueMonkey)
					monkies[m.Test.TrueMonkey].StartingItems = append(monkies[m.Test.TrueMonkey].StartingItems, worry)
				} else {
					// fmt.Printf("    Current worry level is not divisible by %d.\n", m.Test.DivisibleBy)
					// fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", worry, m.Test.FalseMonkey)
					monkies[m.Test.FalseMonkey].StartingItems = append(monkies[m.Test.FalseMonkey].StartingItems, worry)
				}
			}
		}

		fmt.Println()

		fmt.Printf("\nRound %d\n", round)
		for i, m := range monkies {
			fmt.Printf("  Monkey %d: %v\n", i, m.StartingItems)
		}
	}

	slices.Sort(inspections)

	return fmt.Sprintf("%d", inspections[len(inspections)-1]*inspections[len(inspections)-2]), nil
}

func runDay11Part2(ctx context.Context, args []string) (string, error) {
	path := "day11.input"
	if len(args) > 0 {
		path = args[0]
	}
	_, err := readInputDay11(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	panic("not implemented")
}

type monkey struct {
	StartingItems []int
	Operation     worryExpression
	Test          worryTest
}

type worryExpression string

func (e *worryExpression) Apply(old int) int {
	fields := strings.Fields(string(*e))
	if len(fields) != 2 {
		panic("invalid expression")
	}

	left := old
	var right int

	switch {
	case fields[1] == "old":
		right = old
	default:
		v, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(fmt.Sprintf("invalid operand %q: %s", fields[1], err))
		}
		right = v
	}

	switch fields[0] {
	case "+":
		return left + right
	case "*":
		return left * right
	default:
		panic(fmt.Sprintf("unexpected operator %s", fields[0]))
	}
}

type worryTest struct {
	DivisibleBy int
	TrueMonkey  int
	FalseMonkey int
}

func stringAfter(s, prefix string) string {
	if !strings.HasPrefix(s, prefix) {
		panic(fmt.Sprintf("expected %q to have prefix %q", s, prefix))
	}
	return s[len(prefix):]
}
func readInputDay11(path string) ([]monkey, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	monkies := []monkey{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		var m monkey
		var monkeyIndex int
		_, err := fmt.Sscanf(scanner.Text(), "Monkey %d", &monkeyIndex)
		if err != nil {
			return nil, fmt.Errorf("unable to parse monkey index: %w", err)
		}

		if monkeyIndex != len(monkies) {
			panic(fmt.Sprintf("unexpected monkey index %d", monkeyIndex))
		}

		if !scanner.Scan() {
			return nil, fmt.Errorf("unable to move to starting items")
		}
		startingItems := stringAfter(scanner.Text(), "  Starting items: ")
		for _, s := range strings.Split(startingItems, ",") {
			s = strings.TrimSpace(s)
			i, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("unable to parse starting item %q: %w", s, err)
			}
			m.StartingItems = append(m.StartingItems, i)
		}

		if !scanner.Scan() {
			return nil, fmt.Errorf("unable to move to operation")
		}
		op := stringAfter(scanner.Text(), "  Operation: new = old")
		op = strings.TrimSpace(op)
		m.Operation = worryExpression(op)

		if !scanner.Scan() {
			return nil, fmt.Errorf("unable to move to test")
		}
		var divisibleBy int
		_, err = fmt.Sscanf(scanner.Text(), "  Test: divisible by %d", &divisibleBy)
		if err != nil {
			return nil, fmt.Errorf("unable to parse divisible by: %w", err)
		}
		m.Test.DivisibleBy = divisibleBy

		if !scanner.Scan() {
			return nil, fmt.Errorf("unable to move to if true")
		}
		var trueMonkey int
		_, err = fmt.Sscanf(scanner.Text(), "    If true: throw to monkey %d", &trueMonkey)
		if err != nil {
			return nil, fmt.Errorf("unable to parse divisible by: %w", err)
		}
		m.Test.TrueMonkey = trueMonkey

		if !scanner.Scan() {
			return nil, fmt.Errorf("unable to move if false")
		}
		var falseMonkey int
		_, err = fmt.Sscanf(scanner.Text(), "    If false: throw to monkey %d", &falseMonkey)
		if err != nil {
			return nil, fmt.Errorf("unable to parse divisible by: %w", err)
		}
		m.Test.FalseMonkey = falseMonkey

		monkies = append(monkies, m)
	}

	return monkies, nil
}
