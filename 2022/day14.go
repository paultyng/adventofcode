package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func sandPileLine(p point, count int) line {
	return line{p.Move(0, -count-1), p}
}

type cave struct {
	Lines []line
	sand  map[int]map[int]int
}

func (c *cave) addToPile(p point) {
	if c.sand == nil {
		c.sand = map[int]map[int]int{}
	}
	yPiles, ok := c.sand[p.X]
	if !ok {
		yPiles = map[int]int{}
	}
	yPiles[p.Y]++
	c.sand[p.X] = yPiles
}

func (c *cave) isSandPile(p point) bool {
	yPiles, ok := c.sand[p.X]
	if !ok {
		return false
	}
	for y, count := range yPiles {
		if count == 0 {
			continue
		}
		pile := point{X: p.X, Y: y}
		if p == pile {
			return true
		}
		if count == 1 {
			continue
		}
		l := sandPileLine(pile, count)
		if l.IsOnLine(p) {
			return true
		}
	}

	return false
}

func (c *cave) bounds() (minX, maxX, minY, maxY int) {
	maxY = 0
	minX = 500
	maxX = 500
	for _, l := range c.Lines {
		for _, p := range l {
			if p.X < minX {
				minX = p.X
			}
			if p.X > maxX {
				maxX = p.X
			}
			if p.Y > maxY {
				maxY = p.Y
			}
		}
	}
	return minX, maxX, 0, maxY
}

func (c *cave) String() string {
	minX, maxX, minY, maxY := c.bounds()
	return c.string(minX, maxX, minY, maxY)
}

func (c *cave) string(minX, maxX, minY, maxY int) string {
	// minX, maxX, _, maxY := c.bounds()

	s := ""
	for y := minY; y <= maxY+1; y++ {
	XLoop:
		for x := minX - 1; x <= maxX+1; x++ {
			p := point{X: x, Y: y}
			if p.X == 500 && p.Y == 0 {
				s += "+"
				continue XLoop
			}
			for _, l := range c.Lines {
				if l.IsOnLine(p) {
					s += "#"
					continue XLoop
				}
			}
			if c.isSandPile(p) {
				s += "o"
				continue XLoop
			}
			s += "."
		}
		s += "\n"
	}
	return s
}

func (c *cave) isValidMove(p point) bool {
	for _, l := range c.Lines {
		if l.IsOnLine(p) {
			return false
		}
	}

	return !c.isSandPile(p)
}

func (c *cave) isFallingInAnEndlessVoid(grain point) bool {
	for _, l := range c.Lines {
		for _, p := range l {
			if p.Y > grain.Y {
				return false
			}
		}
	}

	return true
}

func (c *cave) PourGrain(entry point) (point, bool) {
	grain := entry
	for {
		if to := grain.Move(0, 1); c.isValidMove(to) {
			if c.isFallingInAnEndlessVoid(to) {
				return to, false
			}
			grain = to
			continue
		}
		if to := grain.Move(-1, 1); c.isValidMove(to) {
			grain = to
			continue
		}
		if to := grain.Move(1, 1); c.isValidMove(to) {
			grain = to
			continue
		}

		c.addToPile(grain)
		return grain, true
	}
}

func runDay14Part1(ctx context.Context, args []string) (string, error) {
	path := "day14.input"
	if len(args) > 0 {
		path = args[0]
	}
	c := &cave{}
	lines, err := readInputDay14(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}
	c.Lines = lines

	entryPoint := point{X: 500, Y: 0}
	grains := 0
	for {
		if _, ok := c.PourGrain(entryPoint); !ok {
			break
		}
		// if slices.Contains([]int{1, 2, 5, 22, 24}, grains-1) {
		// 	fmt.Println()
		// 	fmt.Println(c.String())
		// }
		grains++
	}
	// fmt.Println()
	// fmt.Println(c.String())

	return fmt.Sprintf("%d", grains), nil
}

func runDay14Part2(ctx context.Context, args []string) (string, error) {
	path := "day14.input"
	if len(args) > 0 {
		path = args[0]
	}
	_, err := readInputDay14(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	c := &cave{}
	lines, err := readInputDay14(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}
	c.Lines = lines

	minX, maxX, _, maxY := c.bounds()
	groundBuffer := 2*maxY + (maxX - minX)
	ground := line{point{minX - groundBuffer, maxY + 2}, point{maxX + groundBuffer, maxY + 2}}
	c.Lines = append(c.Lines, ground)

	entryPoint := point{X: 500, Y: 0}
	grains := 0
	for {
		grain, resting := c.PourGrain(entryPoint)
		if !resting {
			// fmt.Println()
			// fmt.Println(c.String())
			panic("endless void! ahhhhh!")
		}
		// if slices.Contains([]int{1, 2, (maxY * 2) + 1}, grains-1) || grains%1000 == 0 {
		// 	fmt.Println(grains, "grains")
		// 	fmt.Println(c.string(minX, maxX, 0, maxY+1))
		// }
		grains++
		if grain == entryPoint {
			break
		}
	}
	// fmt.Println()
	// fmt.Println(c.String())

	return fmt.Sprintf("%d", grains), nil
}

func readInputDay14(path string) ([]line, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	lines := []line{}
	err = readLines(input, func(_ int, text string) error {
		l := []point{}
		for _, coords := range strings.Split(text, " -> ") {
			p := point{}
			_, err := fmt.Sscanf(coords, "%d,%d", &p.X, &p.Y)
			if err != nil {
				return fmt.Errorf("unable to parse coordinates: %w", err)
			}
			l = append(l, p)
		}

		lines = append(lines, l)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	return lines, nil
}
