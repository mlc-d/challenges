package codewars

import (
	"reflect"
	"testing"
)

func TestDiffLists(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{a: []int{1, 2}, b: []int{1}},
			want: []int{2},
		},
		{
			name: "test2",
			args: args{a: []int{1, 2, 2}, b: []int{1}},
			want: []int{2, 2},
		},
		{
			name: "test3",
			args: args{a: []int{1, 2, 2}, b: []int{2}},
			want: []int{1},
		},
		{
			name: "test4",
			args: args{a: []int{1, 2, 2}, b: []int{}},
			want: []int{1, 2, 2},
		},
		{
			name: "test5",
			args: args{a: []int{}, b: []int{1, 2}},
			want: []int{},
		},
		{
			name: "test6",
			args: args{a: []int{1, 2, 3}, b: []int{1, 2}},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffLists(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiffLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
