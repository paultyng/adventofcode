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

	dupes := map[K]int{}

	for _, m := range maps[1:] {
		for key := range m {
			if _, ok := maps[0][key]; ok {
				dupes[key]++
			}
		}
	}

	// only returns one duplicate key in case multiple were found
	for key, count := range dupes {
		if count == len(maps)-1 {
			return &key
		}
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
