package main

import "testing"

func TestRunProgram(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
		want2 int
	}{
		{"sample", `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`, 5, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := RunProgram(tt.args)
			if got != tt.want {
				t.Errorf("%s failed, got: %d, want: %d", tt.name, got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			got := FixProgram(tt.args)
			if got != tt.want2 {
				t.Errorf("%s failed, got: %d, want: %d", tt.name, got, tt.want)
			}
		})
	}
}
