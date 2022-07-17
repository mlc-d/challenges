package hackerrank

import "testing"

func Test_anagram(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test",
			args: args{s: "abccde"},
			want: 2,
		},
		{
			name: "test1",
			args: args{s: "aaabbb"},
			want: 3,
		},
		{
			name: "test2",
			args: args{s: "ab"},
			want: 1,
		},
		{
			name: "test3",
			args: args{s: "abc"},
			want: -1,
		},
		{
			name: "test4",
			args: args{s: "fdhlvosfpafhalll"},
			want: 5,
		},
		{
			name: "test5",
			args: args{s: "xyyx"},
			want: 0,
		},
		{
			name: "test6",
			args: args{s: "xaxbbbxx"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagram(tt.args.s); got != tt.want {
				t.Errorf("anagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
