package codewars

import "testing"

func TestRGB(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{values: []int{0, 0, 0}},
			want: "000000",
		},
		{
			name: "test1",
			args: args{values: []int{255, 285, 255}},
			want: "FFFFFF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RGB(tt.args.values...); got != tt.want {
				t.Errorf("RGB() = %v, want %v", got, tt.want)
			}
		})
	}
}
