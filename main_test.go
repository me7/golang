package main

import "testing"

func TestSumInt(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
	}{
		// TODO: Add test cases.
		{"123", args{[]int{1, 2, 3}}, 6},
		{"1234", args{[]int{1, 2, 3, 4}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := SumInt(tt.args.nums...); gotTotal != tt.wantTotal {
				t.Errorf("SumInt() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
