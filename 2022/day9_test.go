package main

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func (p point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func TestPointTouches(t *testing.T) {
	for _, tc := range []struct {
		expected bool
		p1, p2   point
	}{
		{true, point{0, 0}, point{0, 0}},
		{true, point{0, 0}, point{1, 0}},
		{true, point{0, 0}, point{0, 1}},
		{true, point{0, 0}, point{1, 1}},

		{false, point{0, 0}, point{2, 0}},
		{false, point{0, 0}, point{0, 2}},
		{false, point{0, 0}, point{2, 1}},
		{false, point{0, 0}, point{1, 2}},

		{false, point{0, 0}, point{2, 2}},
	} {
		t.Run(fmt.Sprintf("%v %v", tc.p1, tc.p2), func(t *testing.T) {
			if actual := tc.p1.TouchesDay9(tc.p2); actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}

			if actual := tc.p2.TouchesDay9(tc.p1); actual != tc.expected {
				t.Errorf("expected %t, got %t (inverse test)", tc.expected, actual)
			}
		})
	}
}

func TestMoveTowards(t *testing.T) {
	for _, tc := range []struct {
		expected point
		p        point
		to       point
	}{
		// not testing moving towards a touching point

		{point{1, 0}, point{0, 0}, point{2, 0}},
		{point{0, 1}, point{0, 0}, point{0, 2}},
		{point{1, 1}, point{0, 0}, point{2, 2}},
		{point{1, 1}, point{0, 0}, point{2, 1}},
		{point{1, 1}, point{0, 0}, point{1, 2}},
	} {
		t.Run(fmt.Sprintf("%v %v", tc.p, tc.to), func(t *testing.T) {
			tc.p.MoveTowardsDay9(tc.to)
			if tc.p != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, tc.p)
			}
		})
	}
}

func TestRopeMoveMotion_2Knots(t *testing.T) {
	r := newRope(2)

	for i, tc := range []struct {
		expected []point
		m        motion
	}{
		{[]point{{4, 0}, {3, 0}}, motion{right, 4}},
		{[]point{{4, 4}, {4, 3}}, motion{up, 4}},
		{[]point{{1, 4}, {2, 4}}, motion{left, 3}},
		{[]point{{1, 3}, {2, 4}}, motion{down, 1}},
		{[]point{{5, 3}, {4, 3}}, motion{right, 4}},
		{[]point{{5, 2}, {4, 3}}, motion{down, 1}},
		{[]point{{0, 2}, {1, 2}}, motion{left, 5}},
		{[]point{{2, 2}, {1, 2}}, motion{right, 2}},
	} {
		t.Run(fmt.Sprintf("%d %v", i, tc.m), func(t *testing.T) {
			r.moveMotion(tc.m)
			if !slices.Equal(r.Knots, tc.expected) {
				t.Errorf("expected knots %v, got %v", tc.expected, r.Knots)
			}
		})
	}
}

func TestRopeMoveMotion_10Knots(t *testing.T) {
	r := newRope(10)
	// set initial state to match example
	// []point{{11, 5}, {11, 5}, {11, 5}, {11, 5}, {11, 5}, {11, 5}, {11, 5}, {11, 5}, {11, 5}, {11, 5}}
	for i := range r.Knots {
		r.Knots[i] = point{11, 5}
	}

	for i, tc := range []struct {
		expected []point
		m        motion
	}{
		{[]point{{16, 5}, {15, 5}, {14, 5}, {13, 5}, {12, 5}, {11, 5}, {11, 5}, {11, 5}, {11, 5}, {11, 5}}, motion{right, 5}},
		{[]point{{16, 13}, {16, 12}, {16, 11}, {16, 10}, {16, 9}, {15, 9}, {14, 8}, {13, 7}, {12, 6}, {11, 5}}, motion{up, 8}},
		{[]point{{8, 13}, {9, 13}, {10, 13}, {11, 13}, {12, 13}, {12, 12}, {12, 11}, {12, 10}, {12, 9}, {12, 8}}, motion{left, 8}},
		{[]point{{8, 10}, {8, 11}, {9, 12}, {10, 12}, {11, 12}, {12, 12}, {12, 11}, {12, 10}, {12, 9}, {12, 8}}, motion{down, 3}},
		{[]point{{25, 10}, {24, 10}, {23, 10}, {22, 10}, {21, 10}, {20, 10}, {19, 10}, {18, 10}, {17, 10}, {16, 10}}, motion{right, 17}},
		{[]point{{25, 0}, {25, 1}, {25, 2}, {25, 3}, {25, 4}, {25, 5}, {24, 5}, {23, 5}, {22, 5}, {21, 5}}, motion{down, 10}},
		{[]point{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}, {9, 0}}, motion{left, 25}},
		{[]point{{0, 20}, {0, 19}, {0, 18}, {0, 17}, {0, 16}, {0, 15}, {0, 14}, {0, 13}, {0, 12}, {0, 11}}, motion{up, 20}},
	} {
		t.Run(fmt.Sprintf("%d %v", i, tc.m), func(t *testing.T) {
			r.moveMotion(tc.m)
			if !slices.Equal(r.Knots, tc.expected) {
				t.Errorf("expected knots %v, got %v", tc.expected, r.Knots)
			}
		})
	}
}
