package main

import (
	"bufio"
	"fmt"
	"io"
)

func sum[V int](a []V) V {
	var sum V
	for _, v := range a {
		sum += v
	}
	return sum
}

func duplicateKey[M ~map[K]V, K comparable, V any](maps ...M) *K {
	if len(maps) <= 1 {
		return nil
	}

	for key := range maps[0] {
		for _, m := range maps[1:] {
			_, ok := m[key]
			if !ok {
				goto nextKey
			}
		}
		return &key
	nextKey:
	}

	return nil
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
