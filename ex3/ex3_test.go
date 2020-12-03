package main

import "testing"

func TestCountTrees(t *testing.T) {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	expected := 7
	t.Run("Example", func(t *testing.T) {
		if got := CountTrees(input, 1, 3); got != expected {
			t.Errorf("CountTrees failed: %v, wanted %v", got, expected)

		}

	})

	t.Run("Example B", func(t *testing.T) {
		if got := CountTreesDifferentSlopes(input); got != 336 {
			t.Errorf("CountTreesDifferentSlopes failed: %v, wanted %v", got, expected)

		}

	})

}
