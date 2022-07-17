package hackerrank

import "testing"

func Test_towerBreakers(t *testing.T) {
	type args struct {
		n int32
		m int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test1",
			args: args{
				n: 2,
				m: 6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := towerBreakers(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("towerBreakers() = %v, want %v", got, tt.want)
			}
		})
	}
}
