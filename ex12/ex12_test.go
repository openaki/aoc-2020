package main

import "testing"

func TestRunInstructions(t *testing.T) {
	sample1 := `F10
N3
F7
R90
F11`
	_ = sample1
	tests := []struct {
		name string
		args string
		want int
		want2 int
	}{
		{"sample", sample1, 25, 286},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunInstructions(tt.args); got != tt.want {
				t.Errorf("RunInstructions() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := RunInstructionsWithWayPoint(tt.args); got != tt.want2 {
				t.Errorf("RunInstructions() = %v, want %v", got, tt.want2)
			}
		})
	}
}
