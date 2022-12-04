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

KeyLoop:
	for key := range maps[0] {
		for _, m := range maps[1:] {
			_, ok := m[key]
			if !ok {
				continue KeyLoop
			}
		}
		return &key
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

func keyset[V comparable](v []V) map[V]struct{} {
	m := map[V]struct{}{}
	for _, i := range v {
		m[i] = struct{}{}
	}
	return m
}
