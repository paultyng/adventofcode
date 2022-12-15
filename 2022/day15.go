package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

type sensor struct {
	point
	ClosestBeacon point

	dist *int
}

func (s *sensor) ManhattanDistance() int {
	if s.dist == nil {
		dist := s.point.ManhattanDistance(s.ClosestBeacon)
		s.dist = &dist
	}

	return *s.dist
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
	dist := (*t)[0].ManhattanDistance()
	minX -= dist
	maxX += dist
	minY -= dist
	maxY += dist

	for _, s := range (*t)[1:] {
		xMinX, sMaxX, sMinY, sMaxY := s.bounds()
		dist := s.ManhattanDistance()
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

			dist := s.ManhattanDistance()

			// todo: pre-filter sensor list as just points and distances?
			if s.Y-dist > y || s.Y+dist < y {
				continue
			}

			if s.point.ManhattanDistance(p) <= dist {
				totalX++
				continue NextX
			}
		}
	}

	return totalX
}

func (t *tunnel) FirstPossibleBeacon(maxSearch int) point {
	minX, maxX, minY, maxY := t.boundsWithCoverage()
	if minX < 0 {
		minX = 0
	}
	if maxX > maxSearch {
		maxX = maxSearch
	}
	if minY < 0 {
		minY = 0
	}
	if maxY > maxSearch {
		maxY = maxSearch
	}

	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan point, 1)
	for x := minX; x <= maxX; x++ {
		go func(ctx context.Context, x int) {
		NextY:
			for y := minY; y <= maxY; y++ {
				p := point{X: x, Y: y}
				for _, s := range *t {
					select {
					case <-ctx.Done():
						return
					default:
						sensorDist := s.ManhattanDistance()
						pointDist := s.point.ManhattanDistance(p)
						if p == s.point || p == s.ClosestBeacon || pointDist <= sensorDist {
							if y < s.Y {
								y = s.Y + (sensorDist - pointDist)
							} else {
								y += sensorDist - pointDist
							}
							continue NextY
						}
					}
				}
				result <- p
			}
		}(ctx, x)
	}

	// this may never finish :'(

	found := <-result
	cancel()
	return found
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
	if len(args) < 2 {
		panic("path and max required")
	}
	path := args[0]
	t, err := readInputDay15(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	max, err := strconv.Atoi(args[1])
	if err != nil {
		return "", fmt.Errorf("unable to parse max: %w", err)
	}

	signal := t.FirstPossibleBeacon(max)

	answer := big.NewInt(int64(signal.X))
	answer = answer.Mul(answer, big.NewInt(4000000))
	answer = answer.Add(answer, big.NewInt(int64(signal.Y)))

	return fmt.Sprintf("%d", answer), nil
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
