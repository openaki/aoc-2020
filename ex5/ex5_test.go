package main

import "testing"

func TestGetRowId(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{name: "Test1", args: "FBFBBFFRLR", want: 357},
		{name: "Test2", args: "BFFFBBFRRR", want: 567},
		{name: "Test3", args: "FFFBBBFRRR", want: 119},
		{name: "Test4", args: "BBFFBBFRLL", want: 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRowId(tt.args); got != tt.want {
				t.Errorf("GetRowId() = %v, want %v", got, tt.want)
			}
		})
	}
}