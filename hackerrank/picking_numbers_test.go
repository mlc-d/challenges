package hackerrank

import "testing"

func Test_pickingNumbers(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test",
			args: args{a: []int32{1, 1, 2, 2, 4, 4, 4, 5, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickingNumbers(tt.args.a); got != tt.want {
				t.Errorf("pickingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
