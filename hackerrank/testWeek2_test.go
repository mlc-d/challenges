package hackerrank

import "testing"

func Test_greatestNumber(t *testing.T) {
	type args struct {
		args []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greatestNumber(tt.args.args...); got != tt.want {
				t.Errorf("greatestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flippingMatrix(t *testing.T) {
	type args struct {
		matrix [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test",
			args: args{matrix: [][]int32{
				{112, 42, 83, 119},
				{56, 125, 56, 49},
				{15, 78, 101, 43},
				{62, 98, 114, 108}},
			},
			want: 414,
		},
		{
			name: "test",
			args: args{matrix: [][]int32{
				{1, 2},
				{3, 4}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flippingMatrix(tt.args.matrix); got != tt.want {
				t.Errorf("flippingMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
