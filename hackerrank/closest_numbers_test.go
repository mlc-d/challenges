package hackerrank

import (
	"reflect"
	"testing"
)

func Test_closestNumbers(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "test",
			args: args{arr: []int32{-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854, -520, -470}},
			want: []int32{-520, -470, -20, 30},
		},
		{
			name: "test1",
			args: args{arr: []int32{-5, 15, 25, 71, 63}},
			want: []int32{63, 71},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestNumbers(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
