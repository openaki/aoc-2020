package main

import "testing"

func TestGetSeatsAtFixPoint(t *testing.T) {
	sample := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

	tests := []struct {
		name string
		args string
		want int
		want2 int
	}{
		{"Sample", sample, 37, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSeatsAtFixPoint(tt.args, 4, occupiedNeighbors); got != tt.want {
				t.Errorf("GetSeatsAtFixPoint() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSeatsAtFixPoint(tt.args, 5, occupiedNeighborsB); got != tt.want2 {
				t.Errorf("GetSeatsAtFixPointB() = %v, want %v", got, tt.want2)
			}
		})
	}
}

func Test_occupiedNeighborsB(t *testing.T) {
	sample:= `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`
	sample2 := `.............
.L.L.#.#.#.#.
.............`

	sample3 := `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`
	type args struct {
		grid [][]uint8
		ci   int
		cj   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample", args{getGrid(sample), 4, 3}, 8},
		{"sample2", args{getGrid(sample2), 1, 1}, 0},
		{"sample3", args{getGrid(sample3), 3, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := occupiedNeighborsB(tt.args.grid, tt.args.ci, tt.args.cj); got != tt.want {
				t.Errorf("occupiedNeighborsB() = %v, want %v", got, tt.want)
			}
		})
	}
}