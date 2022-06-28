package codewars

import "testing"

func TestBeeramid(t *testing.T) {
	type args struct {
		bonus int
		price float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				bonus: 1500,
				price: 2,
			},
			want: 12,
		},
		{
			name: "test1",
			args: args{
				bonus: 5000,
				price: 3,
			},
			want: 16,
		},
		{
			name: "test2",
			args: args{
				bonus: 9,
				price: 2.0,
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				bonus: 21,
				price: 1.5,
			},
			want: 3,
		},
		{
			name: "test4",
			args: args{
				bonus: 21,
				price: 3.0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Beeramid(tt.args.bonus, tt.args.price); got != tt.want {
				t.Errorf("Beeramid() = %v, want %v", got, tt.want)
			}
		})
	}
}
