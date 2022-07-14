package hackerrank

import (
	"reflect"
	"testing"
)

func Test_matchingStrings(t *testing.T) {
	type args struct {
		data    []string
		queries []string
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "test",
			args: args{
				data:    []string{"aba", "baba", "aba", "xzxb"},
				queries: []string{"aba", "xzxb", "ab"},
			},
			want: []int32{2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchingStrings(tt.args.data, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matchingStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
