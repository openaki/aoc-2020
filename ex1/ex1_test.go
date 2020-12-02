package main

import (
	"testing"
)

func TestSolve1a(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{name: "example", args: []int{1721, 979, 366, 299, 675, 1456}, want: 514579},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve1a(tt.args); got != tt.want {
				t.Errorf("Solve1a() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve1b(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{name: "example", args: []int{1721, 979, 366, 299, 675, 1456}, want: 241861950},
	}
	type args struct {
		input []int
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve1b(tt.args); got != tt.want {
				t.Errorf("Solve1b() = %v, want %v", got, tt.want)
			}
		})
	}
}
