package main

import (
	"fmt"
	"testing"
)

func TestForestVisibleTrees(t *testing.T) {
	for i, f := range []struct {
		expected int
		f        forest
	}{
		{0, forest{}},
		{1, forest{{1}}},
		{2, forest{{1, 1}}},
		{2, forest{{1}, {1}}},
		{8, forest{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}},
		{8, forest{
			{2, 2, 2},
			{2, 1, 2},
			{2, 2, 2},
		}},
		{9, forest{
			{1, 1, 1},
			{1, 2, 1},
			{1, 1, 1},
		}},
		{13, forest{
			{1, 1, 1, 1},
			{1, 2, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		}},

		// test each direction individually
		{13, forest{
			{2, 1, 2, 2},
			{2, 2, 2, 2},
			{2, 2, 1, 2},
			{2, 2, 2, 2},
		}},
		{13, forest{
			{2, 2, 2, 2},
			{1, 2, 2, 2},
			{2, 2, 1, 2},
			{2, 2, 2, 2},
		}},
		{13, forest{
			{2, 2, 2, 2},
			{2, 2, 2, 2},
			{2, 1, 1, 2},
			{2, 1, 2, 2},
		}},
		{13, forest{
			{2, 2, 2, 2},
			{2, 2, 1, 1},
			{2, 2, 1, 2},
			{2, 2, 2, 2},
		}},

		// test input
		{21, forest{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		}},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if actual := f.f.VisibleTrees(); actual != f.expected {
				t.Errorf("expected %d, got %d", f.expected, actual)
			}
		})
	}
}
