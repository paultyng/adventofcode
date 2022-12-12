package main

import (
	"bufio"
	"fmt"
	"io"

	"golang.org/x/exp/maps"
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

func readLines(r io.Reader, handleLine func(i int, line string) error) error {
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		err := handleLine(i, line)
		if err != nil {
			return fmt.Errorf("unable to handle line %q: %w", line, err)
		}

		i++
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}
	return nil
}

type set[V comparable] map[V]struct{}

func newSet[V comparable](values ...V) set[V] {
	m := set[V]{}
	m.Add(values...)
	return m
}

func (s set[V]) Keys() []V {
	return maps.Keys(s)
}

func (s set[V]) Add(values ...V) {
	for _, v := range values {
		s[v] = struct{}{}
	}
}

func (s set[V]) Union(others ...set[V]) set[V] {
	ret := set[V]{}
	for _, s := range append(others, s) {
		ret.Add(s.Keys()...)
	}
	return ret
}

func (s set[V]) Intersect(others ...set[V]) set[V] {
	ret := set[V]{}
	for _, v := range s.Keys() {
		contains := true
		for _, o := range others {
			_, contains = o[v]
			if !contains {
				break
			}
		}
		if contains {
			ret.Add(v)
		}
	}
	return ret
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

func abs[T int](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

type point struct {
	// TODO: allow more types than int?
	X, Y int
}

type grid[T any] [][]T

func (g *grid[T]) InBounds(p point) bool {
	return p.Y >= 0 && p.Y < len(*g) && p.X >= 0 && p.X < len((*g)[0])
}

func (g *grid[T]) At(p point) T {
	return (*g)[p.Y][p.X]
}

func (g *grid[T]) String() string {
	s := ""
	for _, row := range *g {
		for _, cell := range row {
			s += fmt.Sprintf("%4v", cell)
		}
		s += "\n"
	}

	return s
}
