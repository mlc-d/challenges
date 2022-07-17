package hackerrank

import "testing"

func Test_transform(t *testing.T) {
	type args struct {
		c int32
		k int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test",
			args: args{
				c: 66,
				k: 26,
			},
			want: 66,
		},
		{
			name: "test1",
			args: args{
				c: 95,
				k: 100,
			},
			want: 117,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transform(tt.args.c, tt.args.k); got != tt.want {
				t.Errorf("transform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_caesarCipher(t *testing.T) {
	type args struct {
		s string
		k int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "abcdefghijklmnopqrstuvwxyz",
				k: 3,
			},
			want: "defghijklmnopqrstuvwxyzabc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caesarCipher(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("caesarCipher() = %v, want %v", got, tt.want)
			}
		})
	}
}
