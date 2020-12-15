package main

import "testing"

func TestRunProgram(t *testing.T) {
	sample := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`
	sample2 := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
	_ = sample
	tests := []struct {
		name string
		args string
		want int
	}{
		{"sample", sample2, 165},
	}
	for _, tt := range tests {
		//t.Run(tt.name, func(t *testing.T) {
		//	if got := RunProgram(tt.args); got != tt.want {
		//		t.Errorf("RunProgram() = %v, want %v", got, tt.want)
		//	}
		//})
		t.Run(tt.name, func(t *testing.T) {
			if got := RunMemoryDecoder(tt.args); got != tt.want {
				t.Errorf("RunProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}
