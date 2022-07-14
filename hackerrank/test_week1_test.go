package hackerrank

import "testing"

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
