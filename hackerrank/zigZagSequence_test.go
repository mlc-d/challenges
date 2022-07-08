package hackerrank

import "testing"

func Test_findZigZagSequence(t *testing.T) {
	type args struct {
		arr []int32
		n   int32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findZigZagSequence(tt.args.arr, tt.args.n)
		})
	}
}
