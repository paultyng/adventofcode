package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

func runDay13Part1(ctx context.Context, args []string) (string, error) {
	path := "day13.input"
	if len(args) > 0 {
		path = args[0]
	}
	pairs, err := readInputDay13(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	sum := 0
	for i, pair := range pairs {
		if !pair.CorrectOrder() {
			continue
		}
		// fmt.Printf("Pair %d is correct order\n%v\n%v\n\n", i+1, pair.Left, pair.Right)
		sum += i + 1
	}

	return fmt.Sprintf("%d", sum), nil
}

func runDay13Part2(ctx context.Context, args []string) (string, error) {
	path := "day13.input"
	if len(args) > 0 {
		path = args[0]
	}
	pairs, err := readInputDay13(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	divider1Packet := packet{packetData{List: []packetData{{Int: 2}}}}
	divider2Packet := packet{packetData{List: []packetData{{Int: 6}}}}

	packets := []packet{
		divider1Packet,
		divider2Packet,
	}
	for _, p := range pairs {
		packets = append(packets, p.Left, p.Right)
	}

	slices.SortFunc(packets, func(left, right packet) bool {
		return left.Compare(right) < 0
	})

	divider1 := 0
	divider2 := 0

	for i, p := range packets {
		switch {
		case p.Compare(divider1Packet) == 0:
			divider1 = i + 1
		case p.Compare(divider2Packet) == 0:
			divider2 = i + 1
		}

	}

	return fmt.Sprintf("%d", divider1*divider2), nil
}

type packetData struct {
	List []packetData `json:"-"`
	Int  int          `json:"-"`
}

func (pd *packetData) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if data[0] == '[' {
		var list []packetData
		err := json.Unmarshal(data, &list)
		if err != nil {
			return fmt.Errorf("unable to unmarshal list: %w", err)
		}
		*pd = packetData{List: list}
	} else {
		var i int
		err := json.Unmarshal(data, &i)
		if err != nil {
			return fmt.Errorf("unable to unmarshal int: %w", err)
		}
		*pd = packetData{Int: i}
	}
	return nil
}

func (left packetData) Compare(right packetData) int {
	if left.List != nil || right.List != nil {
		// list comparison
		leftList := left.List
		if left.List == nil {
			leftList = []packetData{{Int: left.Int}}
		}

		rightList := right.List
		if right.List == nil {
			rightList = []packetData{{Int: right.Int}}
		}

		return packet(leftList).Compare(rightList)
	} else {
		// int comparison
		switch {
		case left.Int < right.Int:
			return -1
		case left.Int == right.Int:
			return 0
		case left.Int > right.Int:
			return 1
		}
	}

	panic("unexpected comparison")
}

type packet []packetData

func (left packet) Compare(right packet) int {
	for i := 0; i < len(left); i++ {
		if i >= len(right) {
			return 1
		}

		pdc := left[i].Compare(right[i])
		if pdc != 0 {
			return pdc
		}
	}

	if len(right) > len(left) {
		return -1
	}

	return 0
}

type packetPair struct {
	Left  packet
	Right packet
}

func (pp packetPair) CorrectOrder() bool {
	return pp.Left.Compare(pp.Right) <= 0
}

func readInputDay13(path string) ([]packetPair, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	pairs := []packetPair{}
	var currentPair *packetPair
	err = readLines(input, func(_ int, line string) error {
		if line == "" {
			if currentPair != nil {
				pairs = append(pairs, *currentPair)
				currentPair = nil
			}
			return nil
		}

		var target *packet
		if currentPair == nil {
			currentPair = &packetPair{}
			target = &currentPair.Left
		} else {
			target = &currentPair.Right
		}
		err := json.Unmarshal([]byte(line), &target)
		if err != nil {
			return fmt.Errorf("unable to unmarshal: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	if currentPair != nil {
		pairs = append(pairs, *currentPair)
		currentPair = nil
	}

	return pairs, nil
}
