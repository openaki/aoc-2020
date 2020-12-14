package main

import "testing"

func TestGetNextAvaiableBus(t *testing.T) {
	sample := `939
7,13,x,x,59,x,31,19`
	tests := []struct {
		name string
		args string
		want int
		want2 int
	}{
		{"sample", sample, 295, 1068781},
	}
	for _, tt := range tests {
		//t.Run(tt.name, func(t *testing.T) {
		//	if got := GetNextAvaiableBus(tt.args); got != tt.want {
		//		t.Errorf("GetNextAvaiableBus() = %v, want %v", got, tt.want)
		//	}
		//})
		t.Run(tt.name, func(t *testing.T) {
			if got := BruteForce(tt.args); got != tt.want2 {
				t.Errorf("GetNextAvaiableBus() = %v, want %v", got, tt.want2)
			}
		})
	}
}
