package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

func runDay5Part1(ctx context.Context, args []string) (string, error) {
	path := "day5.input"
	if len(args) > 0 {
		path = args[0]
	}
	ship, moves, err := readInputDay5(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	for _, m := range moves {
		moving, remaining := popN(ship[m.From-1], m.Count)
		ship[m.From-1] = remaining
		// CrateMover 9000 does not reverse the order when moving
		ship[m.To-1] = push(ship[m.To-1], moving...)
	}

	return string(ship.TopCrates()), nil
}

func runDay5Part2(ctx context.Context, args []string) (string, error) {
	path := "day5.input"
	if len(args) > 0 {
		path = args[0]
	}
	ship, moves, err := readInputDay5(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	for _, m := range moves {
		moving, remaining := popN(ship[m.From-1], m.Count)
		ship[m.From-1] = remaining
		// CrateMover 9001 reverses the order
		reverse(moving)
		ship[m.To-1] = push(ship[m.To-1], moving...)
	}

	return string(ship.TopCrates()), nil
}

type ship []stack
type crate rune
type stack []crate
type move struct {
	Count int
	From  int
	To    int
}

func (s *ship) TopCrates() []crate {
	tops := []crate{}
	for _, st := range *s {
		tops = append(tops, st[len(st)-1])
	}
	return tops
}

func readShip(scanner *bufio.Scanner) (ship, error) {
	var stacks ship
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			return stacks, nil
		}

		// this is possibly the number line, confirm it matches our expectations
		// unsure how to handle greater than 1 digit? but doesn't matter
		if stacks != nil && strings.HasSuffix(" "+strings.TrimRight(line, " "), fmt.Sprintf(" %d", len(stacks))) {
			break
		}

		countStacks := (len(line) + 1) / 4

		if stacks == nil {
			stacks = make([]stack, countStacks)
		} else if len(stacks) != countStacks {
			panic("unexpected number of stacks")
		}

		for i := 0; i < countStacks; i++ {
			c := crate(line[i*4+1])
			if c == ' ' {
				continue
			}

			stacks[i] = append([]crate{c}, stacks[i]...)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan: %w", err)
	}

	return stacks, nil
}

func readMoves(scanner *bufio.Scanner) ([]move, error) {
	moves := []move{}
	for scanner.Scan() {
		line := scanner.Text()
		var m move
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &m.Count, &m.From, &m.To)
		if err != nil {
			return nil, fmt.Errorf("unable to parse move: %q %w", line, err)
		}

		moves = append(moves, m)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to read moves: %w", err)
	}

	return moves, nil
}

func readInputDay5(path string) (ship, []move, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	stacks, err := readShip(scanner)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read stacks: %w", err)
	}

	// read a blank line between sections
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("unable to scan: %w", err)
	}

	moves, err := readMoves(scanner)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read moves: %w", err)
	}

	return stacks, moves, nil
}
