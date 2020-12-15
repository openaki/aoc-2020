package main

import "testing"
func TestFindIterationNumber(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{"sample1", "0,3,6", 436},
		{"sample1", "1,3,2", 1},
		{"sample1", "2,1,3", 10},
		{"sample1", "1,2,3", 27},
		{"sample1", "2,3,1", 78},
		{"sample1", "3,2,1", 438},
		{"sample1", "3,1,2", 1836},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIterationNumber(tt.args); got != tt.want {
				t.Errorf("FindIterationNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}