package main

import "testing"

func TestGetScanningRate(t *testing.T) {
	sample := `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

	tests := []struct {
		name string
		args string
		want int
	}{
		{"sample", sample, 71},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetScanningRate(tt.args); got != tt.want {
				t.Errorf("GetScanningRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapLocations(t *testing.T) {
	sample := `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`

	tests := []struct {
		name string
		args string
		want int
	}{
		{"sample", sample, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapLocations(tt.args); got != tt.want {
				t.Errorf("MapLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}