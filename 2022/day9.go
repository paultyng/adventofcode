package main

import (
	"context"
	"fmt"
	"os"
)

func runDay9Part1(ctx context.Context, args []string) (string, error) {
	path := "day9.input"
	if len(args) > 0 {
		path = args[0]
	}
	motions, err := readInputDay9(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	r := newRope(2)
	knotPoints := r.Move(motions)
	uniqueTail := newSet(knotPoints[len(r.Knots)-1]...)
	return fmt.Sprintf("%d", len(uniqueTail)), nil
}

func runDay9Part2(ctx context.Context, args []string) (string, error) {
	path := "day9.input"
	if len(args) > 0 {
		path = args[0]
	}
	motions, err := readInputDay9(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	r := newRope(10)
	knotPoints := r.Move(motions)
	uniqueTail := newSet(knotPoints[len(r.Knots)-1]...)
	return fmt.Sprintf("%d", len(uniqueTail)), nil
}

type motion struct {
	Dir   direction
	Steps int
}

func (a point) TouchesDay9(b point) bool {
	if !(a.X-1 <= b.X && b.X <= a.X+1) {
		return false
	}

	if !(a.Y-1 <= b.Y && b.Y <= a.Y+1) {
		return false
	}

	return true
}

func (a *point) MoveTowardsDay9(b point) {
	if a.TouchesDay9(b) {
		panic("cannot move towards a point that touches the current point")
	}

	dx := b.X - a.X
	dy := b.Y - a.Y

	switch {
	case dx > 0:
		a.X++
	case dx < 0:
		a.X--
	}

	switch {
	case dy > 0:
		a.Y++
	case dy < 0:
		a.Y--
	}
}

type direction string

const (
	up    direction = "U"
	down  direction = "D"
	left  direction = "L"
	right direction = "R"
)

func (p *point) move(dir direction) {
	switch dir {
	case up:
		p.Y++
	case down:
		p.Y--
	case left:
		p.X--
	case right:
		p.X++
	}
}

type rope struct {
	Knots []point
}

func newRope(knots int) *rope {
	return &rope{
		Knots: make([]point, knots),
	}
}

func (r *rope) Move(motions []motion) [][]point {
	if len(r.Knots) < 2 {
		panic("not enough knots to simulate")
	}

	knotPoints := make([][]point, len(r.Knots))
	for i, k := range r.Knots {
		knotPoints[i] = []point{k}
	}
	for _, m := range motions {
		motionKnotPoints := r.moveMotion(m)
		for i, k := range motionKnotPoints {
			knotPoints[i] = append(knotPoints[i], k...)
		}
	}
	return knotPoints
}

func (r *rope) moveMotion(m motion) [][]point {
	knotPoints := make([][]point, len(r.Knots))

	for step := 0; step < m.Steps; step++ {
		// start with the head, then loop over following knots
		lastKnot := &r.Knots[0]
		lastKnot.move(m.Dir)
		knotPoints[0] = append(knotPoints[0], *lastKnot)

		for i := 1; i < len(r.Knots); i++ {
			k := &r.Knots[i]

			if !k.TouchesDay9(*lastKnot) {
				k.MoveTowardsDay9(*lastKnot)
			}

			knotPoints[i] = append(knotPoints[i], *k)

			lastKnot = k
		}
	}

	return knotPoints
}

func readInputDay9(path string) ([]motion, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	motions := []motion{}
	err = readLines(input, func(_ int, line string) error {
		var m motion
		_, err := fmt.Sscan(line, &m.Dir, &m.Steps)
		if err != nil {
			return fmt.Errorf("unable to parse line: %w", err)
		}

		motions = append(motions, m)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	return motions, nil
}
