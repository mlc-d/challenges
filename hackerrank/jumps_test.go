package hackerrank

import "testing"

func Test_kangaroo(t *testing.T) {
	type args struct {
		x1 int32
		v1 int32
		x2 int32
		v2 int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				x1: 2,
				v1: 1,
				x2: 1,
				v2: 2,
			},
			want: "YES",
		},
		{
			name: "test1",
			args: args{
				x1: 2,
				v1: 1,
				x2: 3,
				v2: 3,
			},
			want: "NO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kangaroo(tt.args.x1, tt.args.v1, tt.args.x2, tt.args.v2); got != tt.want {
				t.Errorf("kangaroo() = %v, want %v", got, tt.want)
			}
		})
	}
}
