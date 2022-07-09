package hackerrank

import (
	"reflect"
	"testing"
)

func Test_rotateLeft(t *testing.T) {
	type args struct {
		d   int32
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "test",
			args: args{
				d:   4,
				arr: []int32{1, 2, 3, 4, 5},
			},
			want: []int32{5, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateLeft(tt.args.d, tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
