package main

import "testing"

func TestUseAllAdapters(t *testing.T) {
	sample1 := `16
10
15
5
1
11
7
19
6
12
4`
	sample2 := `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`
	_ = sample2
	tests := []struct {
		name string
		content string
		want int
		want2 int
	}{
		{"sample1", sample1, 35, 8},
		{"sample2", sample2, 220, 19208},
	}
	for _, tt := range tests {
		//t.Run(tt.name, func(t *testing.T) {
		//	if got := UseAllAdapters(tt.content); got != tt.want {
		//		t.Errorf("UseAllAdapters() = %v, want %v", got, tt.want)
		//	}
		//})
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPaths(tt.content); got != tt.want2 {
				t.Errorf("UseAllAdapters() = %v, want %v", got, tt.want2)
			}
		})
	}
}
