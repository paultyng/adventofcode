package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
)

func runDay8Part1(ctx context.Context, args []string) (string, error) {
	path := "day8.input"
	if len(args) > 0 {
		path = args[0]
	}
	f, err := readInputDay8(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	return fmt.Sprintf("%d", f.VisibleTrees()), nil
}

func runDay8Part2(ctx context.Context, args []string) (string, error) {
	path := "day8.input"
	if len(args) > 0 {
		path = args[0]
	}
	f, err := readInputDay8(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	return fmt.Sprintf("%d", f.MaxScenicScore()), nil
}

type forest [][]int

func (f forest) MaxScenicScore() int {
	max := 0
	for x := 0; x < len(f); x++ {
		for y := 0; y < len(f[0]); y++ {
			if score := f.scenicScore(x, y); score > max {
				max = score
			}
		}
	}
	return max
}

func (f forest) scenicScore(treeX, treeY int) int {
	treeHeight := f[treeX][treeY]

	up := 0
	for x := treeX - 1; x >= 0; x-- {
		up++
		if f[x][treeY] >= treeHeight {
			break
		}
	}

	down := 0
	for x := treeX + 1; x < len(f); x++ {
		down++
		if f[x][treeY] >= treeHeight {
			break
		}
	}

	left := 0
	for y := treeY - 1; y >= 0; y-- {
		left++
		if f[treeX][y] >= treeHeight {
			break
		}
	}

	right := 0
	for y := treeY + 1; y < len(f[treeX]); y++ {
		right++
		if f[treeX][y] >= treeHeight {
			break
		}
	}

	// fmt.Printf("up: %d, down: %d, left: %d, right: %d\n", up, down, left, right)

	return up * down * left * right
}

func (f forest) VisibleTrees() int {
	visible := 0
	if len(f) == 0 || len(f[0]) == 0 {
		return 0
	}
	if len(f) <= 2 || len(f[0]) <= 2 {
		return len(f) * len(f[0])
	}

	// add edges
	visible += len(f)*2 + len(f[0])*2 - 4

	// iterate interior
	for x := 1; x < len(f)-1; x++ {
		for y := 1; y < len(f[0])-1; y++ {
			if !f.isVisible(x, y) {
				continue
			}
			visible++
		}
	}

	return visible
}

func (f forest) isVisible(treeX, treeY int) bool {
	treeHeight := f[treeX][treeY]

	visible := true
	for x := 0; x < treeX; x++ {
		if f[x][treeY] >= treeHeight {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for x := treeX + 1; x < len(f); x++ {
		if f[x][treeY] >= treeHeight {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for y := 0; y < treeY; y++ {
		if f[treeX][y] >= treeHeight {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for y := treeY + 1; y < len(f[treeX]); y++ {
		if f[treeX][y] >= treeHeight {
			visible = false
			break
		}
	}
	return visible
}

func readInputDay8(path string) (forest, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	f := forest{}
	err = readLines(input, func(_ int, line string) error {
		if len(f) > 0 && len(line) != len(f[0]) {
			return fmt.Errorf("unexpected line length: %q", line)
		}

		row := []int{}

		for _, tree := range line {
			treeHeight, err := strconv.Atoi(string(tree))
			if err != nil {
				return fmt.Errorf("unable to parse row: %q %w", line, err)
			}
			row = append(row, treeHeight)
		}

		f = append(f, row)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	return f, nil
}
