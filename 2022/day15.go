package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
)

type sensor struct {
	point
	ClosestBeacon point
}

func (s *sensor) bounds() (minX, maxX, minY, maxY int) {
	minX = s.X
	if s.ClosestBeacon.X < minX {
		minX = s.ClosestBeacon.X
	}
	maxX = s.X
	if s.ClosestBeacon.X > maxX {
		maxX = s.ClosestBeacon.X
	}
	minY = s.Y
	if s.ClosestBeacon.Y < minY {
		minY = s.ClosestBeacon.Y
	}
	maxY = s.Y
	if s.ClosestBeacon.Y > maxY {
		maxY = s.ClosestBeacon.Y
	}
	return minX, maxX, minY, maxY
}

type tunnel []sensor

func (t *tunnel) boundsWithCoverage() (minX, maxX, minY, maxY int) {
	if len(*t) == 0 {
		panic("no sensors")
	}

	minX, maxX, minY, maxY = (*t)[0].bounds()
	dist := (*t)[0].ManhattanDistance((*t)[0].ClosestBeacon)
	minX -= dist
	maxX += dist
	minY -= dist
	maxY += dist

	for _, s := range (*t)[1:] {
		xMinX, sMaxX, sMinY, sMaxY := s.bounds()
		dist := s.ManhattanDistance(s.ClosestBeacon)
		xMinX -= dist
		sMaxX += dist
		sMinY -= dist
		sMaxY += dist

		if xMinX < minX {
			minX = xMinX
		}
		if sMaxX > maxX {
			maxX = sMaxX
		}
		if sMinY < minY {
			minY = sMinY
		}
		if sMaxY > maxY {
			maxY = sMaxY
		}
	}
	return minX, maxX, minY, maxY
}

func (t *tunnel) TotalNoBeacon(y int) int {
	minX, maxX, _, _ := t.boundsWithCoverage()

	totalX := 0
NextX:
	for x := minX; x <= maxX; x++ {
		p := point{X: x, Y: y}

		for _, s := range *t {
			if p == s.point || p == s.ClosestBeacon {
				continue
			}

			dist := s.point.ManhattanDistance(s.ClosestBeacon)

			// todo: pre-filter sensor list as just points and distances?
			if s.Y-dist > y || s.Y+dist < y {
				continue
			}

			if s.ManhattanDistance(p) <= dist {
				totalX++
				continue NextX
			}
		}
	}

	return totalX
}

func runDay15Part1(ctx context.Context, args []string) (string, error) {
	if len(args) < 2 {
		panic("path and y required")
	}
	path := args[0]
	t, err := readInputDay15(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	y, err := strconv.Atoi(args[1])
	if err != nil {
		return "", fmt.Errorf("unable to parse y: %w", err)
	}
	noBeacon := t.TotalNoBeacon(y)

	return fmt.Sprintf("%d", noBeacon), nil
}

func runDay15Part2(ctx context.Context, args []string) (string, error) {
	path := "day15.input"
	if len(args) > 0 {
		path = args[0]
	}
	_, err := readInputDay15(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	panic("not implemented")
}

func readInputDay15(path string) (tunnel, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	sensors := []sensor{}
	err = readLines(input, func(_ int, line string) error {
		var s sensor
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.X, &s.Y, &s.ClosestBeacon.X, &s.ClosestBeacon.Y)
		if err != nil {
			return fmt.Errorf("unable to parse line: %w", err)
		}

		sensors = append(sensors, s)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	return sensors, nil
}
