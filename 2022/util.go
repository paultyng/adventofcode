package main

import (
	"bufio"
	"fmt"
	"io"
)

func sumInts(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func readLines(r io.Reader, handleLine func(line string) error) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		err := handleLine(line)
		if err != nil {
			return fmt.Errorf("unable to handle line %q: %w", line, err)
		}
	}

	return nil
}
