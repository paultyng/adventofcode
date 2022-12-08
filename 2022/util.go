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

type set[V comparable] map[V]struct{}

func newSet[V comparable](values ...V) set[V] {
	m := set[V]{}
	m.add(values...)
	return m
}

func (s set[V]) add(values ...V) {
	for _, v := range values {
		s[v] = struct{}{}
	}
}

// push isn't really needed as its the same as append, but included for consistent naming with pop.
func push[T any](s []T, v ...T) []T {
	return append(s, v...)
}

func pop[T any](s []T) (T, []T) {
	return s[len(s)-1], s[:len(s)-1]
}

func popN[T any](s []T, n int) ([]T, []T) {
	popped := s[len(s)-n:]
	// reverse for FILO
	reverse(popped)
	return popped, s[:len(s)-n]
}

// reverse reverses the order of the elements in the slice in place.
func reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func filter[T any](s []T, test func(T) bool) (ret []T) {
	for _, v := range s {
		if test(v) {
			ret = append(ret, v)
		}
	}
	return
}

func all[T any](s []T, test func(T) bool) bool {
	return len(s) == len(filter(s, test))
}

// func any[T any](s []T, test func(T) bool) bool {
// 	for _, v := range s {
// 		if test(v) {
// 			return true
// 		}
// 	}
// 	return false
// }
