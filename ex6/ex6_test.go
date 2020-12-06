package main

import "testing"

func TestFindNumAnsPerGroup(t *testing.T) {

	table := [] struct {
		name string
		arg string
		wantUnion int
		wantIntersection int
	} {
		{name: "sample", arg: `abc

a
b
c

ab
ac

a
a
a
a

b`, wantUnion: 11, wantIntersection: 6},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			got := GetTotalYesAnswers(tt.arg, Union)
			if got != tt.wantUnion {
				t.Fatalf("%s, failed, wanted: %d, got: %d", tt.name, tt.wantUnion, got)
			}
		})
		t.Run(tt.name, func(t *testing.T) {

			got := GetTotalYesAnswers(tt.arg, Intersection)
			if got != tt.wantIntersection {
				t.Fatalf("%s, failed, wanted: %d, got: %d", tt.name, tt.wantIntersection, got)
			}
		})
	}

}
