package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

const (
	startPacketLength  = 4
	startMessageLength = 14
)

func runDay6Part1(ctx context.Context, args []string) (string, error) {
	path := "day6.input"
	if len(args) > 0 {
		path = args[0]
	}
	signal, err := readInputDay6(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	if len(signal) <= startPacketLength {
		panic("invalid signal")
	}

	packetMarker := findUniqueSubstring(signal, 0, startPacketLength)
	if packetMarker < 0 {
		return "", fmt.Errorf("packet marker not found in %q", signal)
	}

	return fmt.Sprintf("%d", packetMarker+startPacketLength), nil
}

func runDay6Part2(ctx context.Context, args []string) (string, error) {
	path := "day6.input"
	if len(args) > 0 {
		path = args[0]
	}
	signal, err := readInputDay6(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	messageMarker := findUniqueSubstring(signal, 0, startMessageLength)
	if messageMarker < 0 {
		return "", fmt.Errorf("message marker not found in %q", signal)
	}

	return fmt.Sprintf("%d", messageMarker+startMessageLength), nil
}

func findUniqueSubstring(s string, startIndex, length int) int {
	if len(s)+startIndex < length {
		return -2
	}
	for i := startIndex; i < len(s)-length; i++ {
		unique := newSet([]rune(s[i : i+length])...)
		if len(unique) == length {
			return i
		}
	}
	return -1
}

func readInputDay6(path string) (string, error) {
	input, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("unable to scan: %w", err)
	}

	return scanner.Text(), nil
}
