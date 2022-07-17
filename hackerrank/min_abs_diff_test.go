package hackerrank

import "testing"

func Test_minimumAbsoluteDifference(t *testing.T) {
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
			args: args{arr: []int32{-2, 2, 4}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAbsoluteDifference(tt.args.arr); got != tt.want {
				t.Errorf("minimumAbsoluteDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
