package main

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestPastAnswers(t *testing.T) {
	for _, part := range []struct {
		run      runPart
		args     []string
		expected string
	}{
		{runDay1Part1, []string{"day1.input.test"}, "24000"},
		{runDay1Part1, []string{"day1.input"}, "66719"},
		{runDay1Part2, []string{"day1.input.test"}, "45000"},
		{runDay1Part2, []string{"day1.input"}, "198551"},

		{runDay2Part1, []string{"day2.input.test"}, "15"},
		{runDay2Part1, []string{"day2.input"}, "11150"},
		{runDay2Part2, []string{"day2.input.test"}, "12"},
		{runDay2Part2, []string{"day2.input"}, "8295"},

		{runDay3Part1, []string{"day3.input.test"}, "157"},
		{runDay3Part1, []string{"day3.input"}, "7850"},
		{runDay3Part2, []string{"day3.input.test"}, "70"},
		{runDay3Part2, []string{"day3.input"}, "2581"},

		{runDay4Part1, []string{"day4.input.test"}, "2"},
		{runDay4Part1, []string{"day4.input"}, "305"},
		{runDay4Part2, []string{"day4.input.test"}, "4"},
		{runDay4Part2, []string{"day4.input"}, "811"},

		{runDay5Part1, []string{"day5.input.test"}, "CMZ"},
		{runDay5Part1, []string{"day5.input"}, "QPJPLMNNR"},
		{runDay5Part2, []string{"day5.input.test"}, "MCD"},
		{runDay5Part2, []string{"day5.input"}, "BQDNWJPVJ"},

		{runDay6Part1, []string{"day6.input.test1"}, "7"},
		{runDay6Part1, []string{"day6.input.test2"}, "5"},
		{runDay6Part1, []string{"day6.input.test3"}, "6"},
		{runDay6Part1, []string{"day6.input.test4"}, "10"},
		{runDay6Part1, []string{"day6.input.test5"}, "11"},
		{runDay6Part1, []string{"day6.input"}, "1702"},
		{runDay6Part2, []string{"day6.input.test1"}, "19"},
		{runDay6Part2, []string{"day6.input.test2"}, "23"},
		{runDay6Part2, []string{"day6.input.test3"}, "23"},
		{runDay6Part2, []string{"day6.input.test4"}, "29"},
		{runDay6Part2, []string{"day6.input.test5"}, "26"},
		{runDay6Part2, []string{"day6.input"}, "3559"},

		{runDay7Part1, []string{"day7.input.test"}, "95437"},
		{runDay7Part1, []string{"day7.input"}, "1989474"},
		{runDay7Part2, []string{"day7.input.test"}, "24933642"},
		{runDay7Part2, []string{"day7.input"}, "1111607"},

		{runDay8Part1, []string{"day8.input.test"}, "21"},
		{runDay8Part1, []string{"day8.input"}, "1715"},
		{runDay8Part2, []string{"day8.input.test"}, "8"},
		{runDay8Part2, []string{"day8.input"}, "374400"},

		{runDay9Part1, []string{"day9.input.test1"}, "13"},
		// no test 2
		{runDay9Part1, []string{"day9.input"}, "6464"},
		{runDay9Part2, []string{"day9.input.test1"}, "1"},
		{runDay9Part2, []string{"day9.input.test2"}, "36"},
		{runDay9Part2, []string{"day9.input"}, "2604"},

		{runDay10Part1, []string{"day10.input.test"}, "13140"},
		{runDay10Part1, []string{"day10.input"}, "13180"},
		{runDay10Part2, []string{"day10.input.test"}, "" +
			"##..##..##..##..##..##..##..##..##..##..\n" +
			"###...###...###...###...###...###...###.\n" +
			"####....####....####....####....####....\n" +
			"#####.....#####.....#####.....#####.....\n" +
			"######......######......######......####\n" +
			"#######.......#######.......#######.....\n",
		},
		{runDay10Part2, []string{"day10.input"}, "" +
			"####.####.####..##..#..#...##..##..###..\n" +
			"#.......#.#....#..#.#..#....#.#..#.#..#.\n" +
			"###....#..###..#....####....#.#..#.###..\n" +
			"#.....#...#....#....#..#....#.####.#..#.\n" +
			"#....#....#....#..#.#..#.#..#.#..#.#..#.\n" +
			"####.####.#.....##..#..#..##..#..#.###..\n",
		},

		{runDay11Part1, []string{"day11.input.test"}, "10605"},
		{runDay11Part1, []string{"day11.input"}, "113232"},
		{runDay11Part2, []string{"day11.input.test"}, "2713310158"},
		{runDay11Part2, []string{"day11.input"}, "29703395016"},

		{runDay12Part1, []string{"day12.input.test"}, "31"},
		{runDay12Part1, []string{"day12.input"}, "504"},
		{runDay12Part2, []string{"day12.input.test"}, "29"},
		// {runDay12Part2, []string{"day12.input"}, "500"}, // skipping for speed

		{runDay13Part1, []string{"day13.input.test"}, "13"},
		{runDay13Part1, []string{"day13.input"}, "5682"},
		{runDay13Part2, []string{"day13.input.test"}, "140"},
		{runDay13Part2, []string{"day13.input"}, "20304"},
	} {
		t.Run(fmt.Sprintf("%s %#v", runtime.FuncForPC(reflect.ValueOf(part.run).Pointer()).Name(), part.args), func(t *testing.T) {
			ctx := context.TODO()
			answer, err := part.run(ctx, part.args)
			if err != nil {
				t.Fatal(err)
			}
			if answer != part.expected {
				t.Fatalf("Expected answer to be:\n%s\ngot:\n%s", part.expected, answer)
			}
		})
	}
}
