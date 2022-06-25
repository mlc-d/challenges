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
			args: args{arr: []int32{7, 2, 234, 23, 45234, 2, 543, 1000}},
			want: 234,
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			radixsort(tt.args.arr)
		})
	}
}
