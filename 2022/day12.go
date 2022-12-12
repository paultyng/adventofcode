package main

import (
	"context"
	"fmt"
	"os"

	"github.com/RyanCarrier/dijkstra"
)

func runDay12Part1(ctx context.Context, args []string) (string, error) {
	path := "day12.input"
	if len(args) > 0 {
		path = args[0]
	}
	hm, err := readInputDay12(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	shortest := hm.ShortestPath([]point{hm.Start})
	return fmt.Sprintf("%d", len(shortest)-1), nil
}

func runDay12Part2(ctx context.Context, args []string) (string, error) {
	path := "day12.input"
	if len(args) > 0 {
		path = args[0]
	}
	hm, err := readInputDay12(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	starts := []point{}
	for y := range hm.G {
		for x := range hm.G[y] {
			if hm.G.At(point{x, y}) != 0 {
				continue
			}

			starts = append(starts, point{x, y})
		}
	}

	shortest := hm.ShortestPath(starts)
	return fmt.Sprintf("%d", len(shortest)-1), nil
}

type heightmap struct {
	G     grid[int]
	Start point
	End   point
}

func (hm *heightmap) validMoves(from point) []point {
	moves := []point{}

	for _, to := range []point{
		{from.X, from.Y - 1}, // up
		{from.X, from.Y + 1}, // down
		{from.X - 1, from.Y}, // left
		{from.X + 1, from.Y}, // right
	} {
		if !hm.G.InBounds(to) {
			continue
		}

		fromHeight := hm.G.At(from)
		toHeight := hm.G.At(to)
		dh := toHeight - fromHeight

		if dh > 1 {
			continue
		}

		moves = append(moves, to)
	}

	return moves
}

const pointUniqueIDOffset = 100000

func (p point) uniqueID() int {
	return p.Y*pointUniqueIDOffset + p.X
}

func fromUniqueID(id int) point {
	return point{id % pointUniqueIDOffset, id / pointUniqueIDOffset}
}

func (hm *heightmap) ShortestPath(starts []point) []point {
	graph := dijkstra.NewGraph()

	for y := range hm.G {
		for x := range hm.G[y] {
			from := point{x, y}
			graph.AddVertex(from.uniqueID())
		}
	}

	for y := range hm.G {
		for x := range hm.G[y] {
			from := point{x, y}
			for _, to := range hm.validMoves(from) {
				graph.AddArc(from.uniqueID(), to.uniqueID(), 1)
			}
		}
	}

	var shortestPath []int
	skip := newSet[point]()
	for _, start := range starts {
		if _, ok := skip[start]; ok {
			continue
		}

		best, err := graph.Shortest(start.uniqueID(), hm.End.uniqueID())
		if err != nil {
			continue
		}

		// bests, err := graph.ShortestAll(start.uniqueID(), hm.End.uniqueID())
		// if err != nil {
		// 	continue
		// }
		// for _, best := range bests {
		if shortestPath == nil || len(best.Path) < len(shortestPath) {
			shortestPath = best.Path
		}

		// see if we can skip any future starts
		foundShorter := false
		for i := len(best.Path) - 1; i > 1; i-- {
			p := fromUniqueID(best.Path[i])
			if hm.G.At(p) == 0 {
				if foundShorter {
					skip.Add(p)
				} else {
					foundShorter = true
				}
			}
		}
		// }
	}

	shortest := []point{}
	for _, id := range shortestPath {
		shortest = append(shortest, fromUniqueID(id))
	}

	return shortest
}

func readInputDay12(path string) (*heightmap, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	hm := &heightmap{}
	err = readLines(input, func(y int, line string) error {
		row := []int{}
		for x, sq := range line {
			switch sq {
			case 'S':
				hm.Start = point{x, y}
				sq = 'a'
			case 'E':
				hm.End = point{x, y}
				sq = 'z'
			}

			height := int(sq) - int('a')
			row = append(row, height)
		}
		hm.G = append(hm.G, row)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	// fmt.Println(hm.G.String())

	return hm, nil
}
