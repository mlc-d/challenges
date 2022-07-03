package hackerrank

import "testing"

func Test_pangrams(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{s: "We promptly judged antique ivory buckles for the prize"},
			want: "not pangram",
		},
		{
			name: "test1",
			args: args{s: "the quick brown fox jumps over the lazy dog"},
			want: "pangram",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pangrams(tt.args.s); got != tt.want {
				t.Errorf("pangrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
