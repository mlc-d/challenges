package codewars

import "testing"

func Test_intPow(t *testing.T) {
	type args struct {
		b   int
		exp int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				b:   5,
				exp: 1,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intPow(tt.args.b, tt.args.exp); got != tt.want {
				t.Errorf("intPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastDigit(t *testing.T) {
	type args struct {
		as []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{as: []int{4, 3, 6}},
			want: 4,
		},
		{
			name: "test",
			args: args{as: []int{3, 4, 5}},
			want: 1,
		},
		// Expect(LastDigit( []int{123232,694022,140249} )).To(Equal(6))
		{
			name: "test2",
			args: args{as: []int{123232, 694022, 140249}},
			want: 6,
		},
		{
			name: "test3",
			args: args{as: []int{1, 2, 3, 3, 2}},
			want: 1,
		},
		// Expect(LastDigit( []int{499942,898102,846073} )).To(Equal(6))
		{
			name: "test4",
			args: args{as: []int{499942, 898102, 846073}},
			want: 6,
		},
		// Expect(LastDigit( []int{937640,767456,981242} )).To(Equal(0))
		{
			name: "test5",
			args: args{as: []int{937640, 767456, 981242}},
			want: 0,
		},
		{
			name: "test6",
			args: args{as: []int{0, 0}},
			want: 1,
		},
		{
			name: "test7",
			args: args{as: []int{1, 2}},
			want: 1,
		},
		{
			name: "test8",
			args: args{as: []int{7, 6, 21}},
			want: 1,
		},
		{
			name: "test9",
			args: args{as: []int{3, 4, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastDigit(tt.args.as); got != tt.want {
				t.Errorf("LastDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
