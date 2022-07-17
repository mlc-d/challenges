package hackerrank

import "testing"

func Test_maxMin(t *testing.T) {
	type args struct {
		k   int32
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test",
			args: args{
				k:   3,
				arr: []int32{100, 200, 300, 350, 400, 401, 402},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMin(tt.args.k, tt.args.arr); got != tt.want {
				t.Errorf("maxMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
