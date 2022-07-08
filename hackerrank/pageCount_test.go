package hackerrank

import "testing"

func Test_pageCount(t *testing.T) {
	type args struct {
		n int32
		p int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test",
			args: args{
				n: 70809,
				p: 46090,
			},
			want: 12359,
		},
		{
			name: "test",
			args: args{
				n: 6,
				p: 4,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pageCount(tt.args.n, tt.args.p); got != tt.want {
				t.Errorf("pageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
