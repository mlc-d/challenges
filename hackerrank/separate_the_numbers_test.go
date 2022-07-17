package hackerrank

import "testing"

func Test_separateNumbers(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{s: "1234"},
			want: true,
		},
		{
			name: "test1",
			args: args{s: "91011"},
			want: true,
		},
		{
			name: "test2",
			args: args{s: "99100"},
			want: true,
		},
		{
			name: "test3",
			args: args{s: "99910001001"},
			want: true,
		},
		{
			name: "test4",
			args: args{s: "999100010001"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := separateNumbers(tt.args.s); got != tt.want {
				t.Errorf("separateNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
