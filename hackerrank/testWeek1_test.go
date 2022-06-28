package hackerrank

import "testing"

func Test_countSort(t *testing.T) {
	type args struct {
		n   int32
		exp int32
		arr []int32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countSort(tt.args.n, tt.args.exp, tt.args.arr)
		})
	}
}

func Test_FindMedian(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test",
			args: args{arr: []int32{37, 2, 234, 23, 1}},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMedian(tt.args.arr); got != tt.want {
				t.Errorf("FindMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMax(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name    string
		args    args
		wantRes int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := getMax(tt.args.arr); gotRes != tt.wantRes {
				t.Errorf("getMax() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_radixsort(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{arr: []int32{15, 132, 1, 99, 1002}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RadixSort(tt.args.arr)
		})
	}
}
