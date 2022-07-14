package codewars

import "testing"

func TestAbs(t *testing.T) {
	type args struct {
		number float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test",
			args: args{number: 2.0},
			want: 2.0,
		},
		{
			name: "test",
			args: args{number: -2.0},
			want: 2.0,
		},
		{
			name: "test",
			args: args{number: -4.5},
			want: 4.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.number); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxLengthDiff(t *testing.T) {
	type args struct {
		a1 []string
		a2 []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				a1: []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"},
				a2: []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxLengthDiff(tt.args.a1, tt.args.a2); got != tt.want {
				t.Errorf("MaxLengthDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
