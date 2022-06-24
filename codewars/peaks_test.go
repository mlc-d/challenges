package codewars

import (
	"reflect"
	"testing"
)

func TestPickPeaks(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want PosPeaks
	}{
		{
			name: "test1",
			args: args{input: []int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1}},
			want: PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		},
		{
			name: "test2",
			args: args{input: []int{13, 9, -2, -5, 8, 8, 14, -2, -3}},
			want: PosPeaks{Pos: []int{6}, Peaks: []int{14}},
		},
		{
			name: "test3",
			args: args{input: []int{2, 2, 5, -4, 2, 4, 12, 0}},
			want: PosPeaks{Pos: []int{2, 6}, Peaks: []int{5, 12}},
		},
		{
			name: "test4",
			args: args{input: []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1}},
			want: PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
		},
		{
			name: "test5",
			args: args{input: []int{3, -5, 14, 11, 11, -4, 12, 12, 7}},
			want: PosPeaks{Pos: []int{2, 6}, Peaks: []int{14, 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PickPeaks(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PickPeaks() = %v, want %v", got, tt.want)
			}
		})
	}
}
