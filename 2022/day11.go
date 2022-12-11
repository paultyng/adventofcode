package main

import (
	"bufio"
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func totalMonkeyBusiness(monkies []monkey, totalRounds int, worryBackoff int) int64 {
	inspections := make([]int64, len(monkies))

	var reducer big.Int
	for _, m := range monkies {
		if reducer.Int64() == 0 {
			reducer = *big.NewInt(int64(m.Test.DivisibleBy))
			continue
		}

		var v big.Int
		v.Mul(&reducer, big.NewInt(int64(m.Test.DivisibleBy)))
		reducer = v
	}
	fmt.Printf("Reducer: %d\n", reducer.Int64())

	for round := 1; round <= totalRounds; round++ {
		// items := [][]int{}
		// for i := range monkies {
		// 	items = append(items, monkies[i].StartingItems)
		// 	monkies[i].StartingItems = []int{}
		// }

		for monkeyIndex, m := range monkies {
			originalItems := m.StartingItems
			monkies[monkeyIndex].StartingItems = []big.Int{}

			// fmt.Printf("Monkey %d:\n", monkeyIndex)

			for _, worry := range originalItems {
				inspections[monkeyIndex]++
				// fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", worry)

				worry = m.Operation.Apply(worry)
				// fmt.Printf("    Worry level after operation is now %d.\n", worry)

				if worryBackoff > 0 {
					var v big.Int
					v.Div(&worry, big.NewInt(int64(worryBackoff)))
					worry = v
					// fmt.Printf("    Monkey gets bored with item. Worry level is divided by %d to %d.\n", worryBackoff, worry)
				}

				if reducer.Int64() != 0 {
					var v big.Int
					v.Mod(&worry, &reducer)
					worry = v
				}

				var mod big.Int
				if mod.Mod(&worry, big.NewInt(int64(m.Test.DivisibleBy))); mod.Int64() == 0 {
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

		// fmt.Println()

		// fmt.Printf("\nAfter round %d, the monkeys are holding items with these worry levels:\n", round)
		// for i, m := range monkies {
		// 	fmt.Printf("Monkey %d: %v\n", i, m.StartingItems)
		// }

		// if slices.Contains([]int{1, 20, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000}, round) {
		// 	fmt.Printf("\n== After round %d ==\n", round)
		// 	for i, m := range inspections {
		// 		fmt.Printf("Monkey %d inspected items %d times.\n", i, m)
		// 	}
		// }
	}

	slices.Sort(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

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

	monkeyBusiness := totalMonkeyBusiness(monkies, 20, 3)
	return fmt.Sprintf("%d", monkeyBusiness), nil
}

func runDay11Part2(ctx context.Context, args []string) (string, error) {
	path := "day11.input"
	if len(args) > 0 {
		path = args[0]
	}
	monkies, err := readInputDay11(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	monkeyBusiness := totalMonkeyBusiness(monkies, 10000, 0)
	return fmt.Sprintf("%d", monkeyBusiness), nil
}

type monkey struct {
	StartingItems []big.Int
	Operation     worryExpression
	Test          worryTest
}

type worryExpression string

func (e *worryExpression) Apply(old big.Int) big.Int {
	fields := strings.Fields(string(*e))
	if len(fields) != 2 {
		panic("invalid expression")
	}

	left := old
	var right big.Int

	switch {
	case fields[1] == "old":
		right = old
	default:
		v, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(fmt.Sprintf("invalid operand %q: %s", fields[1], err))
		}
		right.SetInt64(int64(v))
	}

	switch fields[0] {
	case "+":
		var v big.Int
		v.Add(&left, &right)
		return v
	case "*":
		var v big.Int
		v.Mul(&left, &right)
		return v
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
			var v big.Int
			if _, ok := v.SetString(s, 10); !ok {
				return nil, fmt.Errorf("unable to parse starting item %q", s)
			}
			m.StartingItems = append(m.StartingItems, v)
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
