package hackerrank

import (
	"reflect"
	"testing"
)

func Test_maximumPerimeterTriangle(t *testing.T) {
	type args struct {
		sticks []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "test",
			args: args{sticks: []int32{1, 2, 3, 4, 5, 10}},
			want: []int32{3, 4, 5},
		},
		{
			name: "test1",
			args: args{sticks: []int32{1, 1, 1, 2, 3, 5}},
			want: []int32{1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumPerimeterTriangle(tt.args.sticks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumPerimeterTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
