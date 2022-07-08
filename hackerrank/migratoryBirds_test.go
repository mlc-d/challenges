package hackerrank

import (
	"testing"
)

func Test_migratoryBirds(t *testing.T) {
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
			args: args{arr: []int32{1, 1, 2, 2, 3}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := migratoryBirds(tt.args.arr); got != tt.want {
				t.Errorf("migratoryBirds() = %v, want %v", got, tt.want)
			}
		})
	}
}
