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

func TestReadShip(t *testing.T) {
	actual, err := readShip(bufio.NewScanner(strings.NewReader("" +
		"[N]\n" +
		"[Z]\n" +
		" 1 \n",
	)))
	if err != nil {
		t.Fatalf("unable to read stacks: %v", err)
	}
	expectedTops := "N"
	if actualTops := string(actual.TopCrates()); actualTops != expectedTops {
		t.Fatalf("top crates do not match, expected %q, actual %q", expectedTops, actualTops)
	}
	if !reflect.DeepEqual(actual, ship{
		{'Z', 'N'},
	}) {
		t.Fatal("stacks do not match")
	}

	actual, err = readShip(bufio.NewScanner(strings.NewReader("" +
		"    [D]    \n" +
		"[N] [C]    \n" +
		"[Z] [M] [P]\n" +
		" 1   2   3 \n",
	)))
	if err != nil {
		t.Fatalf("unable to read stacks: %v", err)
	}
	expectedTops = "NDP"
	if actualTops := string(actual.TopCrates()); actualTops != expectedTops {
		t.Fatalf("top crates do not match, expected %q, actual %q", expectedTops, actualTops)
	}
	if !reflect.DeepEqual(actual, ship{
		{'Z', 'N'},
		{'M', 'C', 'D'},
		{'P'},
	}) {
		t.Fatal("stacks do not match")
	}

	actual, err = readShip(bufio.NewScanner(strings.NewReader("" +
		"[M] [H]         [N]                \n" +
		"[S] [W]         [F]     [W] [V]    \n" +
		"[J] [J]         [B]     [S] [B] [F]\n" +
		"[L] [F] [G]     [C]     [L] [N] [N]\n" +
		"[V] [Z] [D]     [P] [W] [G] [F] [Z]\n" +
		"[F] [D] [C] [S] [W] [M] [N] [H] [H]\n" +
		"[N] [N] [R] [B] [Z] [R] [T] [T] [M]\n" +
		"[R] [P] [W] [N] [M] [P] [R] [Q] [L]\n" +
		" 1   2   3   4   5   6   7   8   9 \n",
	)))
	if err != nil {
		t.Fatalf("unable to read stacks: %v", err)
	}
	expectedTops = "MHGSNWWVF"
	if actualTops := string(actual.TopCrates()); actualTops != expectedTops {
		t.Fatalf("top crates do not match, expected %q, actual %q", expectedTops, actualTops)
	}
	if !reflect.DeepEqual(actual, ship{
		{'R', 'N', 'F', 'V', 'L', 'J', 'S', 'M'},
		{'P', 'N', 'D', 'Z', 'F', 'J', 'W', 'H'},
		{'W', 'R', 'C', 'D', 'G'},
		{'N', 'B', 'S'},
		{'M', 'Z', 'W', 'P', 'C', 'B', 'F', 'N'},
		{'P', 'R', 'M', 'W'},
		{'R', 'T', 'N', 'G', 'L', 'S', 'W'},
		{'Q', 'T', 'H', 'F', 'N', 'B', 'V'},
		{'L', 'M', 'H', 'Z', 'N', 'F'},
	}) {
		t.Fatal("stacks do not match")
	}

	actual, err = readShip(bufio.NewScanner(strings.NewReader("" +
		"[M] [H]         [N]                 [H]         [N]             [M] [H]         [N]                 [H]         [N]             [M] [H]         [N]                 [H]         [N]             [M] [H]         [N]                 [H]         [N]             [M] [H]         [N]                 [H]         [N]                                                                                                \n" +
		"[S] [W]         [F]     [W] [V]     [W]         [F]     [W] [V] [S] [W]         [F]     [W] [V]     [W]         [F]     [W] [V] [S] [W]         [F]     [W] [V]     [W]         [F]     [W] [V] [S] [W]         [F]     [W] [V]     [W]         [F]     [W] [V] [S] [W]         [F]     [W] [V]     [W]         [F]                                                                                                \n" +
		"[J] [J]         [B]     [S] [B] [F] [S] [W]     [F]     [W] [J] [J] [B]         [S]     [B] [F] [S] [W]         [F]     [W] [J] [J] [B]         [S]     [B] [F] [S] [W]         [F]     [W] [J] [J] [B]         [S]     [B] [F] [S] [W]         [B] [F] [F] [W] [J] [J]         [B]     [S] [B] [F] [S] [W]     [F]                                                                                                \n" +
		"[L] [F] [G]     [C]     [L] [N] [N] [Z] [D]     [P] [W] [G] [F] [L] [F] [G]     [C]     [L] [N] [N] [Z] [D]     [P] [W] [G] [F] [L] [F] [G]     [C]     [L] [N] [N] [Z] [D]     [P] [W] [G] [F] [L] [F] [G]     [C]     [L] [N] [N] [Z] [D]     [P] [W] [G] [F] [L] [F] [G]     [C]     [L] [N] [N] [Z] [D]     [P] [W] [G]                                                                                        \n" +
		"[V] [Z] [D]     [P] [W] [G] [F] [Z] [D] [C] [S] [W] [M] [N] [H] [V] [Z] [D]     [P] [W] [G] [F] [Z] [D] [C] [S] [W] [M] [N] [H] [V] [Z] [D]     [P] [W] [G] [F] [Z] [D] [C] [S] [W] [M] [N] [H] [V] [Z] [D]     [P] [W] [G] [F] [Z] [D] [C] [S] [W] [M] [N] [H] [V] [Z] [D]     [P] [W] [G] [F] [Z] [D] [C] [S] [W] [M] [N] [H] [V]                                                                                \n" +
		"[F] [D] [C] [S] [W] [M] [N] [H] [H] [N] [R] [B] [Z] [R] [T] [T] [F] [D] [C] [S] [W] [M] [N] [H] [H] [N] [R] [B] [Z] [R] [T] [T] [F] [D] [C] [S] [W] [M] [N] [H] [H] [N] [R] [B] [Z] [R] [T] [T] [F] [D] [C] [S] [W] [M] [N] [H] [H] [N] [R] [B] [Z] [R] [T] [T] [F] [D] [C] [S] [W] [M] [N] [H] [H] [N] [R] [B] [Z] [R] [T] [T] [F] [D] [C] [S]                                                                    \n" +
		"[N] [N] [R] [B] [Z] [R] [T] [T] [M] [D] [C] [S] [W] [M] [N] [H] [N] [N] [R] [B] [Z] [R] [T] [T] [M] [D] [C] [S] [W] [M] [N] [H] [N] [N] [R] [B] [Z] [R] [T] [T] [M] [D] [C] [S] [W] [M] [N] [H] [N] [N] [R] [B] [Z] [R] [T] [T] [M] [D] [C] [S] [W] [M] [N] [H] [N] [N] [R] [B] [Z] [R] [T] [T] [M] [D] [C] [S] [W] [M] [N] [H] [N] [N] [R] [B] [Z] [R] [T] [T]                                                    \n" +
		"[R] [P] [W] [N] [M] [P] [R] [Q] [L] [N] [R] [B] [Z] [R] [T] [T] [R] [P] [W] [N] [M] [P] [R] [Q] [L] [N] [R] [B] [Z] [R] [T] [T] [R] [P] [W] [N] [M] [P] [R] [Q] [L] [N] [R] [B] [Z] [R] [T] [T] [R] [P] [W] [N] [M] [P] [R] [Q] [L] [N] [R] [B] [Z] [R] [T] [T] [R] [P] [W] [N] [M] [P] [R] [Q] [L] [N] [R] [B] [Z] [R] [T] [T] [R] [P] [W] [N] [M] [P] [R] [Q] [L] [N] [R] [B] [Z] [R] [T] [T] [R] [P] [W] [N] [M]\n" +
		" 1   2   3   4   5   6   7   8   9   10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32  33  34  35  36  37  38  39  40  41  42  43  44  45  46  47  48  49  50  51  52  53  54  55  56  57  58  59  60  61  62  63  64  65  66  67  68  69  70  71  72  73  74  75  76  77  78  79  80  81  82  83  84  85  86  87  88  89  90  91  92  93  94  95  96  97  98  98  99 100\n",
	)))
	if err != nil {
		t.Fatalf("unable to read stacks: %v", err)
	}
	expectedTops = "MHGSNWWVFHWSNWWVMHGSNWWVSHDSNWWVMHGSNWWVSHDSNWWVMHGSNWWVSHDSNFWVMHGSNWWVFHWSNWGHVDCSZRTTLNRBZRTTRPWNM"
	if actualTops := string(actual.TopCrates()); actualTops != expectedTops {
		t.Fatalf("top crates do not match, expected %q, actual %q", expectedTops, actualTops)
	}
}
