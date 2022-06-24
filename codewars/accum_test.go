package codewars

import "testing"

func TestAccum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "a"},
			want: "A",
		},
		{
			name: "test2",
			args: args{s: "abcz"},
			want: "A-Bb-Ccc-Zzzz",
		},
		{
			name: "test3",
			args: args{s: "dd"},
			want: "D-Dd",
		},
		{
			name: "test4",
			args: args{s: "Zaa"},
			want: "Z-Aa-Aaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Accum(tt.args.s); got != tt.want {
				t.Errorf("Accum() = %v, want %v", got, tt.want)
			}
		})
	}
}
