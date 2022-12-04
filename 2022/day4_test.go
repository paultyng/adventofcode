package main

import (
	"fmt"
	"testing"
)

func TestAssignmentFullyContains(t *testing.T) {
	a := assignment{2, 8}
	if actual := a.FullyContains(a); !actual {
		t.Errorf("Expected %v to fully contain %v", a, a)
	}
	b := assignment{3, 7}
	if actual := a.FullyContains(b); !actual {
		t.Errorf("Expected %v to fully contain %v", a, b)
	}
	if actual := b.FullyContains(a); actual {
		t.Errorf("Expected %v to not fully contain %v", b, a)
	}

	a = assignment{4, 6}
	b = assignment{6, 6}
	if actual := a.FullyContains(b); !actual {
		t.Errorf("Expected %v to fully contain %v", a, b)
	}
	if actual := b.FullyContains(a); actual {
		t.Errorf("Expected %v to not fully contain %v", b, a)
	}
}

func TestAssignmentOverlaps(t *testing.T) {
	for _, tc := range []pairing{
		{assignment{5, 7}, assignment{7, 9}},
		{assignment{2, 8}, assignment{3, 7}},
		{assignment{6, 6}, assignment{4, 6}},
		{assignment{2, 6}, assignment{4, 8}},
	} {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			if actual := tc.Overlaps(); !actual {
				t.Errorf("Expected %v to overlap", tc)
			}
		})
	}
}
