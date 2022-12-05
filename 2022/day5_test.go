package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func (s ship) String() string {
	max := 0
	for _, st := range s {
		if l := len(st); l > max {
			max = l
		}
	}

	display := ""
	for i := max - 1; i >= 0; i-- {
		for _, st := range s {
			if i < len(st) {
				display += fmt.Sprintf("[%s]", string(st[i]))
			} else {
				display += "   "
			}
		}
		display += "\n"
	}
	// TODO: col numbers?
	return display
}

func TestReadStacks(t *testing.T) {
	sr := strings.NewReader("" +
		"    [D]    \n" +
		"[N] [C]    \n" +
		"[Z] [M] [P]\n" +
		" 1   2   3 \n",
	)
	scanner := bufio.NewScanner(sr)

	actual, err := readShip(scanner)
	if err != nil {
		t.Fatalf("unable to read stacks: %v", err)
	}

	if !reflect.DeepEqual(actual, ship{
		{'Z', 'N'},
		{'M', 'C', 'D'},
		{'P'},
	}) {
		t.Fatal("stacks do not match")
	}
}
