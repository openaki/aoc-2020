package main

import "testing"

func TestRunIterations(t *testing.T) {
	sample := `.#.
..#
###`
	tests := []struct {
		name string
		args string
		want int
		want2 int
	}{
		{"sample", sample, 112, 848},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunIterations(tt.args, 3); got != tt.want {
				t.Errorf("RunIterations() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := RunIterations(tt.args, 4); got != tt.want2 {
				t.Errorf("RunIterations() = %v, want %v", got, tt.want2)
			}
		})
	}
}
