package main

import "testing"

func TestFindInvalid(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		content string
		offset  int
		want int
		want2 int
	}{
		{"Example", `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`, 5, 127, 62},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindInvalid(tt.content, tt.offset); got != tt.want {
				t.Errorf("FindInvalid() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKey(tt.content, tt.want); got != tt.want2 {
				t.Errorf("FindInvalid() = %v, want %v", got, tt.want)
			}
		})
	}
}
