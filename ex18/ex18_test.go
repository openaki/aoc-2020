package main

import "testing"

func TestEval(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
		want2 int
	}{
		{"sample", `((1 + 1) * (1 + 1 * 1 + 1) + 6) + 3`, 15, 23},
		{"sample", "2 * 3 + (4 * 5)", 26, 46},
		//{"sample", "(4 * 5)", 26, 0},
		{"sample", "5 + (8 * 3 + 9 + 3 * 4 * 3)", 437, 1445},
		{"sample", "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240, 669060},
		{"sample", "((2 + 4 * 9))", 54, 54},
		{"sample", "((6 + 9 * 8 + 6) + 6)", 132, 216},
		{"sample", "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6)", 6810, 11664},
		{"sample", "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2", 6812, 11666},
		{"sample", "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632, 23340},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eval(tt.args); got != tt.want {
				t.Errorf("Eval() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalAdvanced(tt.args); got != tt.want2 {
				t.Errorf("Eval() = %v, want %v", got, tt.want2)
			}
		})
	}
}
