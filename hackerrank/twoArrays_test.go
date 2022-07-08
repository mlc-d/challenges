package hackerrank

import "testing"

func Test_twoArrays(t *testing.T) {
	type args struct {
		k int32
		A []int32
		B []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				k: 5,
				A: []int32{1, 2, 2, 1},
				B: []int32{3, 3, 3, 4},
			},
			want: "NO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoArrays(tt.args.k, tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("twoArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
