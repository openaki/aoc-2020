package main

import "testing"

func TestValidatePasswords(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{"example", []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePasswordsA(tt.args); got != tt.want {
				t.Errorf("ValidatePasswordsA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatePasswordsB(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{"example for part B", []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePasswordsB(tt.args); got != tt.want {
				t.Errorf("ValidatePasswordsB() = %v, want %v", got, tt.want)
			}
		})
	}
}
